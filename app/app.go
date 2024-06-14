package app

import (
	"clean-code-unit-test/config"
	"clean-code-unit-test/model/dto"
	"clean-code-unit-test/router"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/joho/godotenv"
)

func initEnv() (dto.ConfigData, error) {
	// Code for load data dari env
	var configData dto.ConfigData
	if err := godotenv.Load(".env"); err != nil {
		return configData, err
	}

	if port := os.Getenv("PORT"); port != "" {
		configData.AppConfig.Port = port
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbMaxIdle := os.Getenv("MAX_IDLE")
	dbMaxConn := os.Getenv("MAX_CONN")
	dbMaxLifeTime := os.Getenv("MAX_LIFE_TIME")
	logMode := os.Getenv("LOG_MODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" ||
		dbName == "" || dbMaxIdle == "" || dbMaxConn == "" || dbMaxLifeTime == "" || logMode == "" {
		return configData, errors.New("DB config is not set")
	}

	var err error
	configData.DbConfig.MaxConn, err = strconv.Atoi(dbMaxConn)
	if err != nil {
		return configData, err
	}

	configData.DbConfig.MaxIdle, err = strconv.Atoi(dbMaxIdle)
	if err != nil {
		return configData, err
	}

	configData.DbConfig.Host = dbHost
	configData.DbConfig.DbPort = dbPort
	configData.DbConfig.User = dbUser
	configData.DbConfig.Pass = dbPass
	configData.DbConfig.Database = dbName
	configData.DbConfig.MaxLifeTime = dbMaxLifeTime
	configData.DbConfig.LogMode, err = strconv.Atoi(logMode)
	if err != nil {
		return configData, err
	}

	return configData, nil
}

func RunService() {
	//adding zerolog
	zerolog.TimeFieldFormat = "02-01-2006 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	//load config di dalam file env
	configData, err := initEnv()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Info().Msg(fmt.Sprintf("config data %v", configData))

	conn, err := config.ConnectToDB(configData, log.Logger)
	if err != nil {
		log.Error().Msg("RunService.ConnectToDB.err : " + err.Error())
		return
	}

	duration, err := time.ParseDuration(configData.DbConfig.MaxLifeTime)
	if err != nil {
		log.Error().Msg("RunService.duration.err : " + err.Error())
		return
	}

	conn.SetConnMaxLifetime(duration)
	conn.SetMaxIdleConns(configData.DbConfig.MaxIdle)
	conn.SetMaxOpenConns(configData.DbConfig.MaxConn)

	defer func() {
		errClose := conn.Close()
		if errClose != nil {
			log.Error().Msg(errClose.Error())
		}
	}()

	// set up timezone
	time.Local = time.FixedZone("Asia/Jakarta", 7*60*60)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"POST", "DELETE", "GET", "OPTIONS", "PUT"},
		AllowHeaders: []string{
			"Origin", "Content-Type",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           120 * time.Second,
	}))

	log.Logger = log.With().Caller().Logger()

	r.Use(logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(os.Stdout).With().Logger()
		}),
	))

	//gin recovery for handle panic
	r.Use(gin.Recovery())

	//initialize domain module
	initializeDomainModule(r, conn)

	version := "0.0.1"
	log.Info().Msg(fmt.Sprintf("Service Running version %s", version))
	addr := flag.String("port: ", ":"+configData.AppConfig.Port, "Address to listen and serve")
	err = r.Run(*addr)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func initializeDomainModule(r *gin.Engine, db *sql.DB) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")
	//check health
	router.InitRoute(v1Group, db)
}

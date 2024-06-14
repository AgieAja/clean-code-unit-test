package config

import (
	"clean-code-unit-test/model/dto"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func ConnectToDB(in dto.ConfigData, logger zerolog.Logger) (*sql.DB, error) {
	// Code for connect to DB
	logger.Info().Msg("Trying Connect to DB . . .")

	// Initialize database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", in.DbConfig.Host,
		in.DbConfig.User, in.DbConfig.Pass, in.DbConfig.Database, in.DbConfig.DbPort)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to open database connection")
		return nil, err
	}

	// Test database connection
	if err = db.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping database")
		return nil, err
	}

	logger.Info().Msg("Successfully connected to the database")
	return db, nil
}

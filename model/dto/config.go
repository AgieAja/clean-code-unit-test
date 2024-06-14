package dto

import "database/sql"

type (
	ConfigData struct {
		DbConfig  DbConfig
		AppConfig AppConfig
	}

	DbConfig struct {
		Host        string
		DbPort      string
		User        string
		Pass        string
		Database    string
		MaxIdle     int
		MaxConn     int
		MaxLifeTime string
		LogMode     int
	}

	AppConfig struct {
		Port string
	}

	Db struct {
		*sql.DB
	}
)

package core_config

import (
	core_model "gin-api/core/model"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDBSys() {
	logger := &log.Logger
	logger.Info().Str("logtype", "ConnectDB").Msg("Connecting to Sys DB...")

	var err error
	if ENV.DB_SYS_DRIVER == "sqlite" {
		DB_SYS, err = gorm.Open(sqlite.Open(ENV.DB_SYS_URL), &gorm.Config{})
	}

	if err != nil {
		panic("Failed connect to Sys DB")
	} else {
		logger.Info().Str("logtype", "ConnectDB").Msg("CONNECTED to Sys DB!")
		logger.Info().Str("logtype", "ConnectDB").Msg("Migrate Sys DB...")
		migrateDBSys()
	}
}

func migrateDBSys() {
	DB_SYS.Table("m_user").AutoMigrate(&core_model.User{})
}

func DisconnectDBSys() {
	logger := &log.Logger
	logger.Info().Str("logtype", "DisconnectDB").Msg("Disconnecting from Sys DB...")

	db, err := DB_SYS.DB()
	db.Close()

	if err != nil {
		panic("Failed disconnect from Sys DB")
	} else {
		logger.Info().Str("logtype", "DisconnectDB").Msg("DISCONNECTED from Sys DB!")
	}
}

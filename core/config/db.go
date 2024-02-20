package config

import (
	"fmt"
	"gin-api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDBSys(ENV *Config, DB *gorm.DB) {
	fmt.Println("Connecting to Sys DB...")

	var err error
	if ENV.DB_SYS_DRIVER == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(ENV.DB_SYS_URL), &gorm.Config{})
	}

	if err != nil {
		panic("Failed connect to Sys DB")
	} else {
		fmt.Println("CONNECTED to Sys DB!")
		fmt.Println("Migrate Sys DB...")
		migrateDBSys(DB)
	}
}

func migrateDBSys(DB *gorm.DB) {
	DB.Table("m_user").AutoMigrate(&model.User{})
}

func DisconnectDBSys(ENV *Config, DB *gorm.DB) {
	fmt.Println("Disconnecting from Sys DB...")

	db, err := DB.DB()
	db.Close()

	if err != nil {
		panic("Failed disconnect from Sys DB")
	} else {
		fmt.Println("DISCONNECTED from Sys DB!")
	}
}

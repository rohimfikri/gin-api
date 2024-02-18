package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupGin(ENV *Config) {
	fmt.Printf("Setup Gin on '%v' mode...\n", ENV.GIN_MODE)

	if ENV.GIN_MODE == "release" || ENV.GIN_MODE == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	fmt.Printf("Gin '%v' mode setup has been DONE!\n", ENV.GIN_MODE)
}

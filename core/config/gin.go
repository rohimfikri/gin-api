package core_config

import (
	"fmt"
	core_middleware "gin-api/core/middleware"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupGin() *gin.Engine {

	fmt.Printf("Setup Gin on '%v' mode...\n", ENV.GIN_MODE)

	if strings.ToUpper(ENV.GIN_MODE) == "RELEASE" || strings.ToUpper(ENV.GIN_MODE) == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()
	engine.Use(core_middleware.ZeroJSONLogger(&ENV.GIN_MODE, &ENV.LOG_LEVEL, &ENV.LOG_PRETTY))
	engine.Use(gin.Recovery())

	fmt.Printf("Gin '%v' mode setup has been DONE!\n", ENV.GIN_MODE)

	return engine
}

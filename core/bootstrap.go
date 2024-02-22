package core

import (
	core_config "gin-api/core/config"

	"github.com/gin-gonic/gin"
)

func SetupApp() *gin.Engine {
	core_config.LoadConfig()
	ginEngine := core_config.SetupGin()
	core_config.ConfigValidator()
	core_config.ConnectDBSys()
	SetupRouter(core_config.ENV, ginEngine)

	return ginEngine
}

func FlushApp() {
	core_config.DisconnectDBSys()
}

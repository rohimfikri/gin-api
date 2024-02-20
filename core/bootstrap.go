package core

import (
	"gin-api/core/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ENV *config.Config
var DB_SYS *gorm.DB

func SetupApp() *gin.Engine {
	config.LoadConfig(&ENV)
	ginEngine := config.SetupGin(ENV)
	config.ConnectDBSys(ENV, DB_SYS)
	SetupRouter(ENV, ginEngine)

	return ginEngine
}

func FlushApp() {
	config.DisconnectDBSys(ENV, DB_SYS)
}

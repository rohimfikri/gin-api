package core

import (
	"gin-api/core/config"

	"github.com/gin-gonic/gin"
)

var ENV *config.Config

func SetupApp(r *gin.Engine) {
	config.LoadConfig(&ENV)
	config.SetupGin(ENV)
	config.ConnectDBSys(ENV)
	SetupRouter(r)
}

func FlushApp() {
	// config.DisconnectDBSys(ENV)
}

package core

import (
	"fmt"
	"gin-api/api"
	"gin-api/core/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ENV *config.Config, r *gin.Engine) {
	fmt.Println("Configuring app router...")

	r.ForwardedByClientIP = true
	r.SetTrustedProxies(strings.Split(ENV.TRUSTED_PROXIES, ","))
	api.SetupPublicRouter(ENV, r)
	api.SetupAuthRouter(ENV, r)

	fmt.Println("App router has been CONFIGURED!")
}

package core

import (
	"gin-api/api"
	core_type "gin-api/core/type"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupRouter(ENV *core_type.Config, r *gin.Engine) {
	logger := &log.Logger
	logger.Info().Str("logtype", "SetupApp").Msg("Configuring app router...")

	r.ForwardedByClientIP = true
	r.SetTrustedProxies(strings.Split(ENV.TRUSTED_PROXIES, ","))
	api.SetupPublicRouter(ENV, r)
	api.SetupAuthRouter(ENV, r)
	api.SetupUserRouter(ENV, r)

	logger.Info().Str("logtype", "SetupApp").Msg("App router has been CONFIGURED!")
}

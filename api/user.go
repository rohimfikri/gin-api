package api

import (
	api_user "gin-api/api/user"
	core_type "gin-api/core/type"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupUserRouter(ENV *core_type.Config, r *gin.Engine) {
	logger := &log.Logger
	logger.Info().Str("logtype", "SetupRouter").Msg("Configuring 'user' router...")

	api_user.SetupUserApi(r)

	logger.Info().Str("logtype", "SetupRouter").Msg("'user' router has been CONFIGURED!")
}

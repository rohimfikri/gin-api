package api

import (
	api_auth "gin-api/api/auth"
	core_type "gin-api/core/type"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupAuthRouter(ENV *core_type.Config, r *gin.Engine) {
	logger := &log.Logger
	logger.Info().Str("logtype", "SetupRouter").Msg("Configuring 'auth' router...")

	auth_r := r.Group("/auth")
	api_auth.SetupProfileApi(auth_r)

	logger.Info().Str("logtype", "SetupRouter").Msg("'auth' router has been CONFIGURED!")
}

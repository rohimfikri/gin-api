package api

import (
	api_public "gin-api/api/public"
	core_type "gin-api/core/type"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupPublicRouter(ENV *core_type.Config, r *gin.Engine) {
	logger := &log.Logger
	logger.Info().Str("logtype", "SetupRouter").Msg("Configuring 'public' router...")

	api_public.SetupPingApi(r)
	api_public.SetupEnvApi(ENV, r)
	api_public.SetupGeneratorApi(ENV, r)
	api_public.SetupUserApi(r)

	logger.Info().Str("logtype", "SetupRouter").Msg("'public' router has been CONFIGURED!")
}

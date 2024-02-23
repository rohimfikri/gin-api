package api

import (
	api_public "gin-api/api/public"
	core_type "gin-api/core/type"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupPublicRouter(PUBLIC_ROUTER *[]string, ENV *core_type.Config, r *gin.Engine) {
	logger := &log.Logger
	logger.Info().Str("logtype", "SetupRouter").Msg("Configuring 'public' router...")

	SetupPublicPath(PUBLIC_ROUTER)
	api_public.SetupPingApi(r)
	api_public.SetupEnvApi(ENV, r)
	api_public.SetupGeneratorApi(ENV, r)
	api_public.SetupUserApi(r)

	logger.Info().Str("logtype", "SetupRouter").Msg("'public' router has been CONFIGURED!")
}

func SetupPublicPath(PUBLIC_ROUTER *[]string) {
	*PUBLIC_ROUTER = append(*PUBLIC_ROUTER,
		"/ping",
		"/generate/short-uuid",
		"/generate/bcrypt",
		"/user/register",
		"/user/change-password",
	)
}

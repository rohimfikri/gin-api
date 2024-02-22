package api_public

import (
	public_controller "gin-api/api/public/controller"
	core_type "gin-api/core/type"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupGeneratorApi(ENV *core_type.Config, r *gin.Engine) {
	logger := &log.Logger
	g_r := r.Group("/generate")

	logger.Info().Str("logtype", "Router").Msg("Add GET::/generate/short-uuid router")
	g_r.GET("/short-uuid", public_controller.ShortUUID())

	logger.Info().Str("logtype", "Router").Msg("Add POST::/generate/bcrypt router")
	g_r.POST("/bcrypt", public_controller.Bcrypt())
}

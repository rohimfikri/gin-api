package api_public

import (
	core_type "gin-api/core/type"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupEnvApi(ENV *core_type.Config, r *gin.Engine) {
	logger := &log.Logger
	env_r := r.Group("/env")

	logger.Info().Str("logtype", "Router").Msg("Add GET::/env/port router")
	env_r.GET("/port", func(c *gin.Context) {
		c.String(http.StatusOK, ENV.PORT)
	})
}

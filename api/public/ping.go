package api_public

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupPingApi(r *gin.Engine) {
	logger := &log.Logger

	logger.Info().Str("logtype", "Router").Msg("Add GET::/ping router")
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

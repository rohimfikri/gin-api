package api_auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupProfileApi(r *gin.RouterGroup) {
	logger := &log.Logger
	logger.Info().Str("logtype", "Router").Msg("Add GET::/me router")
	r.GET("/me", func(c *gin.Context) {
		c.String(http.StatusOK, "My Profile")
	})
}

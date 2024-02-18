package api_public

import (
	"fmt"
	"gin-api/core/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupEnvApi(ENV *config.Config, r *gin.Engine) {
	env_r := r.Group("/env")

	fmt.Println("Add GET::/env/port router")
	// Ping test
	env_r.GET("/port", func(c *gin.Context) {
		c.String(http.StatusOK, ENV.PORT)
	})
}

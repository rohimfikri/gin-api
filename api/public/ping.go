package api_public

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupPingApi(r *gin.Engine) {
	fmt.Println("Add GET::/ping router")
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

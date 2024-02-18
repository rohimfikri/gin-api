package api_auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProfileApi(r *gin.RouterGroup) {
	fmt.Println("Add GET::/me router")
	// Ping test
	r.GET("/me", func(c *gin.Context) {
		c.String(http.StatusOK, "My Profile")
	})
}

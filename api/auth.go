package api

import (
	"fmt"
	api_auth "gin-api/api/auth"
	"gin-api/core/config"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(ENV *config.Config, r *gin.Engine) {
	fmt.Println("Configuring 'auth' router...")

	auth_r := r.Group("/auth")
	api_auth.SetupProfileApi(auth_r)

	fmt.Println("'auth' router has been CONFIGURED!")
}

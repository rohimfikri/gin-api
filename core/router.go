package core

import (
	"fmt"
	"gin-api/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	fmt.Println("Configuring app router...")

	api.SetupPublicRouter(ENV, r)
	api.SetupAuthRouter(ENV, r)

	fmt.Println("App router has been CONFIGURED!")
}

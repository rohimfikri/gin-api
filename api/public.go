package api

import (
	"fmt"
	api_public "gin-api/api/public"
	"gin-api/core/config"

	"github.com/gin-gonic/gin"
)

func SetupPublicRouter(ENV *config.Config, r *gin.Engine) {
	fmt.Println("Configuring 'public' router...")

	api_public.SetupPingApi(r)
	api_public.SetupEnvApi(ENV, r)
	api_public.SetupGeneratorApi(ENV, r)

	fmt.Println("'public' router has been CONFIGURED!")
}

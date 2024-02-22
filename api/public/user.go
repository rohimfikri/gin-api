package api_public

import (
	public_controller "gin-api/api/public/controller"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupUserApi(r *gin.Engine) {
	logger := &log.Logger
	user_r := r.Group("/user")

	logger.Info().Str("logtype", "Router").Msg("Add GET::/user/register router")
	user_r.POST("/register", public_controller.UserRegister())

	logger.Info().Str("logtype", "Router").Msg("Add GET::/user/change-password router")
	user_r.PUT("/change-password", public_controller.ChangePassword())
}

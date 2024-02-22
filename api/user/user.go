package api_user

import (
	user_controller "gin-api/api/user/controller"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupUserApi(r *gin.Engine) {
	logger := &log.Logger
	users_r := r.Group("/users")

	logger.Info().Str("logtype", "Router").Msg("Add GET::/users/active router")
	users_r.GET("/active", user_controller.GetActiveUsers())

	user_r := r.Group("/user")

	logger.Info().Str("logtype", "Router").Msg("Add GET::/user/:id router")
	user_r.GET("/:id", user_controller.GetUserById())
}

package public_controller

import (
	public_handler "gin-api/api/public/handler"
	public_type "gin-api/api/public/type"
	core_handler "gin-api/core/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		objReq := public_type.UserRegisterRequest{}
		if err := c.ShouldBind(&objReq); err == nil {

			if user, err := public_handler.RegisterUser(&objReq); err == nil {
				resp := core_handler.ResponseParams{
					StatusCode: http.StatusCreated,
					Message:    "created",
					Data:       user,
				}
				resp.HandleResponse(c)
			} else {
				core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
			}
		} else {
			core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
		}
	}
}

func ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		objReq := public_type.ChangePasswordRequest{}
		if err := c.ShouldBind(&objReq); err == nil {

			if user, err := public_handler.ChangePassword(&objReq); err == nil {
				resp := core_handler.ResponseParams{
					StatusCode: http.StatusOK,
					Message:    "updated",
					Data:       user,
				}
				resp.HandleResponse(c)
			} else {
				core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
			}
		} else {
			core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
		}
	}
}

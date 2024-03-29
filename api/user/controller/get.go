package user_controller

import (
	user_handler "gin-api/api/user/handler"
	user_type "gin-api/api/user/type"
	core_handler "gin-api/core/handler"
	core_helper "gin-api/core/helper"
	core_type "gin-api/core/type"
	"net/http"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func GetActiveUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		objReq := user_type.GetActiveUsersRequest{}
		if err := c.ShouldBind(&objReq); err == nil {
			if users, err := user_handler.GetActiveUsers(&objReq); err == nil {
				var resp *core_type.DataResponse
				exclude := []string{"password", "updated_at"}
				if resp, err = core_helper.ConvertToDataResponse(users, &exclude); err == nil {
					resp := core_handler.ResponseParams{
						StatusCode: http.StatusOK,
						Message:    "ok",
						Data:       resp,
					}
					resp.HandleResponse(c)
				} else {
					core_handler.HandleError(c, &core_handler.InternalServerError{Message: err.Error()})
				}
			} else {
				core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
			}
		} else {
			core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
		}
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		objReq := user_type.GetUserByIDRequest{}
		if err := c.ShouldBindUri(&objReq); err == nil {
			if user, err := user_handler.GetUserByID(&objReq); err == nil {
				user.Password = ""
				user.UpdatedAt = time.Time{}
				m := structs.Map(user)
				resp := core_handler.ResponseParams{
					StatusCode: http.StatusOK,
					Message:    "ok",
					Data:       m,
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

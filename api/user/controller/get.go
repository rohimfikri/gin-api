package user_controller

import (
	user_handler "gin-api/api/user/handler"
	user_type "gin-api/api/user/type"
	core_handler "gin-api/core/handler"
	core_helper "gin-api/core/helper"
	core_type "gin-api/core/type"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
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

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		objReq := user_type.GetUserByIdRequest{}
		if err := c.ShouldBind(&objReq); err == nil {
			// if users, err := user_handler.GetActiveUsers(&objReq); err == nil {
			// 	var resp *core_type.DataResponse
			// 	exclude := []string{"password", "updated_at"}
			// 	if resp, err = core_helper.ConvertToDataResponse(users, &exclude); err == nil {
			// 		resp := core_handler.ResponseParams{
			// 			StatusCode: http.StatusOK,
			// 			Message:    "ok",
			// 			Data:       resp,
			// 		}
			// 		resp.HandleResponse(c)
			// 	} else {
			// 		core_handler.HandleError(c, &core_handler.InternalServerError{Message: err.Error()})
			// 	}
			// } else {
			// 	core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
			// }
			u := shortuuid.New()
			resp := core_handler.ResponseParams{
				StatusCode: http.StatusOK,
				Message:    "ok",
				Data:       u,
			}

			resp.HandleResponse(c)
		} else {
			core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
		}
	}
}

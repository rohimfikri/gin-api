package public_controller

import (
	public_type "gin-api/api/public/type"
	core_handler "gin-api/core/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt() gin.HandlerFunc {
	return func(c *gin.Context) {
		objReq := public_type.GenerateBcryptRequest{}
		if err := c.ShouldBind(&objReq); err == nil {
			if hash, err := bcrypt.GenerateFromPassword([]byte(objReq.Password), bcrypt.DefaultCost); err == nil {
				ret := map[string]interface{}{
					"pass": objReq.Password,
					"hash": string(hash),
				}
				resp := core_handler.ResponseParams{
					StatusCode: http.StatusOK,
					Message:    "ok",
					Data:       ret,
				}
				resp.HandleResponse(c)
			} else {
				core_handler.HandleError(c, &core_handler.InternalServerError{Message: err.Error()})
			}
		} else {
			core_handler.HandleError(c, &core_handler.BadRequestError{Message: err.Error()})
		}
	}
}

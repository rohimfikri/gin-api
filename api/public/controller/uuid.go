package public_controller

import (
	core_handler "gin-api/core/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
)

func ShortUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := shortuuid.New()
		resp := core_handler.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       u,
		}

		resp.HandleResponse(c)
	}
}

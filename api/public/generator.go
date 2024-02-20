package api_public

import (
	"fmt"
	"gin-api/core/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"golang.org/x/crypto/bcrypt"
)

func SetupGeneratorApi(ENV *config.Config, r *gin.Engine) {
	g_r := r.Group("/generate")

	fmt.Println("Add GET::/generate/short-uuid router")
	g_r.GET("/short-uuid", ShortUuid)

	fmt.Println("Add POST::/generate/bcrypt router")
	g_r.POST("/bcrypt", Bcrypt)
}

func ShortUuid(c *gin.Context) {
	u := shortuuid.New()
	c.String(http.StatusOK, u)
}

func Bcrypt(c *gin.Context) {
	objReq := GenerateBcryptRequest{}
	if err := c.ShouldBind(&objReq); err == nil {
		if hash, err := bcrypt.GenerateFromPassword([]byte(objReq.Password), bcrypt.DefaultCost); err == nil {
			ret := map[string]interface{}{
				"pass": objReq.Password,
				"hash": string(hash),
			}
			c.JSON(http.StatusOK, ret)
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}

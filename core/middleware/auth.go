package core_middleware

import (
	core_handler "gin-api/core/handler"
	core_helper "gin-api/core/helper"

	"github.com/gin-gonic/gin"
)

func JWTAuthenticate(secretKey *string, PUBLIC_ROUTER *[]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println(c.Request.URL.Path)
		// fmt.Println(c.FullPath())

		path := c.FullPath()
		// var rolePermission map[string][]string
		if core_helper.CheckPublicRouter(&path, PUBLIC_ROUTER) {
			c.Next()
		} else {
			clientToken := c.Request.Header.Get("token")
			if clientToken == "" {
				core_handler.HandleError(c, &core_handler.UnauthorizedError{Message: "Not Authorized"})
				return
			}

			claims, err := core_helper.ValidateToken(&clientToken, secretKey)
			if err != nil {
				core_handler.HandleError(c, &core_handler.UnauthorizedError{Message: err.Error()})
				return
			}
			c.Set("email", claims.Email)
			c.Set("first_name", claims.FirstName)
			c.Set("last_name", claims.LastName)
			c.Set("username", claims.Username)
			c.Next()
		}
	}
}

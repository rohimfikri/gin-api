package handler

import (
	"gin-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e *NotFoundError) Error() string {
	return e.Message
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func (e *InternalServerError) Error() string {
	return e.Message
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusNotImplemented
	}

	response := helper.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	}.Response()

	c.JSON(statusCode, response)
}

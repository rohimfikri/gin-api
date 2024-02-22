package core_handler

import (
	core_type "gin-api/core/type"

	"github.com/gin-gonic/gin"
)

type ResponseParams struct {
	StatusCode int
	Message    string
	Data       any
}

type ResponsePagingParams struct {
	StatusCode int
	Message    string
	Paginate   *core_type.Paginate
	Data       any
}

func (params ResponseParams) HandleResponse(c *gin.Context) {
	var response any
	status := false

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = true
	}

	if params.Data != nil {
		response = &core_type.ResponseWithoutPagingData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &core_type.ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	c.JSON(params.StatusCode, response)
}

func (params ResponsePagingParams) HandleResponse(c *gin.Context) {
	var response any
	status := false

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = true
	}

	if params.Data != nil {
		response = &core_type.ResponseWithPagingData{
			Code:     params.StatusCode,
			Status:   status,
			Message:  params.Message,
			Paginate: params.Paginate,
			Data:     params.Data,
		}
	} else {
		response = &core_type.ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	c.JSON(params.StatusCode, response)
}

package helper

func (params ResponseParams) Response() any {
	var response any
	status := "failed"

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = "success"
	}

	if params.Data != nil {
		response = &ResponseWithoutPagingData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	return response
}

func (params ResponsePagingParams) Response() any {
	var response any
	status := "failed"

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = "success"
	}

	if params.Data != nil {
		response = &ResponseWithPagingData{
			Code:     params.StatusCode,
			Status:   status,
			Message:  params.Message,
			Paginate: params.Paginate,
			Data:     params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	return response
}

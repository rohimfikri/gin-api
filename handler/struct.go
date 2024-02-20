package handler

type NotFoundError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

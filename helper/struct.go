package helper

import "github.com/golang-jwt/jwt/v5"

type Paginate struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type ResponseParams struct {
	StatusCode int
	Message    string
	Data       any
}

type ResponsePagingParams struct {
	StatusCode int
	Message    string
	Paginate   *Paginate
	Data       any
}

type ResponseWithPagingData struct {
	Code     int       `json:"code"`
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Paginate *Paginate `json:"paginate,omitempty"`
	Data     any       `json:"data"`
}

type ResponseWithoutPagingData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type JWTUserClaims struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}

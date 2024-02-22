package public_type

type UserRegisterRequest struct {
	Username        string `json:"username" form:"username" binding:"required" validate:"required,alphanum"`
	Password        string `json:"password" form:"password" binding:"required" validate:"required"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm" binding:"required" validate:"required,eqfield=Password"`
	Email           string `json:"email" form:"email" binding:"required" validate:"required,email"`
	FirstName       string `json:"first_name" form:"first_name" binding:"required" validate:"required,alpha"`
	LastName        string `json:"last_name" form:"last_name" validate:"alpha"`
}

type ChangePasswordRequest struct {
	Username  string `json:"username" form:"username" binding:"required" validate:"required,alphanum"`
	Password  string `json:"password" form:"password" binding:"required" validate:"required"`
	SecretKey string `json:"secret_key" form:"secret_key" binding:"required" validate:"required"`
}

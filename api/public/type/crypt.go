package public_type

type GenerateBcryptRequest struct {
	Password string `json:"password" form:"password" binding:"required"`
}

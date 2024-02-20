package api_public

type GenerateBcryptRequest struct {
	Password string `json:"password" binding:"required"`
}

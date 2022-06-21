package reqmodel

type Refresh struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

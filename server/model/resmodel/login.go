package resmodel

type JwtTokens struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

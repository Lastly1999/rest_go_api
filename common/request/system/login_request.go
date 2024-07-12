package request

type SignRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}

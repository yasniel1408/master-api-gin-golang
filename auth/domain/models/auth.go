package auth_models

type Authentication struct {
	Token         string `json:"token"`
	Refresh_Token string `json:"refresh_token"`
}

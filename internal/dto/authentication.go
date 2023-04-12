package dto

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

package user

type LoginResponse struct {
	HP    string `json:"hp"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

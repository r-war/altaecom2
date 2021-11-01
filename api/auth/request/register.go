package request

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	RePassword string `json:"re_password"`
	Address string `json:"address"`
}


type RegisterAdminRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	RePassword string `json:"re_password"`
}

package auth

import (
	"AltaEcom/api/auth/request"
	"AltaEcom/business/user"
)

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	Login(username string, isAdmin bool) (*user.User, error)
	RegisterAdmin(request request.RegisterAdminRequest) (*request.RegisterAdminRequest, error)
	RegisterUser(request request.RegisterUserRequest) (*request.RegisterUserRequest, error)
}

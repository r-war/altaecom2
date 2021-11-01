package response

import (
	"AltaEcom/business/user"
	"time"
)

//GetUserResponse Get user by ID response payload
type GetUserResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	ModifiedAt time.Time `json:"modified_at"`
	Version    int       `json:"version"`
}

//NewGetUserResponse construct GetUserResponse
func NewGetUserResponse(user user.User) *GetUserResponse {
	var getUserResponse GetUserResponse

	getUserResponse.ID = user.ID
	getUserResponse.Name = user.Name
	getUserResponse.Username = user.Username
	getUserResponse.ModifiedAt = user.ModifiedAt
	getUserResponse.Version = user.Version

	return &getUserResponse
}

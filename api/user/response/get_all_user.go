package response

import (
	"AltaEcom/api/paginator"
	"AltaEcom/business/user"
)

type getAllUserResponse struct {
	Users []GetUserResponse `json:"users"`
	Meta  paginator.Meta    `json:"meta"`
}

//NewGetAllUserResponse construct GetAllUserResponse
func NewGetAllUserResponse(users []user.User, page int, rowPerPage int) getAllUserResponse {

	var (
		lenUsers = len(users)
	)

	getAllUserResponse := getAllUserResponse{}
	getAllUserResponse.Meta.BuildMeta(lenUsers, page, rowPerPage)

	for index, value := range users {
		if index == getAllUserResponse.Meta.RowPerPage {
			continue
		}

		var getUserResponse GetUserResponse

		getUserResponse.ID = value.ID
		getUserResponse.Name = value.Name
		getUserResponse.Username = value.Username
		getUserResponse.ModifiedAt = value.ModifiedAt
		getUserResponse.Version = value.Version

		getAllUserResponse.Users = append(getAllUserResponse.Users, getUserResponse)
	}

	if getAllUserResponse.Users == nil {
		getAllUserResponse.Users = []GetUserResponse{}
	}

	return getAllUserResponse
}

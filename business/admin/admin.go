package admin

import "time"

//User product User that available to rent or sell
type Admin struct {
	ID         int
	Name       string
	Username   string
	Password   string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
}

//NewUser create new User
func NewAdmin(
	id int,
	name string,
	username string,
	password string,
	creator string,
	createdAt time.Time) Admin {

	return Admin{
		ID:         id,
		Name:       name,
		Username:   username,
		Password:   password,
		CreatedAt:  createdAt,
		CreatedBy:  creator,
		ModifiedAt: createdAt,
		ModifiedBy: creator,
	}
}

//ModifyUser update existing User data
func (oldData *Admin) ModifyAdmin(newName string, modifiedAt time.Time, updater string) Admin {
	return Admin{
		ID:         oldData.ID,
		Name:       newName,
		Username:   oldData.Username,
		Password:   oldData.Password,
		CreatedAt:  oldData.CreatedAt,
		CreatedBy:  oldData.CreatedBy,
		ModifiedAt: modifiedAt,
		ModifiedBy: updater,
	}
}



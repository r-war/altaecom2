package admin

import (
	"AltaEcom/business/admin"
	"time"

	"gorm.io/gorm"
)

//GormRepository The implementation of user.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type Admin struct {
	ID         int       `gorm:"id;primaryKey;autoIncrement"`
	Name       string    `gorm:"name"`
	Username   string    `gorm:"email;index:idx_email,unique"`
	Password   string    `gorm:"password"`
	CreatedAt  time.Time `gorm:"created_at"`
	CreatedBy  string    `gorm:"created_by"`
	ModifiedAt time.Time `gorm:"modified_at"`
	ModifiedBy string    `gorm:"modified_by"`
}

func newAdmin(user admin.Admin) *Admin {

	return &Admin{
		user.ID,
		user.Name,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.CreatedBy,
		user.ModifiedAt,
		user.ModifiedBy,
	}

}

func (col *Admin) ToUser() admin.Admin {
	var user admin.Admin

	user.ID = col.ID
	user.Name = col.Name
	user.Username = col.Username
	user.Password = col.Password
	user.CreatedAt = col.CreatedAt
	user.CreatedBy = col.CreatedBy
	user.ModifiedAt = col.ModifiedAt
	user.ModifiedBy = col.ModifiedBy

	return user
}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//FindUserByID If data not found will return nil without error
func (repo *GormRepository) FindUserByID(id int) (*admin.Admin, error) {

	var userData Admin

	err := repo.DB.First(&userData, id).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//FindUserByID If data not found will return nil without error
func (repo *GormRepository) FindUserByUsernameAndPassword(username string, password string) (*admin.Admin, error) {

	var userData Admin

	err := repo.DB.Where("username = ?", username).First(&userData).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
func (repo *GormRepository) FindAllUser(skip int, rowPerPage int) ([]admin.Admin, error) {

	var users []Admin

	err := repo.DB.Offset(skip).Limit(rowPerPage).Find(&users).Error
	if err != nil {
		return nil, err
	}

	var result []admin.Admin
	for _, value := range users {
		result = append(result, value.ToUser())
	}

	return result, nil
}

//InsertUser Insert new User into storage
func (repo *GormRepository) InsertUser(user admin.Admin) error {

	userData := newAdmin(user)
	userData.ID = 0

	err := repo.DB.Create(userData).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateItem Update existing item in database
func (repo *GormRepository) UpdateUser(user admin.Admin, currentVersion int) error {

	userData := newAdmin(user)

	err := repo.DB.Model(&userData).Updates(Admin{Name: userData.Name}).Error
	if err != nil {
		return err
	}

	return nil
}

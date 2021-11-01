package category

import (
	"AltaEcom/business"
	"AltaEcom/business/admin"
	"time"
)

type CategorySpec struct {
	adminID int
	Name    string
}

type service struct {
	adminService admin.Service
	repo         Repository
}

func NewService(admin admin.Service, repo Repository) Service {
	return &service{admin, repo}
}

func (s *service) GetCategories() (*[]Category, error) {
	return s.repo.GetCategories()
}

func (s *service) FindCategoryById(id uint) (*Category, error) {
	return s.repo.FindCategoryById(id)
}

func (s *service) InsertCategory(cat *CategorySpec) error {
	_, err := s.adminService.FindUserByID(cat.adminID)
	if err != nil {
		return business.ErrNotHavePermission
	}
	dataCat := NewCategory(cat.Name, time.Now(), time.Now())

	return s.repo.InsertCategory(dataCat)
}

func (s *service) UpdateCategory(id uint, cat *CategorySpec) error {
	_, err := s.adminService.FindUserByID(cat.adminID)
	if err != nil {
		return business.ErrNotHavePermission
	}

	dataCat := ModifyCategory(cat.Name, time.Now())

	return s.repo.UpdateCategory(id, dataCat)

}

func (s *service) DeleteCategory(id uint, adminID int) error {
	_, err := s.adminService.FindUserByID(adminID)
	if err != nil {
		return business.ErrNotHavePermission
	}

	return s.repo.DeleteCategory(id)
}

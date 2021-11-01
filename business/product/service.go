package product

import (
	"AltaEcom/business"
	"AltaEcom/business/admin"
	"AltaEcom/util/validator"
	"time"
)

type ProductSpec struct {
	Name        string
	Description string
	Price       int
	Qty         int
	Image       string
	CategoryID  int
	adminID     int
}
type service struct {
	adminService admin.Service
	repository   Repository
}

func NewService(adminService admin.Service, repository Repository) Service {
	return &service{adminService, repository}
}

func (s *service) GetProductsByCategoryID(CategoryID int) ([]Product, error) {
	product, err := s.repository.GetProductsByCategoryID(CategoryID)
	if err != nil {
		return []Product{}, err
	}

	return product, err
}

func (s *service) GetProducts() ([]Product, error) {
	return s.repository.GetProducts()
}

func (s *service) FindProductByid(id int) (*Product, error) {
	return s.repository.FindProductByid(id)
}

func (s *service) InsertProduct(product *ProductSpec) error {
	if err := validator.GetValidator().Struct(product); err != nil {
		return business.ErrInvalidSpec
	}
	_, err := s.adminService.FindUserByID(product.adminID)
	if err != nil {
		return business.ErrNotHavePermission
	}

	data := NewProduct(
		product.CategoryID,
		product.Name,
		product.Price,
		product.Qty,
		product.Description,
		product.Image,
		time.Now(),
		time.Now())

	return s.repository.InsertProduct(data)
}

func (s *service) UpdateProduct(id int, updateProduct *ProductSpec) error {
	product, err := s.repository.FindProductByid(id)
	if err != nil {
		return err
	} else if product != nil {
		return business.ErrNotFound
	} else if product.DeletedAt.IsZero() {
		return business.ErrDeleted
	}

	if err := validator.GetValidator().Struct(product); err != nil {
		return business.ErrInvalidSpec
	}

	_, e := s.adminService.FindUserByID(updateProduct.adminID)
	if e != nil {
		return business.ErrNotHavePermission
	}
	dataProduct := ModifyProduct(
		updateProduct.Name,
		updateProduct.Price,
		updateProduct.CategoryID,
		updateProduct.Qty,
		updateProduct.Description,
		updateProduct.Image,
		time.Now(),
	)
	return s.repository.UpdateProduct(id, dataProduct)
}

func (s *service) DeleteProduct(id int, adminID int) error {
	if _, err := s.adminService.FindUserByID(adminID); err != nil {
		return business.ErrNotHavePermission
	}
	return s.repository.DeleteProduct(id)
}

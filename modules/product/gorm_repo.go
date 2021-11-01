package product

import (
	"AltaEcom/business/product"
	"AltaEcom/modules/category"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Product struct {
	ID          int               `gorm:"id;primaryKey;autoIncrement"`
	CategoryId  int               `gorm:"category_id"`
	Name        string            `gorm:"name"`
	Price       int               `gorm:"price"`
	Qty         int               `gorm:"qty"`
	Description string            `gorm:"description"`
	Image       string            `gorm:"image"`
	CreatedAt   time.Time         `gorm:"created_at"`
	UpdatedAt   time.Time         `gorm:"updated_at"`
	DeletedAt   time.Time         `gorm:"DeletedAt"`
	Category    category.Category `gorm:"foreignkey:CategoryId"`
}

func newProductTable(product product.Product) *Product {
	return &Product{
		product.ID,
		product.CategoryId,
		product.Name,
		product.Price,
		product.Qty,
		product.Description,
		product.Image,
		product.CreatedAt,
		product.UpdatedAt,
		product.DeletedAt,
		category.Category{},
	}
}

func (col *Product) ToProduct() product.Product {
	return product.Product{
		ID:          col.ID,
		CategoryId:  col.CategoryId,
		Name:        col.Name,
		Price:       col.Price,
		Qty:         col.Qty,
		Description: col.Description,
		Image:       col.Image,
		CreatedAt:   col.CreatedAt,
		UpdatedAt:   col.UpdatedAt,
		DeletedAt:   col.DeletedAt,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) GetProductsByCategoryID(CategoryID int) ([]product.Product, error) {
	var products []Product

	err := repo.DB.Preload("Categories").Joins("inner join categories on products.category_id = category.id").Where("category_id = ?", CategoryID).Find(&products).Error

	if err != nil {
		return nil, err
	}

	var result []product.Product
	var temp product.Product
	for _, val := range products {
		temp = val.ToProduct()
		temp.CategoryName = val.Category.Name
		result = append(result, temp)
	}

	return result, nil
}

func (repo *Repository) GetProducts() ([]product.Product, error) {
	var products []Product

	err := repo.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	var result []product.Product
	for _, val := range products {
		result = append(result, val.ToProduct())
	}
	return result, nil
}

func (repo *Repository) FindProductByid(id int) (*product.Product, error) {
	var product Product

	err := repo.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	result := product.ToProduct()
	return &result, nil
}

func (repo *Repository) InsertProduct(product product.Product) error {
	productData := newProductTable(product)

	if err := repo.DB.Create(productData).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) UpdateProduct(id int, product product.Product) error {
	var productData Product
	err := repo.DB.Find(&productData, "id = ?", id).Error

	if err != nil {
		return err
	}
	err = repo.DB.Model(&productData).Updates(Product{
		Name:        product.Name,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Qty:         product.Qty,
		Description: product.Description,
		UpdatedAt:   time.Now(),
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) DeleteProduct(id int) error {
	var productData Product

	err := repo.DB.First(productData, "id = ?", id).Error
	if err != nil {
		return err
	}
	err = repo.DB.Model(productData).Updates(Product{DeletedAt: time.Now()}).Error
	if err != nil {
		return err
	}

	return nil
}

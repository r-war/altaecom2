package category

import (
	"AltaEcom/business/category"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type Category struct {
	ID        uint   `gorm:"id;primaryKey;autoIncrement"`
	Name      string `gorm:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewCategoryTable(category category.Category) *Category {
	return &Category{
		category.ID,
		category.Name,
		category.CreatedAt,
		category.UpdatedAt,
		category.DeletedAt,
	}
}

func (col *Category) ToCategory() category.Category {
	var category category.Category

	category.ID = col.ID
	category.Name = col.Name
	category.CreatedAt = col.CreatedAt
	category.UpdatedAt = col.UpdatedAt
	category.DeletedAt = col.DeletedAt

	return category
}

func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db}
}

func (r *GormRepository) GetCategories() (*[]category.Category, error) {
	var cat []Category

	err := r.DB.Find(&cat).Error

	if err != nil {
		return nil, err
	}

	var result []category.Category
	for _, val := range cat {
		result = append(result, val.ToCategory())
	}

	return &result, nil
}

func (r *GormRepository) FindCategoryById(id uint) (*category.Category, error) {
	var data Category

	err := r.DB.First(&data, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	category := data.ToCategory()

	return &category, nil
}

func (r *GormRepository) InsertCategory(cat category.Category) error {
	catData := NewCategoryTable(cat)

	if err := r.DB.Create(catData).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormRepository) UpdateCategory(id uint, cat category.Category) error {
	var category Category
	err := r.DB.Find(&category, "id = ?", id).Error
	if err != nil {
		return err
	}
	err = r.DB.Model(&category).Updates(Category{Name: cat.Name, UpdatedAt: cat.UpdatedAt}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) DeleteCategory(id uint) error {
	var category Category

	err := r.DB.First(&category, "id = ?", id).Error

	if err != nil {
		return err
	}

	err = r.DB.Model(&category).Updates(Category{DeletedAt: time.Now()}).Error
	if err != nil {
		return err
	}

	return nil
}

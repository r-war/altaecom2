package category

type Service interface {
	GetCategories() (*[]Category, error)

	FindCategoryById(ID uint) (*Category, error)

	UpdateCategory(ID uint, category *CategorySpec) error

	InsertCategory(category *CategorySpec) error

	DeleteCategory(ID uint, adminId int) error
}

type Repository interface {
	GetCategories() (*[]Category, error)

	FindCategoryById(ID uint) (*Category, error)

	UpdateCategory(ID uint, category Category) error

	InsertCategory(category Category) error

	DeleteCategory(ID uint) error
}

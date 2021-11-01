package product

type Service interface {
	GetProductsByCategoryID(CategoryID int) ([]Product, error)

	GetProducts() ([]Product, error)

	FindProductByid(Id int) (*Product, error)

	InsertProduct(product *ProductSpec) error

	UpdateProduct(id int, product *ProductSpec) error

	DeleteProduct(id int, adminID int) error
}

type Repository interface {
	GetProductsByCategoryID(CategoryID int) ([]Product, error)

	GetProducts() ([]Product, error)

	FindProductByid(Id int) (*Product, error)

	InsertProduct(product Product) error

	UpdateProduct(id int, product Product) error

	DeleteProduct(id int) error
}

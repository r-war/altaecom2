package category

import "time"

type Category struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewCategory(
	Name string,
	CreatedAt time.Time,
	UpdatedAt time.Time,
) Category {
	return Category{
		Name:      Name,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}

func ModifyCategory(Name string, updatedAt time.Time) Category {
	return Category{
		Name:      Name,
		UpdatedAt: updatedAt,
	}
}

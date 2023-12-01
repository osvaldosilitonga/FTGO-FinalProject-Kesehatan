package web

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int64  `json:"price" validate:"required,min=1000"`
	Stock       int32  `json:"stock" validate:"required,min=1"`
}

type UpdateProductRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	Price       int64  `json:"price,omitempty" validate:"omitempty,min=1000"`
	Stock       int32  `json:"stock,omitempty" validate:"omitempty,min=1"`
}

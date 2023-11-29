package web

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int64  `json:"price" validate:"required,min=1000"`
	Stock       int32  `json:"stock" validate:"required,min=1"`
}

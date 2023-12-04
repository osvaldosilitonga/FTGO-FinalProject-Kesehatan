package web

type ProductDatas struct {
	Id       string `json:"id" validate:"required"`
	Quantity int    `json:"qty" validate:"required,min=1"`
}

type CreateOrderProductRequest struct {
	Products []ProductDatas `json:"products" validate:"required"`
}

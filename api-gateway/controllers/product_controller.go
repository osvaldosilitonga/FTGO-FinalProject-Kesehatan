package controllers

import (
	pb "gateway/internal/product"
	"gateway/models/web"
	"gateway/service"
	"gateway/utils"

	"github.com/labstack/echo/v4"
)

type Product interface {
	ListProduct(c echo.Context) error
	CreateProduct(c echo.Context) error
	FindByID(c echo.Context) error
}

type ProductImpl struct {
	ProductService service.Product
}

func NewProductController(ps service.Product) *ProductImpl {
	return &ProductImpl{
		ProductService: ps,
	}
}

func (p *ProductImpl) ListProduct(c echo.Context) error {
	list, err := p.ProductService.ListProduct(c.Request().Context())
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, list)
}

func (p *ProductImpl) CreateProduct(c echo.Context) error {
	req := &web.CreateProductRequest{}
	if err := c.Bind(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}

	product := &pb.CreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	res, err := p.ProductService.CreateProduct(c.Request().Context(), product)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, res)
}

func (p *ProductImpl) FindByID(c echo.Context) error {
	id := c.Param("id")

	req := &pb.GetProductRequest{
		Id: id,
	}

	res, err := p.ProductService.GetProduct(c.Request().Context(), req)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

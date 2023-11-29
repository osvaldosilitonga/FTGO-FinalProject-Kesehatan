package controllers

import (
	"errors"
	pb "gateway/internal/product"
	"gateway/models/web"
	"gateway/service"
	"gateway/utils"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Product interface {
	ListProduct(c echo.Context) error
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
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
			case codes.Internal:
				return utils.ErrorMessage(c, &utils.ApiInternalServer, "internal server error")
			default:
				return utils.ErrorMessage(c, &utils.ApiInternalServer, "Failed get list product")
			}
		}
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
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
			case codes.Internal:
				return utils.ErrorMessage(c, &utils.ApiInternalServer, errors.New("Internal server error"))
			default:
				return utils.ErrorMessage(c, &utils.ApiInternalServer, "Failed get list product")
			}
		}
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, res)
}

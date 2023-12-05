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

// @Summary 	List Product
// @Description List Product
// @Tags 			Product
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} web.SwProductList
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/products [get]
func (p *ProductImpl) ListProduct(c echo.Context) error {
	list, err := p.ProductService.ListProduct(c.Request().Context())
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, list)
}

// @Summary 	Create Product (Admin Only)
// @Description Create new product
// @Tags 			Product
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 		data body web.CreateProductRequest true "Product Data"
// @Success 	201 {object} web.SwProductCreate
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/products [post]
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

// @Summary 	Product Detail
// @Description Get product by ID
// @Tags 			Product
// @Accept 		json
// @Produce 	json
// @Param 			id path integer true "Product ID"
// @Success 	200 {object} web.SwProductFindById
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/products/{id} [get]
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

// @Summary 	Update Product (Admin Only)
// @Description Update product by ID
// @Tags 			Product
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path integer true "Product ID"
// @Param 		data body pb.UpdateProductRequest true "Product Data"
// @Success 	200 {object} web.SwProductUpdate
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/products/{id} [put]
func (p *ProductImpl) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	req := &web.UpdateProductRequest{}
	if err := c.Bind(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}

	update := &pb.UpdateProductRequest{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	res, err := p.ProductService.UpdateProduct(c.Request().Context(), update)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

// @Summary 	Delete Product (Admin Only)
// @Description Delete product by ID
// @Tags 			Product
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path integer true "Product ID"
// @Success 	200 {object} web.SwProductDelete
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/products/{id} [delete]
func (p *ProductImpl) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	req := &pb.DeleteProductRequest{
		Id: id,
	}

	res, err := p.ProductService.DeleteProduct(c.Request().Context(), req)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

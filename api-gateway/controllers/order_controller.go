package controllers

import (
	"fmt"
	"gateway/internal/order"
	"gateway/models/web"
	"gateway/service"
	"gateway/utils"
	"strconv"
	"strings"

	pb "gateway/internal/order"
	pbProduct "gateway/internal/product"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController interface {
	CreateOrderProduct(c echo.Context) error
	CancelOrder(c echo.Context) error
	ListOrder(c echo.Context) error
	OrderDetail(c echo.Context) error
	ConfirmOrder(c echo.Context) error
	FindByUserId(c echo.Context) error
}

type OrderControllerImpl struct {
	OrderService   service.Order
	PaymentService service.Payment
	ProductService service.Product
}

func NewOrderController(so service.Order, ps service.Payment, prs service.Product) OrderController {
	return &OrderControllerImpl{
		OrderService:   so,
		PaymentService: ps,
		ProductService: prs,
	}
}

// @Summary 	Create order (Customer Only)
// @Description Create a new order for customer. Need to login first
// @Tags 			Order
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 		createOrder body web.CreateOrderProductRequest true "Order details"
// @Success 	201 {object} web.SwCreateOrder
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/order [post]
func (o *OrderControllerImpl) CreateOrderProduct(c echo.Context) error {
	id := c.Get("id").(int)

	req := &web.CreateOrderProductRequest{}
	if err := c.Bind(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}

	products := []*pb.Product{}
	for _, product := range req.Products {
		p := &order.Product{}
		p.Id = product.Id
		p.Qty = int32(product.Quantity)

		products = append(products, p)
	}

	order := &pb.CreateOrderProductRequest{
		User: &pb.User{
			Id:    int32(id),
			Email: c.Get("email").(string),
		},
		Products: products,
	}

	res, err := o.OrderService.CreateOrderProduct(c.Request().Context(), order)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	orderId, _ := primitive.ObjectIDFromHex(res.OrderId)

	paymentRequest := &web.CreatePaymentRequest{
		OrderID:     orderId,
		UserID:      id,
		Email:       c.Get("email").(string),
		Amount:      int(res.TotalAmount),
		Description: "Payment for product orders",
	}

	// Create payment
	resp, code, err := o.PaymentService.CreatePayment(paymentRequest)
	for err != nil || code != 201 {
		fmt.Println("Payment service error, retrying...")

		resp, code, err = o.PaymentService.CreatePayment(paymentRequest)
	}

	// Update product stock
	updateReq := &pbProduct.UpdateStockRequest{
		Type: "decrease",
	}
	for _, product := range res.Products {
		d := &pbProduct.Data{
			Id:       product.Id,
			Quantity: product.Qty,
		}

		updateReq.Datas = append(updateReq.Datas, d)
	}

	_, err = o.ProductService.UpdateStock(c.Request().Context(), updateReq)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, resp)
}

// @Summary 	Cancel order (Customer Only)
// @Description Canceling customer order. Need to login first
// @Tags 			Order
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path string true "Order ID"
// @Success 	200 {object} web.SwCancelOrder
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/order/cancel/{id} [post]
func (o *OrderControllerImpl) CancelOrder(c echo.Context) error {
	userId := c.Get("id").(int)
	orderId := c.Param("id")

	req := &pb.FindByOrderIdRequest{
		OrderId: orderId,
	}

	// Find order by order id
	res, err := o.OrderService.FindByOrderId(c.Request().Context(), req)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	// Check if user is the owner of the order
	if int(res.UserId) != userId {
		return utils.ErrorMessage(c, &utils.ApiForbidden, nil)
	}

	// Cancel order
	res, err = o.OrderService.UpdateStatus(c.Request().Context(), &pb.UpdateOrderStatusRequest{
		OrderId: orderId,
		Status:  "CANCEL",
	})
	if err != nil {
		return utils.GrpcError(c, err)
	}

	// Update product stock
	updateReq := &pbProduct.UpdateStockRequest{
		Type: "increase",
	}
	for _, product := range res.Products {
		d := &pbProduct.Data{
			Id:       product.Id,
			Quantity: product.Qty,
		}

		updateReq.Datas = append(updateReq.Datas, d)
	}

	_, err = o.ProductService.UpdateStock(c.Request().Context(), updateReq)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

// @Summary 	List all order (Admin Only)
// @Description List of all customer orders
// @Tags 			Order
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param        page    query     integer  false  "Page"
// @Param        status    query     string  false  "status"
// @Success 	200 {object} web.SwOrderDetail
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/order/admin [get]
func (o *OrderControllerImpl) ListOrder(c echo.Context) error {
	status := strings.ToUpper(c.QueryParam("status"))
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}

	req := &pb.ListOrderRequest{
		Status: status,
		Page:   int32(pageInt),
	}

	res, err := o.OrderService.ListOrder(c.Request().Context(), req)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	if len(res.Orders) == 0 {
		return utils.ErrorMessage(c, &utils.ApiNotFound, nil)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

// @Summary 	Order Detail (Customer and Admin)
// @Description Get order detail by id. Need to login first
// @Tags 			Order
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path string true "Order ID"
// @Success 	200 {object} web.SwOrderDetail
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/order/{id} [get]
func (o *OrderControllerImpl) OrderDetail(c echo.Context) error {
	orderId := c.Param("id")
	userId := c.Get("id").(int)
	userRole := c.Get("role").(string)

	req := &pb.FindByOrderIdRequest{
		OrderId: orderId,
	}

	res, err := o.OrderService.FindByOrderId(c.Request().Context(), req)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	if int(res.UserId) != userId && userRole != "admin" {
		return utils.ErrorMessage(c, &utils.ApiForbidden, nil)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

// @Summary 	Confirm Order (Admin Only)
// @Description Confirm the order to be completed.
// @Tags 			Order
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path string true "Order ID"
// @Success 	200 {object} web.SwOrderConfirm
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/order/admin/confirm/{id} [put]
func (o *OrderControllerImpl) ConfirmOrder(c echo.Context) error {
	orderId := c.Param("id")

	req := &pb.UpdateOrderStatusRequest{
		OrderId: orderId,
		Status:  "SUCCESS",
	}

	res, err := o.OrderService.UpdateStatus(c.Request().Context(), req)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

// @Summary 	Find Order by User ID (User)
// @Description Find order by user id
// @Tags 			Order
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path string true "Order ID"
// @Success 	200 {object} web.SwOrderFindByID
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	404 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/order/user/{id} [get]
func (o *OrderControllerImpl) FindByUserId(c echo.Context) error {
	userId := c.Get("id").(int)

	paramId := c.Param("id")
	pId, error := strconv.Atoi(paramId)
	if error != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, error)
	}

	if userId != pId {
		return utils.ErrorMessage(c, &utils.ApiForbidden, nil)
	}

	status := strings.ToUpper(c.QueryParam("status"))
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}

	req := &pb.FindByUserIdRequest{
		UserId: int32(pId),
		Page:   int32(pageInt),
		Status: status,
	}

	res, err := o.OrderService.FindByUserId(c.Request().Context(), req)
	if err != nil {
		return utils.GrpcError(c, err)
	}

	if len(res.Orders) == 0 {
		return utils.ErrorMessage(c, &utils.ApiNotFound, nil)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

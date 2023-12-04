package controllers

import (
	"fmt"
	"gateway/internal/order"
	"gateway/models/web"
	"gateway/service"
	"gateway/utils"

	pb "gateway/internal/order"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController interface {
	CreateOrderProduct(c echo.Context) error
}

type OrderControllerImpl struct {
	OrderService   service.Order
	PaymentService service.Payment
}

func NewOrderController(so service.Order, ps service.Payment) OrderController {
	return &OrderControllerImpl{
		OrderService:   so,
		PaymentService: ps,
	}
}

func (o *OrderControllerImpl) CreateOrderProduct(c echo.Context) error {
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
			Id:    1,
			Email: "test",
			Role:  "user",
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
		UserID:      int(res.UserId),
		Email:       "jacksparrow257257@gmail.com",
		Amount:      int(res.TotalAmount),
		Description: "Payment for product order",
	}

	resp, code, err := o.PaymentService.CreatePayment(paymentRequest)
	for err != nil || code != 201 {
		fmt.Println("Payment service error, retrying...")
		fmt.Println(err, "<---------------- err")
		resp, code, err = o.PaymentService.CreatePayment(paymentRequest)
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, resp)
}

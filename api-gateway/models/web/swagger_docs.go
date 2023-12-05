package web

type SwCreateOrder struct {
	Code   int                     `json:"code"`
	Status string                  `json:"status"`
	Data   CreateOrderResponseData `json:"data"`
}

type CreateOrderResponseData struct {
	InvoiceUrl   string `json:"invoice_url"`
	InvoiceId    string `json:"invoice_id"`
	Status       string `json:"status"`
	Description  string `json:"description"`
	Created      string `json:"created"`
	ExpairyDate  string `json:"expairy_date"`
	ExternalId   string `json:"external_id"`
	PaymentEmail string `json:"payment_email"`
	Amount       int    `json:"amount"`
}

// ---------------------------------------------

type SwCancelOrder struct {
	Code   int                     `json:"code"`
	Status string                  `json:"status"`
	Data   CreateOrderResponseData `json:"data"`
}

type SwCancelOrderData struct {
	OrderId     string                 `json:"order_id"`
	UserId      int                    `json:"user_id"`
	Type        string                 `json:"type"`
	TotalAmount int                    `json:"total_amount"`
	Status      string                 `json:"status"`
	Products    []SwCancelOrderProduct `json:"products"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

type SwCancelOrderProduct struct {
	Id       string `json:"id"`
	Quantity int    `json:"qty"`
}

// -----------------------------------------------
type SwOrderDetail struct {
	Code   int                     `json:"code"`
	Status string                  `json:"status"`
	Data   CreateOrderResponseData `json:"data"`
}

type SwOrderConfirm struct {
	Code   int                     `json:"code"`
	Status string                  `json:"status"`
	Data   CreateOrderResponseData `json:"data"`
}

// -----------------------------------------------
type SwUserLogin struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   SwUserLoginData `json:"data"`
}

type SwUserLoginData struct {
	Token string `json:"token"`
}

// -----------------------------------------------
type SwUserRegister struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Data   SwUserRegisterData `json:"data"`
}
type SwUserRegisterData struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// -----------------------------------------------
type SwUserProfile struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   SwUserProfileData `json:"data"`
}
type SwUserProfileData struct {
	UserID    int    `json:"user_id"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Name      string `json:"name"`
}

// -----------------------------------------------
type SwUserProfileUpdate struct {
	Code   int                     `json:"code"`
	Status string                  `json:"status"`
	Data   SwUserProfileDataUpdate `json:"data"`
}
type SwUserProfileDataUpdate struct {
	UserProfile string `json:"user_profile"`
}

type UserProfileData struct {
	UserID    int    `json:"user_id"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
}

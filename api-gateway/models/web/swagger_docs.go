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

// -----------------------------------------------
type SwProductList struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   SwProductListData `json:"data"`
}
type SwProductListData struct {
	Products []SwProduct `json:"products"`
}

type SwProduct struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// -----------------------------------------------
type SwProductFindById struct {
	Code   int                   `json:"code"`
	Status string                `json:"status"`
	Data   SwProductFindByIdData `json:"data"`
}
type SwProductFindByIdData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// -----------------------------------------------

type SwProductUpdate struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Data   SwProductUpdateData `json:"data"`
}
type SwProductUpdateData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// -----------------------------------------------
type SwProductDelete struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Data   SwProductDeleteData `json:"data"`
}
type SwProductDeleteData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// -----------------------------------------------
type SwProductCreate struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Data   SwProductCreateData `json:"data"`
}
type SwProductCreateData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// ------------------ PAYMENT --------------------------
type SwPaymentInvoice struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Data   SwPaymentInvoiceData `json:"data"`
}
type SwPaymentInvoiceData struct {
	Id            string `json:"id"`
	InvoiceId     string `json:"invoice_id"`
	OrderId       string `json:"order_id"`
	UserID        int    `json:"user_id"`
	Email         string `json:"email"`
	Amount        int    `json:"amount"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	MerchantName  string `json:"merchant_name"`
	Currency      string `json:"currency"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type SwPaymentByUserID struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Data   []SwPaymentInvoiceData `json:"data"`
}

// ------------------ ORDER --------------------------
type SwOrderFindByID struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Data   SwOrderFindByIDData `json:"data"`
}

type SwOrderFindByIDData struct {
	OrderId     string                   `json:"order_id"`
	UserId      int                      `json:"user_id"`
	Type        string                   `json:"type"`
	TotalAmount int                      `json:"total_amount"`
	Status      string                   `json:"status"`
	Products    []SwOrderFindByIDProduct `json:"products"`
	CreatedAt   string                   `json:"created_at"`
	UpdatedAt   string                   `json:"updated_at"`
}

type SwOrderFindByIDProduct struct {
	Id       string `json:"id"`
	Quantity int    `json:"qty"`
}

package dto

import "time"

type CreateOrder struct {
	OrderName string    `json:"name"`
	Product   []Product `json:"product"`
}
type Product struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CreateOrderHTTPResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    CreateOrderResponse `json:"data"`
}

type CreateOrderResponse struct {
	ID string `json:"id"`
}

type GetOrderHTTPResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetOrderResponse struct {
	ID          string                 `json:"id"`
	OrderName   string                 `json:"name"`
	OrderValue  float64                `json:"ordervalue"`
	OrderStatus string                 `json:"orderstatus"`
	Product     []OrderProductResponse `json:"products"`
}

type GetOrderWithDispatchResponse struct {
	ID           string                 `json:"id"`
	OrderName    string                 `json:"name"`
	OrderValue   float64                `json:"ordervalue"`
	OrderStatus  string                 `json:"orderstatus"`
	DispatchDate time.Time              `json:"dispatch"`
	Product      []OrderProductResponse `json:"products"`
}

type OrderProductResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type UpdateOrderHTTPResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    UpdateOrders `json:"data"`
}

type UpdateOrders struct {
	ID           string    `json:"id"`
	OrderName    string    `json:"name"`
	OrderValue   float64   `json:"ordervalue"`
	OrderStatus  string    `json:"orderstatus"`
	DispatchDate time.Time `json:"dispatch"`
	Product      []Product `json:"products"`
}

type UpdateOrder struct {
	ID          string `json:"id"`
	OrderStatus string `json:"orderstatus"`
}

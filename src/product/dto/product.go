package dto

type CreateProduct struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Availability bool    `json:"availability"`
}

type ProductHTTPResponse struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    CreateProductResponse `json:"data"`
}

type CreateProductResponse struct {
	ID string `json:"id"`
}

type UpdateProductHTTPResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    UpdateProduct `json:"data"`
}

type UpdateProduct struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Availability bool    `json:"availability"`
}

type GetProductHTTPResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    GetProductResponse `json:"data"`
}

type ProductListHTTPResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    []GetProductResponse `json:"data"`
}

type GetProductResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Availability bool    `json:"availability"`
}

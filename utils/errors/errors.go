package errors

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Err struct {
	INTERNAL_ERR  Error
	INVALID_ERR   Invalid_Error
	OBJ_NOT_FOUND Error
}

// custom error types
type Invalid_Error struct {
	PRODUCT_ID   Error
	PRODUCTNAME  Error
	PRICE        Error
	AVAILABILITY Error
	CATEGORY     Error
	QUANTITY     Error
	ORDER_ID     Error
	ORDERSTATUS  Error
}

// common error response
type InvalidErrorResponse struct {
	Code        int             `json:"code"`
	Message     string          `json:"message"`
	Description string          `json:"description"`
	Errors      []InvalidErrors `json:"errors"`
}

type CommonErrorResponse struct {
	Code        int                      `json:"code"`
	Message     string                   `json:"message"`
	Description string                   `json:"description"`
	Errors      []map[string]interface{} `json:"errors"`
}

type InvalidErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code    int
	Type    string
	Field   string
	Message string
}

// custom error handler
func ErrorHandler(ctx *gin.Context, err interface{}) {

	if e, ok := err.(*Error); ok {
		errResponse := CommonErrorResponse{
			Code:        e.Code,
			Message:     e.Type,
			Description: e.Message,
			Errors:      make([]map[string]interface{}, 0),
		}

		ctx.JSON(errResponse.Code, errResponse)
		var m map[string]interface{}
		json.NewDecoder(ctx.Request.Body).Decode(&m)
		return

	} else if er, ok := err.([]*Error); ok {
		errResponse := InvalidErrorResponse{}
		errResponse.Code = 400
		errResponse.Message = "invalid_request_error"
		errResponse.Description = "The request was unacceptable, due to missing a required parameter or invalid parameter."
		var invalid_fields []InvalidErrors

		for _, e := range er {
			invalid := InvalidErrors{
				Field:   e.Field,
				Message: e.Message,
			}
			invalid_fields = append(invalid_fields, invalid)
		}
		errResponse.Errors = append(errResponse.Errors, invalid_fields...)
		ctx.JSON(errResponse.Code, errResponse)
		var m map[string]interface{}
		json.NewDecoder(ctx.Request.Body).Decode(&m)

		return
	}

	errResponse := CommonErrorResponse{}
	errResponse.Code = 500
	errResponse.Message = "internal_server_error"
	errResponse.Description = "Internal server error."

	ctx.JSON(errResponse.Code, errResponse)
	var m map[string]interface{}
	json.NewDecoder(ctx.Request.Body).Decode(&m)

}

// errors
func NewErr() Err {
	e := Err{
		INTERNAL_ERR: Error{
			Code:    500,
			Type:    "internal_server_error",
			Message: "Internal server error",
		},

		INVALID_ERR: Invalid_Error{

			PRODUCT_ID: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "id",
				Message: "This field should be a valid Product id",
			},
			PRODUCTNAME: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "name",
				Message: "This field should be valid string not greater than 255 characters",
			},
			CATEGORY: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "category",
				Message: "This field should be valid category Premium or Regular or Budget",
			},
			PRICE: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "price",
				Message: "This field should be valid float",
			},
			AVAILABILITY: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "availability",
				Message: "This field should be valid boolean",
			},
			ORDER_ID: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "id",
				Message: "This field should be a valid order id",
			},
			QUANTITY: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "quantity",
				Message: "This field should be valid quantity",
			},
			ORDERSTATUS: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "orderstatus",
				Message: "This field should be valid order status",
			},
		},

		OBJ_NOT_FOUND: Error{
			Code:    404,
			Type:    "object_not_found",
			Message: "The requested object does not exist or already deleted.",
		},
	}

	return e
}

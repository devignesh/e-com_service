package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"e-com/src/order/dto"
	"e-com/src/order/service"
	"e-com/src/order/validator"
	errorsutils "e-com/utils/errors"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

var Err = errorsutils.NewErr()
var orderService = new(service.OrderService)

// create order contrller
func (orderController *OrderController) CreateOrder(ctx *gin.Context) {

	var r map[string]interface{}
	var req dto.CreateOrder

	if err := ctx.BindJSON(&r); err != nil {
		errorsutils.ErrorHandler(ctx, &Err.INTERNAL_ERR)
		return
	}

	j, er := json.Marshal(r)
	if er != nil {
		fmt.Printf(er.Error())
	}
	json.Unmarshal(j, &req)

	errs := validator.CreateOrderValidator(ctx, req)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := orderService.CreateOrder(ctx, req)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// get order by Id controller
func (orderController *OrderController) GetOrderByID(ctx *gin.Context) {

	//order id
	id := ctx.Param("id")

	errs := validator.GetOrderValidator(ctx, id)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := orderService.GetOrderByID(ctx, id)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// update order controller
func (orderController *OrderController) UpdateOrder(ctx *gin.Context) {

	Order_id := ctx.Param("id")

	var req dto.UpdateOrder

	if err := ctx.BindJSON(&req); err != nil {
		errorsutils.ErrorHandler(ctx, &Err.INTERNAL_ERR)
		return
	}

	errs := validator.UpdateOrderValidator(ctx, Order_id, req)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := orderService.UpdateOrder(ctx, Order_id, req)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

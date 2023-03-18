package controller

import (
	"e-com/src/product/dto"
	"e-com/src/product/service"
	"e-com/src/product/validator"
	errorsutils "e-com/utils/errors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var productService = new(service.ProductService)
var Err = errorsutils.NewErr()

type ProductController struct{}

// product controller create product
func (productController *ProductController) CreateProduct(ctx *gin.Context) {

	var r map[string]interface{}
	var req dto.CreateProduct

	if err := ctx.BindJSON(&r); err != nil {
		errorsutils.ErrorHandler(ctx, &Err.INTERNAL_ERR)
		return
	}

	j, er := json.Marshal(r)
	if er != nil {
		fmt.Printf(er.Error())
	}
	json.Unmarshal(j, &req)

	errs := validator.CreateProductValidator(ctx, req)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := productService.CreateProduct(ctx, req)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// get product controller
func (productController *ProductController) GetProductByID(ctx *gin.Context) {

	//product id
	id := ctx.Param("id")

	errs := validator.GetProductValidator(ctx, id)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := productService.GetProductByID(ctx, id)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// get product list contrller
func (productController *ProductController) GetProductList(ctx *gin.Context) {

	response, e := productService.GetProductList(ctx)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// update produt
func (productController *ProductController) UpdateProduct(ctx *gin.Context) {

	product_id := ctx.Param("id")
	var req dto.UpdateProduct

	if err := ctx.BindJSON(&req); err != nil {
		errorsutils.ErrorHandler(ctx, &Err.INTERNAL_ERR)
		return
	}

	errs := validator.UpdateProductValidator(ctx, req)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := productService.UpdateProduct(ctx, product_id, req)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

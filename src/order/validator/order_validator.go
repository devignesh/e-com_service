package validator

import (
	"context"
	"e-com/src/order/dto"
	order "e-com/src/order/repository"
	product "e-com/src/product/repository"
	errorsutils "e-com/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Err = errorsutils.NewErr()

var ProductRepository = new(product.Product)

func validateStatus(status string) bool {

	switch status {
	case "dispatched":
		return true
	case "completed":
		return true
	case "cancelled":
		return true
	}

	return false
}

func CreateOrderValidator(ctx context.Context, req dto.CreateOrder) interface{} {

	var invalidErrs []*errorsutils.Error

	if len(req.OrderName) == 0 || len(req.OrderName) > 255 {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCTNAME)

	}

	for _, product := range req.Product {

		prod_id, err := primitive.ObjectIDFromHex(product.ProductID)
		if err != nil {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCT_ID)

		} else if prod_id != primitive.NilObjectID {
			product_data, err := ProductRepository.FindProductByID(ctx, prod_id)
			if err != nil {
				invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCT_ID)
			}
			if !product_data.Availability {
				invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCT_ID)
			}

		}
		if product.Quantity > 10 {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.QUANTITY)
		}
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}

func GetOrderValidator(ctx context.Context, id string) interface{} {

	var invalidErrs []*errorsutils.Error

	var orderrepo = new(order.Order)

	order_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.ORDER_ID)
		return invalidErrs
	}

	var filter bson.D
	filter = append(filter, primitive.E{Key: "_id", Value: order_id})
	filter = append(filter, primitive.E{Key: "is_deleted", Value: false})

	orders, err := orderrepo.FindOrderByBson(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.ORDER_ID)
			return invalidErrs
		}

		return &Err.INTERNAL_ERR
	}

	if orders.ID == primitive.NilObjectID {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.ORDER_ID)
		return invalidErrs
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}

func UpdateOrderValidator(ctx context.Context, req dto.UpdateOrder) interface{} {

	var invalidErrs []*errorsutils.Error

	var orderrepo = new(order.Order)

	order_id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.ORDER_ID)
		return invalidErrs
	}

	var filter bson.D
	filter = append(filter, primitive.E{Key: "_id", Value: order_id})
	filter = append(filter, primitive.E{Key: "is_deleted", Value: false})

	orders, err := orderrepo.FindOrderByBson(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.ORDER_ID)
			return invalidErrs
		}

		return &Err.INTERNAL_ERR
	}

	if orders.ID == primitive.NilObjectID {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.ORDER_ID)
		return invalidErrs
	}

	if valid := validateStatus(req.OrderStatus); !valid {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.ORDERSTATUS)
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}

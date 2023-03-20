package validator

import (
	"context"
	"e-com/src/product/dto"
	"e-com/src/product/repository"
	errorsutils "e-com/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Err = errorsutils.NewErr()

// common validators for product service
func validateCategory(category string) bool {
	switch category {
	case "premium":
		return true
	case "regular":
		return true
	case "budget":
		return true
	}

	return false
}

func CreateProductValidator(ctx context.Context, req dto.CreateProduct) interface{} {

	var invalidErrs []*errorsutils.Error

	if len(req.Name) == 0 || len(req.Name) > 255 {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCTNAME)

	}

	if !(req.Availability || !req.Availability) {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.AVAILABILITY)
	}

	if valid := validateCategory(req.Category); !valid {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.CATEGORY)
	}

	if req.Price == 0 {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRICE)
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}

func GetProductValidator(ctx context.Context, id string) interface{} {

	var invalidErrs []*errorsutils.Error

	var productrepo = new(repository.Product)

	product_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCT_ID)
		return invalidErrs
	}

	var filter bson.D
	filter = append(filter, primitive.E{Key: "_id", Value: product_id})
	filter = append(filter, primitive.E{Key: "is_deleted", Value: false})

	product, err := productrepo.FindProductByBson(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCT_ID)
			return invalidErrs
		}

		return &Err.INTERNAL_ERR
	}

	if product.ID == primitive.NilObjectID {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCT_ID)
		return invalidErrs
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}

func UpdateProductValidator(ctx context.Context, req dto.UpdateProduct) interface{} {

	var invalidErrs []*errorsutils.Error

	if len(req.Name) > 0 {
		if len(req.Name) == 0 || len(req.Name) > 255 {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRODUCTNAME)

		}
	}

	if req.Availability != true || req.Availability != false {
		if !(req.Availability || !req.Availability) {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.AVAILABILITY)
		}
	}

	if valid := validateCategory(req.Category); !valid {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.CATEGORY)
	}

	if req.Price != -0 {
		if req.Price == 0 {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.PRICE)
		}
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}

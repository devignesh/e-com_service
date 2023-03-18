package service

import (
	"context"
	"e-com/src/product/dto"
	"e-com/src/product/repository"
	errorutils "e-com/utils/errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ProductRepository = new(repository.Product)
var Err = errorutils.NewErr()

type ProductService struct{}

func (ps *ProductService) CreateProduct(c context.Context, req dto.CreateProduct) (dto.ProductHTTPResponse, *errorutils.Error) {

	var product repository.Product

	product.Name = req.Name
	product.Category = req.Category
	product.Availability = req.Availability
	product.Price = req.Price
	product.CreatedAt = time.Now().UTC()
	product.UpdatedAt = time.Now().UTC()

	ProductID, err := ProductRepository.CreateProduct(c, product)
	if err != nil {
		fmt.Printf("Error creating product: %+v", err.Error())
		return dto.ProductHTTPResponse{}, &Err.INTERNAL_ERR
	}

	response := dto.CreateProductResponse{
		ID: ProductID.(primitive.ObjectID).Hex(),
	}

	return dto.ProductHTTPResponse{
		Status:  201,
		Message: "Product Created successfully",
		Data:    response,
	}, nil

}

func (ps *ProductService) GetProductByID(c context.Context, id string) (dto.GetProductHTTPResponse, *errorutils.Error) {

	ctx, cancelFunc := context.WithCancel(c)
	defer cancelFunc()

	product_id, _ := primitive.ObjectIDFromHex(id)

	product, err := ProductRepository.FindProductByID(ctx, product_id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dto.GetProductHTTPResponse{}, &Err.INVALID_ERR.PRODUCT_ID
		}
		fmt.Printf("Error fetching Product: %+v", err.Error())
		return dto.GetProductHTTPResponse{}, &Err.INTERNAL_ERR
	}

	response := dto.GetProductResponse{
		ID:           product.ID.Hex(),
		Name:         product.Name,
		Price:        product.Price,
		Category:     product.Category,
		Availability: product.Availability,
	}

	return dto.GetProductHTTPResponse{
		Status:  200,
		Message: "The details of the product for requested Id",
		Data:    response,
	}, nil
}

func (ps *ProductService) GetProductList(c context.Context) (dto.ProductListHTTPResponse, *errorutils.Error) {

	ctx, cancelFunc := context.WithCancel(c)
	defer cancelFunc()

	var product_resp []dto.GetProductResponse

	products, err := ProductRepository.FindProductList(ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dto.ProductListHTTPResponse{}, nil
			fmt.Printf("Error fetching products: %+v", err.Error())
			return dto.ProductListHTTPResponse{}, &Err.INTERNAL_ERR
		}

	}

	for _, productData := range products {
		response := dto.GetProductResponse{
			ID:           productData.ID.Hex(),
			Name:         productData.Name,
			Price:        productData.Price,
			Category:     productData.Category,
			Availability: productData.Availability,
		}
		product_resp = append(product_resp, response)
	}

	return dto.ProductListHTTPResponse{
		Status:  200,
		Message: "product List success",
		Data:    product_resp,
	}, nil
}

func (ps *ProductService) UpdateProduct(c context.Context, id string, req dto.UpdateProduct) (dto.UpdateProductHTTPResponse, *errorutils.Error) {

	productUpdate := dto.UpdateProduct{
		ID:           id,
		Name:         req.Name,
		Price:        req.Price,
		Category:     req.Category,
		Availability: req.Availability,
	}

	product, err := ProductRepository.UpdateProduct(c, productUpdate)
	if err != nil {
		fmt.Printf("Error updating Product: %+v", err.Error())
		return dto.UpdateProductHTTPResponse{}, &Err.INTERNAL_ERR
	}

	response := dto.UpdateProduct{
		ID:           product.ID.Hex(),
		Name:         product.Name,
		Category:     product.Category,
		Price:        product.Price,
		Availability: product.Availability,
	}

	return dto.UpdateProductHTTPResponse{
		Status:  201,
		Message: "The product data is updated successfully for given id.",
		Data:    response,
	}, nil
}

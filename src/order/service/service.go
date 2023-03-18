package service

import (
	"context"
	"e-com/src/order/dto"
	orders "e-com/src/order/repository"
	products "e-com/src/product/repository"
	errorsutils "e-com/utils/errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct{}

var OrderRepository = new(orders.Order)
var ProductRepository = new(products.Product)

var Err = errorsutils.NewErr()

// create order service function
func (os *OrderService) CreateOrder(c context.Context, req dto.CreateOrder) (dto.CreateOrderHTTPResponse, *errorsutils.Error) {

	var order orders.Order
	var prod dto.Product

	order.Name = req.OrderName

	var prod_ids []primitive.ObjectID
	var premium []string

	for _, product := range req.Product {

		product_id, _ := primitive.ObjectIDFromHex(product.ProductID)
		ProductID, err := ProductRepository.FindProductByID(c, product_id)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return dto.CreateOrderHTTPResponse{}, &Err.INVALID_ERR.PRODUCT_ID
			}
			fmt.Printf("Error fetching Product: %+v", err.Error())
			return dto.CreateOrderHTTPResponse{}, &Err.INTERNAL_ERR
		}

		if ProductID.Availability {

			prod_ids = append(prod_ids, ProductID.ID)

			if ProductID.Category == "premium" {
				premium = append(premium, ProductID.Category)
			}

			prod.ProductID = ProductID.ID.Hex()
			if product.Quantity <= 10 {
				prod.Quantity = product.Quantity
			}

			if product.Quantity == 10 {
				availablity := false
				_, err := ProductRepository.UpdateProductAvailability(c, product.ProductID, availablity)
				if err != nil {

					fmt.Printf("Error updating product availablity: %+v", err.Error())
					return dto.CreateOrderHTTPResponse{}, &Err.INTERNAL_ERR

				}

			}

			ordervalues := ProductID.Price * float64(product.Quantity)
			order.OrderValue += ordervalues
			order.Product = append(order.Product, prod)
		}

	}

	order.CreatedAt = time.Now().UTC()
	order.UpdatedAt = time.Now().UTC()

	if len(prod_ids) >= 3 && len(premium) >= 3 {
		order.OrderValue = order.OrderValue * 0.9
	}

	order.OrderStatus = "placed"

	if order.OrderStatus == "dispatched" {
		order.DispatchDate = time.Time{}
	}

	orderID, err := OrderRepository.CreateOrder(c, order)
	if err != nil {
		fmt.Printf("Error creating order: %+v", err.Error())
		return dto.CreateOrderHTTPResponse{}, &Err.INTERNAL_ERR
	}

	response := dto.CreateOrderResponse{
		ID: orderID.(primitive.ObjectID).Hex(),
	}

	return dto.CreateOrderHTTPResponse{
		Status:  201,
		Message: "Order Created successfully",
		Data:    response,
	}, nil
}

// get order by its id function
func (os *OrderService) GetOrderByID(c context.Context, id string) (dto.GetOrderHTTPResponse, *errorsutils.Error) {

	ctx, cancelFunc := context.WithCancel(c)
	defer cancelFunc()

	order_id, _ := primitive.ObjectIDFromHex(id)

	order, err := OrderRepository.FindOrderByID(ctx, order_id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dto.GetOrderHTTPResponse{}, &Err.INVALID_ERR.ORDER_ID
		}
		fmt.Printf("Error fetching order: %+v", err.Error())
		return dto.GetOrderHTTPResponse{}, &Err.INTERNAL_ERR
	}

	var prod dto.OrderProductResponse
	var orderProduct []dto.OrderProductResponse

	for _, product := range order.Product {

		product_id, _ := primitive.ObjectIDFromHex(product.ProductID)
		ProductID, err := ProductRepository.FindProductByID(c, product_id)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return dto.GetOrderHTTPResponse{}, nil
			}
			fmt.Printf("Error fetching Product: %+v", err.Error())
			return dto.GetOrderHTTPResponse{}, &Err.INTERNAL_ERR
		}

		prod.ID = ProductID.ID.Hex()
		prod.Name = ProductID.Name
		prod.Price = ProductID.Price
		prod.Quantity = product.Quantity
		orderProduct = append(orderProduct, prod)
	}

	if order.OrderStatus == "placed" {

		response := dto.GetOrderResponse{
			ID:          order.ID.Hex(),
			OrderName:   order.Name,
			OrderValue:  order.OrderValue,
			OrderStatus: order.OrderStatus,
			Product:     orderProduct,
		}

		return dto.GetOrderHTTPResponse{
			Status:  200,
			Message: "The details of the order for requested Id,",
			Data:    response,
		}, nil

	} else {

		response := dto.GetOrderWithDispatchResponse{
			ID:           order.ID.Hex(),
			OrderName:    order.Name,
			OrderValue:   order.OrderValue,
			OrderStatus:  order.OrderStatus,
			Product:      orderProduct,
			DispatchDate: order.DispatchDate,
		}

		return dto.GetOrderHTTPResponse{
			Status:  200,
			Message: "The details of the order for requested Id,",
			Data:    response,
		}, nil
	}

}

// update order function
func (os *OrderService) UpdateOrder(c context.Context, id string, req dto.UpdateOrder) (dto.UpdateOrderHTTPResponse, *errorsutils.Error) {

	OrderUpdate := dto.UpdateOrder{
		ID:          id,
		OrderStatus: req.OrderStatus,
	}

	var dispatch_date time.Time

	if OrderUpdate.OrderStatus == "dispatched" {
		dispatch_date = time.Now().AddDate(0, 0, 3)
	}

	order, err := OrderRepository.UpdateOrder(c, OrderUpdate, dispatch_date)
	if err != nil {
		fmt.Printf("Error updating order: %+v", err.Error())
		return dto.UpdateOrderHTTPResponse{}, &Err.INTERNAL_ERR
	}

	response := dto.UpdateOrders{
		ID:           order.ID.Hex(),
		OrderName:    order.Name,
		OrderValue:   order.OrderValue,
		OrderStatus:  order.OrderStatus,
		DispatchDate: order.DispatchDate,
		Product:      order.Product,
	}

	return dto.UpdateOrderHTTPResponse{
		Status:  201,
		Message: "The Order data is updated successfully for given id.",
		Data:    response,
	}, nil

}

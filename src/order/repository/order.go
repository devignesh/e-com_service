package repository

import (
	"context"
	"e-com/src/order/dto"
	"e-com/utils/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongo db colletion
var OrderRepo = database.Db().Database("josh-com").Collection("order")

// order schema
type Order struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Product      []dto.Product      `json:"product" bson:"product"`
	Name         string             `json:"name" bson:"name,omitempty"`
	OrderValue   float64            `json:"value" bson:"value,omitempty"`
	OrderStatus  string             `json:"status" bson:"status,omitempty"`
	DispatchDate time.Time          `json:"dispatch" bson:"dispatch,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	IsDeleted    bool               `json:"is_deleted" bson:"is_deleted"`
}

// create new order
func (o *Order) CreateOrder(ctx context.Context, order Order) (interface{}, error) {

	r, err := OrderRepo.InsertOne(ctx, order)
	if err != nil {
		return "", err
	}

	return r.InsertedID, nil
}

// find a oder by its id
func (o *Order) FindOrderByID(ctx context.Context, id primitive.ObjectID) (Order, error) {

	var order Order

	err := OrderRepo.FindOne(ctx, bson.M{"_id": id}).Decode((&order))
	if err != nil {
		return order, err
	}

	return order, nil
}

// find order by filter
func (o *Order) FindOrderByBson(ctx context.Context, filter bson.D) (Order, error) {

	var order Order

	err := OrderRepo.FindOne(ctx, filter).Decode((&order))
	if err != nil {
		return order, err
	}

	return order, nil
}

// update order
func (o *Order) UpdateOrder(ctx context.Context, req dto.UpdateOrder, date time.Time) (Order, error) {

	var order Order

	after := options.After

	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var update bson.D

	if len(req.OrderStatus) != 0 {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"status": req.OrderStatus,
		}})
	}

	update = append(update, primitive.E{Key: "$set", Value: bson.M{
		"dispatch": date,
	}})

	if update != nil {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"updated_at": time.Now().UTC(),
		}})
	}

	id, _ := primitive.ObjectIDFromHex(req.ID)

	err := OrderRepo.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, &returnOpt).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}

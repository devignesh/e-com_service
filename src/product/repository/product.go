package repository

import (
	"context"
	"e-com/src/product/dto"
	"e-com/utils/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productRepo = database.Db().Database("josh-com").Collection("product")

type Product struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Price        float64            `json:"price" bson:"price,omitempty"`
	Category     string             `json:"category" bson:"category,omitempty"`
	Availability bool               `json:"availability" bson:"availability,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	IsDeleted    bool               `json:"is_deleted" bson:"is_deleted"`
}

func (p *Product) CreateProduct(ctx context.Context, product Product) (interface{}, error) {

	r, err := productRepo.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	return r.InsertedID, nil
}

func (p *Product) FindProductByID(ctx context.Context, id primitive.ObjectID) (Product, error) {

	var product Product

	err := productRepo.FindOne(ctx, bson.M{"_id": id}).Decode((&product))
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *Product) FindProductByBson(ctx context.Context, filter bson.D) (Product, error) {

	var product Product

	err := productRepo.FindOne(ctx, filter).Decode((&product))
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *Product) FindProductList(ctx context.Context) ([]Product, error) {

	var product []Product

	options := options.Find()

	cur, err := productRepo.Find(ctx, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var products Product

		err := cur.Decode(&products)
		if err != nil {
			return nil, err
		}

		product = append(product, products)

	}

	return product, nil
}

func (p *Product) UpdateProduct(ctx context.Context, req dto.UpdateProduct) (Product, error) {

	var product Product

	after := options.After

	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var update bson.D

	if len(req.Name) != 0 {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"name": req.Name,
		}})
	}

	if req.Price != 0 {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"price": req.Price,
		}})
	}

	if len(req.Category) != 0 {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"category": req.Category,
		}})
	}

	if req.Availability || !req.Availability {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"availability": req.Availability,
		}})
	}

	if update != nil {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"updated_at": time.Now().UTC(),
		}})
	}

	id, _ := primitive.ObjectIDFromHex(req.ID)

	err := productRepo.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, &returnOpt).Decode(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *Product) UpdateProductAvailability(ctx context.Context, prod_id string, req bool) (Product, error) {

	var product Product

	after := options.After

	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var update bson.D

	if req || !req {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"availability": req,
		}})
	}

	if update != nil {
		update = append(update, primitive.E{Key: "$set", Value: bson.M{
			"updated_at": time.Now().UTC(),
		}})
	}

	id, _ := primitive.ObjectIDFromHex(prod_id)

	err := productRepo.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, &returnOpt).Decode(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}

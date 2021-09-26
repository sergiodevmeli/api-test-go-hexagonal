package services

import (
	"api_test_hexagonal/config/databases/mongodb"
	"api_test_hexagonal/pkg/domain/entities"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductService struct{}

var collection = mongodb.Connect("products")
var ctx = context.Background()

func (p *ProductService) GetAllProducts() (entities.IProductsEntity, error) {
	var products entities.IProductsEntity
	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}


	for cur.Next(ctx) {
		var product entities.IProductEntity
		err = cur.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	fmt.Println("response get all products: ", products)
	}

	return products, nil
}

func (p *ProductService) CreateProduct(product entities.IProductEntity) error {
	var err error
	_, err = collection.InsertOne(ctx, product)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) UpdateProduct(product entities.IProductEntity, productID string)error {

	var err error

	objectID, _ := primitive.ObjectIDFromHex(productID)
	filter := bson.M{"_id": objectID}

	updateProduct := bson.M{
		"$set": bson.M{
			"name": product.Name,
			"status": product.Status,
			"description": product.Description,
			"price": product.Price,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, updateProduct)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductService) DeleteProduct(productID string) error {
	return nil
}

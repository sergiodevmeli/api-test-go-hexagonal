package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// IProductEntity data
type IProductEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `json:"name"`
	Status bool `json:"status"`
	Description string `json:"description"`
	Price int `json:"price,omitempty"`
	CreatedAt	time.Time	`bson:"created_At" json:"created_at"`
	UpdatedAt	time.Time	`bson:"updated_at" json:"updated_at,omitempty"`
}

// IProductsEntity list
type IProductsEntity []*IProductEntity
package repositories

import (
	"api_test_hexagonal/pkg/domain/entities"
)

type IProductRepository interface {
	CreateProduct(product entities.IProductEntity) error
	GetAllProducts() (entities.IProductsEntity, error)
	UpdateProduct(entities.IProductEntity, string) error
	DeleteProduct(string) error
}

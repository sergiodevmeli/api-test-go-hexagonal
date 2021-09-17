package use_case

import (
	"api_test_hexagonal/pkg/products/domain/entities"
	"api_test_hexagonal/pkg/products/domain/repositories"
)

type ProductUseCase struct {
	Repository repositories.IProductRepository
}

func (repo *ProductUseCase) GetAllProducts() (entities.IProductsEntity, error) {
	return repo.Repository.GetAllProducts()
}

func (repo *ProductUseCase) CreateProduct(product entities.IProductEntity) error {
	return repo.Repository.CreateProduct(product)
}

func (repo *ProductUseCase) UpdateProduct(product entities.IProductEntity, productID string) error {
	return repo.Repository.UpdateProduct(product, productID)
}

func (repo *ProductUseCase) DeleteProducts(productID string) error {
	return repo.Repository.DeleteProduct(productID)
}

package usecase

import (
	"product-manager/model"
	"product-manager/repository"
)

type ProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		productRepository: repository,
	}
}

func (pu *ProductUsecase) GetProductUseCase() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}

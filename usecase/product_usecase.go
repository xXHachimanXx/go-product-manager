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

func (pu *ProductUsecase) GetProductById(id int) (model.Product, error) {
	return pu.productRepository.GetProductById(id)
}

func (pu *ProductUsecase) GetProductUseCase() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.productRepository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

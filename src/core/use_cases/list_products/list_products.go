package listproducts

import (
	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
)

type ListProductsUseCase struct {
	productRepository repositories.ProductRepository
}

func (useCase *ListProductsUseCase) Execute(params repositories.ListProductsParams) []entities.Product {
	return useCase.productRepository.List(params)
}

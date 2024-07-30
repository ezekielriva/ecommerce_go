package showproduct

import (
	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
)

type ShowProductUseCase struct {
	productRepository repositories.ProductRepository
}

func NewShowProductUseCase(productRepository repositories.ProductRepository) *ShowProductUseCase {
	return &ShowProductUseCase{
		productRepository: productRepository,
	}
}

func (u *ShowProductUseCase) Execute(id entities.ProductID) (*entities.Product, error) {
	product, err := u.productRepository.Get(id)

	if err != nil {
		return nil, &ProductNotFoundError{id: id}
	}

	return product, nil
}

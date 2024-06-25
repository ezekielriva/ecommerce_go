package showproduct

import (
	"errors"
	"testing"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
	"github.com/stretchr/testify/mock"
)

type MProductRepository struct {
	mock.Mock
}

func (m *MProductRepository) List(params repositories.ListProductsParams) []entities.Product {
	args := m.Called()
	return args.Get(0).([]entities.Product)
}

func (m *MProductRepository) Get(id entities.ProductID) (*entities.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func TestShowProductUseCase(t *testing.T) {
	testCases := []struct {
		desc      string
		id        entities.ProductID
		productUT *entities.Product
		err       error
	}{
		{
			desc:      "Show Existing Product",
			id:        entities.ProductID(1),
			productUT: &entities.Product{Id: 1, Name: "Test"},
		},
		{
			desc: "Throw an error when Product doesnt exist",
			id:   entities.ProductID(1),
			err:  &ProductNotFoundError{id: 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var repository *MProductRepository = &MProductRepository{}
			var useCase *ShowProductUseCase = NewShowProductUseCase(repository)

			repository.On("Get", tC.id).Return(tC.productUT, tC.err).Times(1)

			product, err := useCase.Execute(tC.id)

			if product != nil && product.Id != tC.productUT.Id {
				t.Errorf("Product doesnt match. Actual (%d) Expected (%d)", product.Id, tC.productUT.Id)
			}

			if err != nil && errors.Is(err, tC.err) {
				t.Errorf("Error doesnt match. Actual (%d) Expected (%d)", err, tC.err)
			}

			repository.AssertExpectations(t)
		})
	}
}

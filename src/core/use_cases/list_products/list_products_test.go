package listproducts

import (
	"testing"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
	"github.com/stretchr/testify/mock"
)

type MProductRepository struct {
	mock.Mock
}

func (m *MProductRepository) List(params repositories.ListProductsParams) []entities.Product {
	args := m.Called(params)
	return args.Get(0).([]entities.Product)
}

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		len    int
		params repositories.ListProductsParams
	}{
		{
			desc:   "Returns all products",
			len:    2,
			params: repositories.ListProductsParams{},
		},
		{
			desc: "Returns only products with id=1",
			len:  1,
			params: repositories.ListProductsParams{Filters: []repositories.TListFilter{
				{Field: "id", Value: "1"},
			}},
		},
		{
			desc: "Returns only products with name=pepsi",
			len:  3,
			params: repositories.ListProductsParams{Filters: []repositories.TListFilter{
				{Field: "name", Value: "pepsi", Operation: "eq"},
			}},
		},
		{
			desc: "Returns only products with name=pepsi and price greather than 0",
			len:  3,
			params: repositories.ListProductsParams{Filters: []repositories.TListFilter{
				{Field: "name", Value: "pepsi"},
				{Field: "price", Value: "0", Operation: "gt"},
			}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			repository := &MProductRepository{}

			repository.On("List", tC.params).Return(make([]entities.Product, tC.len)).Times(1)

			var useCase ListProductsUseCase = ListProductsUseCase{
				productRepository: repository,
			}

			useCase.Execute(tC.params)

			repository.AssertExpectations(t)
		})
	}
}

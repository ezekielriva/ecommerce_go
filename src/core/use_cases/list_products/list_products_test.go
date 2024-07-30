package listproducts

import (
	"errors"
	"testing"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
	mock_repositories "github.com/ezekielriva/ecommerce_go/src/core/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		len      int
		params   repositories.ListProductsParams
		products []entities.Product
	}{
		{
			desc:     "Returns all products",
			len:      2,
			params:   repositories.ListProductsParams{},
			products: make([]entities.Product, 2),
		},
		{
			desc: "Returns only products with id=1",
			len:  1,
			params: repositories.ListProductsParams{Filters: []repositories.TListFilter{
				{Field: "id", Value: "1"},
			}},
			products: make([]entities.Product, 1),
		},
		{
			desc: "Returns only products with name=pepsi",
			len:  3,
			params: repositories.ListProductsParams{Filters: []repositories.TListFilter{
				{Field: "name", Value: "pepsi", Operation: "eq"},
			}},
			products: make([]entities.Product, 3),
		},
		{
			desc: "Returns only products with name=pepsi and price greather than 0",
			len:  3,
			params: repositories.ListProductsParams{Filters: []repositories.TListFilter{
				{Field: "name", Value: "pepsi"},
				{Field: "price", Value: "0", Operation: "gt"},
			}},
			products: make([]entities.Product, 3),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_repositories.NewMockProductRepository(ctrl)

			repo.EXPECT().List(tC.params).Times(1).Return(tC.products)

			var useCase ListProductsUseCase = ListProductsUseCase{
				productRepository: repo,
			}

			products := useCase.Execute(tC.params)

			assert.Len(t, products, tC.len, errors.New("Not enough items"))
		})
	}
}

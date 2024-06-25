package repositories

import (
	"github.com/ezekielriva/ecommerce_go/src/core/entities"
)

type TListFilter struct {
	Field     string
	Value     string
	Operation string
}

type ListProductsParams struct {
	Filters []TListFilter
}

type ProductRepository interface {
	List(params ListProductsParams) []entities.Product
}

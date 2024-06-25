package showproduct

import (
	"fmt"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
)

type ProductNotFoundError struct {
	id entities.ProductID
}

func (e *ProductNotFoundError) Error() string {
	return fmt.Sprintf("Product %d Not Found", e.id)
}

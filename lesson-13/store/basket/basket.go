package basket

import "lesson-13/store/product"

type Basket struct {
	ProductList product.ProductList
	Total       int
}

func NewBasket() Basket {
	return Basket{}
}

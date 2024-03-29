package customer

import (
	"fmt"
	"lesson-13/store/basket"
	"lesson-13/store/product"
)

type Customer struct {
	Name   string
	Cash   int
	Basket basket.Basket
}

func NewCustomer(name string, cash int, basket basket.Basket) Customer {
	return Customer{
		Name:   name,
		Cash:   cash,
		Basket: basket,
	}
}

func (c *Customer) AddProduct(p product.Product) {
	c.Basket.ProductList.AddProduct(p)
	for _, p := range c.Basket.ProductList {
		c.Basket.Total += p.Price * p.Quantity
	}
}

func GetCustomerInfo() (name string, cash int) {
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your cash: ")
	fmt.Scan(&cash)

	return
}

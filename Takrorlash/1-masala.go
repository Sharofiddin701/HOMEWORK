// package main

// import "fmt"

// type Basket interface {
// }

// type Client struct {
// 	Id     int
// 	Name   string
// 	Basket []BasketProduct
// }

// type BasketProduct struct {
// 	ProductId int
// 	Quantity  int
// }

// type Product struct {
// 	Id    int
// 	Name  string
// 	Price int
// }

// var productDatabase = []Product{
// 	{Id: 1, Name: "Product1", Price: 10},
// 	{Id: 2, Name: "Product2", Price: 20},
// 	{Id: 3, Name: "Product3", Price: 15},
// 	{Id: 4, Name: "Product4", Price: 25},
// }

// func (c *Client) AddtoBasket(product Product, quantity int) {

// 	for i, item := range c.Basket {
// 		if item.ProductId == product.Id {
// 			c.Basket[i].Quantity += quantity
// 			return
// 		}

// 	}
// 	c.Basket = append(c.Basket, BasketProduct{ProductId: product.Id, Quantity: quantity})



// }
// func main() {
// 	client := Client{

// 		Id:   1,
// 		Name: "Jahon",
// 		Basket: []BasketProduct{
// 			BasketProduct{
// 				ProductId: 1,
// 				Quantity:  21,
// 			},
// 		},
// 	}
// 	client.AddtoBasket(productDatabase[0], 4)
// 	client.AddtoBasket(productDatabase[1], 2)
// 	fmt.Print(client.Basket)

// }

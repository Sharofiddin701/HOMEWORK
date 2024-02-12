// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Branch struct {
// 	Id      int
// 	Name    string
// 	Address string
// }

// type Transaction struct {
// 	Id        int
// 	BranchId  int
// 	ProductId int
// 	Quantity  int
// }

// type Product struct {
// 	Id    int
// 	Name  string
// 	Price int
// }

// var products = []Product{
// 	{Id: 1, Name: "Olma", Price: 12000},
// 	{Id: 2, Name: "Banan", Price: 22000},
// 	{Id: 3, Name: "Olcha", Price: 25000},
// }

// var transactions = []Transaction{
// 	{Id: 1, BranchId: 1, ProductId: 1, Quantity: 12},
// 	{Id: 2, BranchId: 2, ProductId: 2, Quantity: 10},
// 	{Id: 3, BranchId: 1, ProductId: 3, Quantity: 8},
// 	{Id: 4, BranchId: 2, ProductId: 1, Quantity: 14},
// 	{Id: 5, BranchId: 1, ProductId: 2, Quantity: 13},
// 	{Id: 6, BranchId: 2, ProductId: 3, Quantity: 7},
// }

// var branches = []Branch{
// 	{Id: 1, Name: "MGorkiy", Address: "Mirzo Ulug'bek 17 uy"},
// 	{Id: 2, Name: "Chilonzor", Address: "Chilonzor Metro"},
// }

// func main() {
// 	productBranchQuantities := make(map[int]int)

// 	for _, branch := range branches {
// 		for _, transaction := range transactions {

// 			productBranchQuantities[transaction.ProductId] += transaction.Quantity

// 		}
// 	}
// 	maxQuantity := math.MinInt64
// 	var maxProduct Product

// 	for _, product := range products {
// 		quantity := productBranchQuantities[product.Id]
// 		fmt.Printf("Branch: %s,Product: %s, Quantity: %d\n", branch.Name, product.Name, quantity)

// 		if quantity > maxQuantity {
// 			maxQuantity = quantity
// 			maxProduct = product
// 		}
// 	}

// 	fmt.Printf("Product : %s, Quantity: %d\n", maxProduct.Name, maxQuantity)
// }

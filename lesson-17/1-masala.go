// package main

// import "fmt"

// type Branch struct {
// 	Id      int
// 	Name    string
// 	Address string
// }

// type Transacation struct {
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

// var product = []Product{
// 	{Id: 1, Name: "Olma", Price: 12000},
// 	{Id: 2, Name: "Banan", Price: 22000},
// 	{Id: 3, Name: "Olcha", Price: 25000},
// }
// var transaction = []Transacation{
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

// 	for _, item := range branches {
// 		summ1:= 0
// 		fmt.Println("Branch")
// 		for _, item1 := range transaction {
// 			for _, item2 := range product {

// 				if item1.ProductId == item2.Id && item.Id == item1.BranchId {

// 					summ1 += item1.Quantity * item2.Price

// 				}
// 			}

// 		}
// 		fmt.Println("Summ1=", summ1)
// 	}

// }

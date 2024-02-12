// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type BranchTransaction struct {
// 	Id        int    `json:"id"`
// 	BranchId  int    `json:"branch_id"`
// 	ProductId int    `json:"product_id"`
// 	Type      string `json:"type"`
// 	Quantity  int    `json:"quantity"`
// 	Data      string `json:"created_at"`
// }

// type Product struct {
// 	Id          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Price       int    `json:"price"`
// 	Category_id int    `json:"category_id"`
// }

// func main() {

// 	data, err := os.ReadFile("branch_pr_transaction.json")
// 	if err != nil {
// 		log.Fatalf("Bunday Json fayli yo'q: %v", err)
// 	}

// 	data1, err := os.ReadFile("products.json")
// 	if err != nil {
// 		log.Fatalf("Bunday Json fayli yo'q: %v", err)
// 	}

// 	var branchTransactions []BranchTransaction
// 	var products []Product

// 	if err := json.Unmarshal(data, &branchTransactions); err != nil {
// 		log.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	if err := json.Unmarshal(data1, &products); err != nil {
// 		log.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	for _, branchTransaction := range branchTransactions {
// 		var k1, ch1 int

// 		for _, product := range products {
// 			if branchTransaction.ProductId == product.Id {
// 				if branchTransaction.Type == "plus" {
// 					k1 += branchTransaction.Quantity
// 				} else if branchTransaction.Type == "minus" {
// 					ch1 += branchTransaction.Quantity
// 				}
// 				fmt.Println("Name:", product.Name, "Kiritilgan(Plus):", k1, "Chiqarilgan(Minus):", ch1)
// 			}
// 		}
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type Branchtransacation struct {
// 	Id        int    `json:"id"`
// 	BranchId  int    `json:"branch_id"`
// 	ProductId int    `json:"product_id"`
// 	Type      string `json:"type"`
// 	Quantity  int    `json:"quantity"`
// 	Data      string `json:"created_at"`
// }
// type Products struct {
// 	Id          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Price       int    `json:"price"`
// 	Category_id int    `json:"category_id"`
// }

// func main() {

// 	data, err := os.ReadFile("./data/branch_pr_transaction.json")
// 	if err != nil {
// 		log.Fatalf("Bunday Json fayli yo'q: %v", err)
// 	}

// 	data1, err := os.ReadFile("./data/products.json")
// 	if err != nil {
// 		log.Fatalf("Bunday Json fayli yo'q: %v", err)
// 	}

// 	var products []Products

// 	var branchtransacations []Branchtransacation

// 	var total1, total2, plus, minus int

// 	if err := json.Unmarshal(data, &branchtransacations); err != nil {
// 		log.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	if err := json.Unmarshal(data1, &products); err != nil {
// 		log.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	for _, branchtransacation := range branchtransacations {

// 		for _, product := range products {
// 			if branchtransacation.Id == product.Id && branchtransacation.Type == "plus" {
// 				total1 += branchtransacation.Quantity * product.Price
// 				plus++
// 			} else if branchtransacation.Id == product.Id && branchtransacation.Type == "minus" {
// 				total2 += branchtransacation.Quantity * product.Price
// 				minus++
// 			}

// 		}
// 	fmt.Println("Plus Summa:", total1, "Minus Summa:", total2, "Pluslar soni :", plus, "Minuslar soni:", minus)

// }
// }
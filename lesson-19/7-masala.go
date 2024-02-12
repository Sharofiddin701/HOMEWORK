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

// func main() {

// 	data, err := os.ReadFile("./data/branch_pr_transaction.json")
// 	if err != nil {
// 		log.Fatalf("Bunday Json fayli yo'q: %v", err)
// 	}

// 	var branchtransacations []Branchtransacation

// 	if err := json.Unmarshal(data, &branchtransacations); err != nil {
// 		log.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	for _, branchtransacation := range branchtransacations {

// 		if branchtransacation.Data

// 		}

// }

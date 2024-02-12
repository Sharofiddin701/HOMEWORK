// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	//"os"
// )

// type Product struct {
// 	ID        int    `json:"id"`
// 	BranchID  int    `json:"branch_id"`
// 	ProductID int    `json:"product_id"`
// 	Type1     string `json:"type1"`
// 	Quantity  int    `json:"quantity"`
// 	CreatedAt string `json:"created_at"`
// }

// func main() {
// 	fileName := "./data/branch_pr_transaction.json"
// 	content, err := os.ReadFile(fileName)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}

// 	var product []Product
// 	err = json.Unmarshal(content, &product)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}

// 	productIDCount := make(map[int]int)
// 	for _, m := range product {
// 		productIDCount[m.ProductID]++
// 	}
// 	for productID, count := range productIDCount {
// 		fmt.Printf("product_id: %d, soni: %d\n", productID, count)
// 	}
// }

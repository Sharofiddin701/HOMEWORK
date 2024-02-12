package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	type Branchtransacation struct {
		Id        int    `json:"id"`
		BranchId  int    `json:"branch_id"`
		ProductId int    `json:"product_id"`
		Type      string `json:"type"`
		Quantity  int    `json:"quantity"`
		Data      string `json:"created_at"`
	}
	type Branches struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	type Products struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Category_id int    `json:"category_id"`
	}

	branchTransactionsData, err := os.ReadFile("./data/branch_pr_transaction.json")
	if err != nil {
		log.Fatalf("Error reading branch transaction JSON file: %v", err)
	}

	branchesData, err := os.ReadFile("./data/branches.json")
	if err != nil {
		log.Fatalf("Error reading branches JSON file: %v", err)
	}

	productsData, err := os.ReadFile("./data/products.json")
	if err != nil {
		log.Fatalf("Error reading products JSON file: %v", err)
	}

	var branchTransactions []Branchtransacation
	if err := json.Unmarshal(branchTransactionsData, &branchTransactions); err != nil {
		log.Fatalf("Error unmarshalling branch transaction JSON: %v", err)
	}

	var branches []Branches
	if err := json.Unmarshal(branchesData, &branches); err != nil {
		log.Fatalf("Error unmarshalling branches JSON: %v", err)
	}

	var products []Products
	if err := json.Unmarshal(productsData, &products); err != nil {
		log.Fatalf("Error unmarshalling products JSON: %v", err)
	}

	branchSales := make(map[string]int)

	for _, transaction := range branchTransactions {

		for _, branch := range branches {

			for _, product := range products {

				if transaction.BranchId == branch.Id && transaction.ProductId == product.Id {
					branchSales[branch.Name] += transaction.Quantity * product.Price
				}
			}
		}
	}

	for branchName, totalSales := range branchSales {
		fmt.Printf("%s: %d\n", branchName, totalSales)
	}
}

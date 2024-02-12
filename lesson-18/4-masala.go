package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Car struct {
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Engine string `json:"engine"`
}

func main() {

	data, err := os.ReadFile("cars.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var cars []Car
	if err := json.Unmarshal(data, &cars); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, car := range cars {
		fmt.Printf("Brand: %s, Model: %s, Year: %d, Engine: %s\n", car.Brand, car.Model, car.Year, car.Engine)
	}
}

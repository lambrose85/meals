package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lambrose85/meals/meals"
)

type Meal struct {
	Name        string       `json:"name"`
	Day         string       `json:"day"`
	Ingredients []Ingredient `json:"ingredients"`
}
type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

func main() {

	var mealz Meal

	filePath := "data/test.json"

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	err = json.Unmarshal(file, &mealz) //parse JSON into Meal Struct
	if err != nil {
		fmt.Println("JSON ERROR: ", err)
		return
	}

	meals.NewRecipe("")

}

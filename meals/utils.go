package meals

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var data = "data/"

// will take input from the command line, parse the arguments and
// output the result to a JSON file
func NewRecipe(path string) {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("FILE EXISTS EXITING...\n")

	} else if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("ERROR ON CREATION EXITING with code : %d", err)
		}
		fmt.Printf("FILE CREATION SUCCESSFUL: %s\n", path)
		defer file.Close()
	}

}

// query directory that contains user recipes and return file names &
// parse string value returned to remove file extension
func ViewRecipes(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileInfo := file.Name()
		fmt.Println(strings.Trim(fileInfo, ".json"))

	}
}

// diplay the json contents of a specified JSON file
func ViewJSON(file string) {
	var builder strings.Builder
	builder.WriteString(data)
	builder.WriteString(file)
	str := builder.String()
	fmt.Println(str)

}

package meals

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var data = "data/"

// will take input from the command line, parse the arguments and
// output the result to a file locally for further use
func NewRecipe(recipe []Meal, path string) error {
	return nil
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
		//fmt.Println(strings.Trim(fileInfo, ".json"))
		ViewJSON(fileInfo)
	}
}

// diplay the json contents of each file
func ViewJSON(file string) {
	var builder strings.Builder
	builder.WriteString(data)
	builder.WriteString(file)
	str := builder.String()
	fmt.Println(str)
	//file, err := os.Open(string(str))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//b, err := io.ReadAll(file)
	//.Println(b)
}

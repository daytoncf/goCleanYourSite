package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/net/html"
)

func main() {

	dir := flag.String("directory", "./content/", "Path to html files")

	classList := GetClassesHTMLFiles(*dir)
	cleanedList := separateAllClassNames(classList)
	classSet := mapset.NewSet(cleanedList...)

	iter := classSet.Iterator()

	for classname := range iter.C {
		fmt.Println(classname)
	}

}

// Some HTML elements had multiple classes, and their class values would be appended to the list as one class
// This function separates those elements and returns a new slice containing the separated elements
func separateAllClassNames(classList []string) []string {
	var cleanList []string
	// var newList []string
	for _, classes := range classList {
		if strings.Contains(classes, " ") {
			choppedString := strings.Split(classes, " ")
			// Append the chopped up string slice's values to the slice
			cleanList = append(cleanList, choppedString...)
		} else {
			cleanList = append(cleanList, classes)
		}
	}

	return cleanList
}

// Function that will read a directory for html files
func GetClassesHTMLFiles(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("ERR1")

		return []string{}
	}

	files, err := f.Readdirnames(0)
	f.Close()
	if err != nil {
		fmt.Println("ERR2")
		return []string{}
	}

	classes := make([]string, 0)

	for _, file := range files {
		fullFilename := path + file
		classes = append(classes, GetClassesFromHTMLFile(fullFilename)...)
	}

	return classes
}

// Function that will iterate over an html file and generate a list of all class names in the file
func GetClassesFromHTMLFile(path string) []string {
	//read the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//create a new tokenizer
	tokenizer := html.NewTokenizer(file)

	//create a map to store the class names
	classes := make([]string, 0)

	//iterate over the tokens
	for {
		TokenType := tokenizer.Next()
		token := tokenizer.Token()

		//check if the token is a class
		for _, attribute := range token.Attr {
			if attribute.Key == "class" {
				classes = append(classes, attribute.Val)
			}
		}

		if TokenType == html.ErrorToken {
			break
		}
	}

	return classes
}

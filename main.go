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

	fmt.Println(*dir)
	classList := GetClassesHTMLFiles(*dir)
	cleanedList := separateAllClassNames(classList)
	classSet := mapset.NewSet(cleanedList...)

	iter := classSet.Iterator()

	for classname := range iter.C {
		fmt.Println(classname)
	}

	cleanAllCSSFiles(*dir)
}

func cleanCSSFile(path string) {

	// create string array that represents the css file, each value being a different line
	fileString := fileToString(path)
	fileLines := strings.Split(fileString, "\n")

	for _, s := range fileLines {
		fmt.Println(s)
	}
}

func cleanAllCSSFiles(path string) {
	f, err := os.Open(path)
	checkErr(err)
	defer f.Close()

	files, err := f.Readdirnames(0)
	checkErr(err)

	for _, file := range files {
		if strings.HasSuffix(file, ".css") {
			// Do stuff here
			fullFilename := path + file
			cleanCSSFile(fullFilename)
		}
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
	checkErr(err)

	files, err := f.Readdirnames(0)
	checkErr(err)
	defer f.Close()

	classes := make([]string, 0)

	for _, file := range files {
		if strings.HasSuffix(file, ".html") {
			fullFilename := path + file
			classes = append(classes, GetClassesFromHTMLFile(fullFilename)...)
		}
	}

	return classes
}

// Function that will iterate over an html file and generate a list of all class names in the file
func GetClassesFromHTMLFile(path string) []string {
	//read the file
	file, err := os.Open(path)
	checkErr(err)
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

// Function that dumps a file as a string
func fileToString(filename string) string {
	file, err := os.ReadFile(filename)
	checkErr(err)
	return string(file)
}

// for repeated error checking functionality
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

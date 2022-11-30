package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {

	// val := GetClassesFromHTMLFile("/Users/cvad.itservices/dev/cleanss/content/emergency-index.html")

	fmt.Print(GetClassesHTMLFiles("/Users/cvad.itservices/dev/cleanss/content/"))

}

// Function that will read a directory for html files
func GetClassesHTMLFiles(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("ERR1")

		return []string{}
	}

	files, err := f.Readdirnames(0)
	// f.Close()
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
	file, err :=
		os.Open(path)
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

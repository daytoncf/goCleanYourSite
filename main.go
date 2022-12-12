package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	css "github.com/daytoncf/goCleanYourSite/css"
	lib "github.com/daytoncf/goCleanYourSite/pkg/lib"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/net/html"
)

func main() {

	dir := flag.String("directory", "./content/", "Path to html files")

	classList := GetClassesHTMLFiles(*dir)
	cleanedList := separateAllClassNames(classList)
	classSet := mapset.NewSet(cleanedList...)
	cleanAllCSSFiles(*dir, classSet)
}

func cleanCSSFile(filename string, classSet mapset.Set[string]) {

	// Initialize variable that will contain the new file contents
	var newFile string
	// Parse stylesheet
	stylesheet := css.Tokenizer(filename)

	// Iterate over each top level token
	for _, token := range stylesheet.Tokens {
		// Check to see if the token's selector uses a class
		if classes := extractClassesFromSelector(token.Selector); len(classes) > 0 && token.TokenType == css.RULESET {
			for _, class := range classes {
				// If the class is within the class set, add to string
				// Using [1:] to ignore the '.' at the start of the class name
				if classSet.Contains(class[1:]) {
					newFile += token.Serialize() + "\n"
					break // assuming that if it has the first class, it has all of them. Need a better solution
				}
			}
		} else if token.TokenType == css.RULESET {
			newFile += token.Serialize() + "\n"
		}
	}

	// Iterate over at-rules
	for _, atRule := range stylesheet.AtRules {
		newFile += atRule.Serialize() + "\n"
	}

	// Get index of last '/' to find the end of the directory prefix
	dirEnd := strings.LastIndex(filename, "/") + 1
	// Generate new filename with path prefixed
	newFileName := filename[:dirEnd] + "new_" + filename[dirEnd:]

	// Write to the new file and check errors
	err := os.WriteFile(newFileName, []byte(newFile), 0666)
	lib.CheckErr(err)
}

// Iterates over all css files in a directory and runs cleanCSSFile
func cleanAllCSSFiles(path string, classSet mapset.Set[string]) {
	f, err := os.Open(path)
	lib.CheckErr(err)
	defer f.Close()

	files, err := f.Readdirnames(0)
	lib.CheckErr(err)

	for _, file := range files {
		if strings.HasSuffix(file, ".css") {
			// Do stuff here
			fullFilename := path + file
			cleanCSSFile(fullFilename, classSet)
		}
	}
}

// Function that takes in full css selector and returns the classname(s)
// EX: "a.myClass:hover" => ["myClass"], "p.myClass.anotherClass" => ["myClass", "anotherClass"]
func extractClassesFromSelector(fullSelector string) []string {
	if startIndex := strings.Index(fullSelector, "."); startIndex != -1 {
		// Create regular expression to match classnames
		classRegex := regexp.MustCompile("[.][a-zA-Z0-9-_]+")

		// If matches to pattern are found, return them
		if matches := classRegex.FindAll([]byte(fullSelector), -1); matches != nil {
			return lib.ByteSlicesToStringSlice(matches)
		}
		fmt.Println("Pattern found no matches")
	}
	fmt.Println("No class selector found")
	return []string{}
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
	lib.CheckErr(err)

	files, err := f.Readdirnames(0)
	lib.CheckErr(err)
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
	lib.CheckErr(err)
	defer file.Close()

	//create a new tokenizer
	tokenizer := html.NewTokenizer(file)

	//create a map to store the class names
	classes := make([]string, 0)

	//iterate over the tokens
	for {
		TokenType := tokenizer.Next()
		token := tokenizer.Token()

		//check if the token has a class
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

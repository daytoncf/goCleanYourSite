package lib

import (
	"strings"
	"os"
	"log"
)

func PopRuneArrToString(chars *[]rune) string {
	startingLength := len(*chars)
	stringOfRunes := make([]rune, startingLength)

	for i := 0; i < startingLength; i++ {
		// Pop from front of queue
		stringOfRunes[i] = (*chars)[0] // have to wrap with parentheses to access an index in dereferenced pointer
		if i != startingLength-1 {
			*chars = (*chars)[1:]
		} else {
			// if only 1 rune left, make rune empty string
			*chars = []rune{}
		}
	}

	return string(stringOfRunes)
}

func RemoveWhitespace(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.ReplaceAll(input, " ", "")
	return input
}

func FileToString(filename string) string {
	file, err := os.ReadFile(filename)
	CheckErr(err)
	return string(file)
}

// for repeated error checking functionality
func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}


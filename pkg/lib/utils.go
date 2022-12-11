package lib

import (
	"log"
	"os"
	"strings"
)

type Stack []rune

// Check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push value onto stack
func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

// Remove the last value and return
func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return rune(0) // return null rune if stack is empty
	}
	topIndex := len(*s) - 1     // Index of top of stack
	poppedVal := (*s)[topIndex] // assign `poppedVal` value at top of stack
	*s = (*s)[:topIndex]        // remove top most element
	return poppedVal
}

type Queue []rune

// Check if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Push value into queue
func (q *Queue) Push(r rune) {
	*q = append(*q, r)
}

// Remove the first value and return
func (q *Queue) Pop() rune {
	if q.IsEmpty() {
		return rune(0) // return null rune if stack is empty
	}
	poppedVal := (*q)[0] // assign `poppedVal` value at top of stack
	if len(*q) > 1 {
		*q = (*q)[1:] // If length of queue is >= 2, set queue equal to all elements beyond the first
	} else {
		*q = []rune{} // If length of queue is 1, set queue equal to empty slice
	}
	return poppedVal
}

func (q *Queue) PopQueueToString() string {
	startingLength := len(*q)
	stringOfRunes := make([]rune, startingLength)

	for i := 0; i < startingLength; i++ {
		// Pop from front of queue
		stringOfRunes[i] = q.Pop() // have to wrap with parentheses to access an index in dereferenced pointer
	}
	return string(stringOfRunes)
}

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

// Convert a slice of byte slices to a slice of strings
func ByteSlicesToStringSlice(input [][]byte) []string {
	// Initialize new slice
	newSlice := make([]string, len(input))

	// Convert and insert
	for i, v := range input {
		newSlice[i] = string(v)
	}
	return newSlice
}

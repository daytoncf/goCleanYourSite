package css

import (
	"strings"
	lib "github.com/daytoncf/goCleanSS/pkg/lib"
)

type TokenType int

const (
	COMMENT TokenType = iota
	RULESET
	ERR
)

type Declaration struct {
	Property string
	Value    string
}

type Token struct {
	TokenType    TokenType
	Selector     string
	Declarations []Declaration
}

// Factory function for CSSToken
func newToken(t TokenType, selector string, declarations []Declaration) Token {
	return Token{t, selector, declarations}
}

// Factory function for CSSDeclaration
func newDeclaration(property, value string) Declaration {
	return Declaration{property, value}
}

func ParseDeclarationBlock(declarationBlock string) []Declaration {
	// Initialize empty array
	declarations := make([]Declaration, 0)
	// Trim trailing and leading whitespace.
	minDeclarationBlock := strings.TrimSpace(declarationBlock)

	// Initialize temp variables to store parsed values
	var tempProperty string
	var tempValue string

	var charQ []rune
	for _, char := range minDeclarationBlock {
		switch char {
		case ':':
			tempProperty = strings.TrimSpace(popRuneArrToString(&charQ))
		case ';':
			tempValue = strings.TrimSpace(popRuneArrToString(&charQ))
			declarations = append(declarations, newDeclaration(tempProperty, tempValue))
		default:
			charQ = append(charQ, char)
		}
	}
	return declarations
}

func popRuneArrToString(chars *[]rune) string {
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

// Removes all whitespace characters within a given string
func tokenizeCSSFile(path string) []Token {
	var tokens []Token

	// Convert file into string to make it easily iterable
	fileString := lib.FileToString(path)
	var charQueue []rune
	var readingComment bool = false

	for _, v := range fileString {
		if !readingComment {
			if v == '/' {
				readingComment = true
			} else {
				charQueue = append(charQueue, v)
			}
		} else {
			if v == '/' {
				readingComment = false
			}
		}

	}
	return tokens
}
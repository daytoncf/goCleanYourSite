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
func NewToken(t TokenType, selector string, declarations []Declaration) Token {
	return Token{t, selector, declarations}
}

// Factory function for CSSDeclaration
func NewDeclaration(property, value string) Declaration {
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
			tempProperty = strings.TrimSpace(lib.PopRuneArrToString(&charQ))
		case ';':
			tempValue = strings.TrimSpace(lib.PopRuneArrToString(&charQ))
			declarations = append(declarations, NewDeclaration(tempProperty, tempValue))
		default:
			charQ = append(charQ, char)
		}
	}
	return declarations
}

// Removes all whitespace characters within a given string
func Tokenizer(path string) []Token {
	var tokens []Token

	// Convert file into string to make it easily iterable
	fileString := lib.FileToString(path)
	var charQueue []rune
	var readingComment bool = false

	var selector string
	var decBlock string
	for _, v := range fileString {
		if !readingComment {
			switch v {
			case '/':
				readingComment = true
			case '{':
				selector = strings.TrimSpace(lib.PopRuneArrToString(&charQueue))
			case '}':
				decBlock = strings.TrimSpace(lib.PopRuneArrToString(&charQueue))
				// fmt.Printf("Selector: %v\n", selector)
				// fmt.Println(decBlock)
				tokens = append(tokens, NewToken(RULESET, selector, ParseDeclarationBlock(decBlock)))
			default:
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

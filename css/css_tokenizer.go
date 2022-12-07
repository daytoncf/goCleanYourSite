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

type AtRuleType int

const (
	CHARSET AtRuleType = iota
	COUNTERSTYLE
	FONTFACE
	IMPORT
	KEYFRAMES
	MEDIA
	PAGE
	SUPPORTS
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

type AtRule struct {
	AtRuleType AtRuleType
	Tokens     []Token
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
	for i, v := range fileString {
		switch v {
		case '/':
			if readingComment && peekForCommentEnd(fileString, i) {
				comment := strings.TrimSpace(lib.PopRuneArrToString(&charQueue))
				// Create Token for comment, using its contents for the selector exluding asterisk, [1:len(s)-1]
				tokens = append(tokens, NewToken(COMMENT, comment[1:len(comment)-1], []Declaration{}))
			} else {
				// Check to see if there is an asterisk following this /
				readingComment = peekForCommentStart(fileString, i)
			}
		case '{':
			selector = strings.TrimSpace(lib.PopRuneArrToString(&charQueue))
		case '}':
			decBlock = strings.TrimSpace(lib.PopRuneArrToString(&charQueue))

			// Create new token after declaration block finishes :)
			tokens = append(tokens, NewToken(RULESET, selector, ParseDeclarationBlock(decBlock)))
		default:
			charQueue = append(charQueue, v)
		}
	}
	return tokens
}

// Function that is called after a '/' rune is encountered.
// Check to see if the following rune in string is an asterisk. Returns true if so.
func peekForCommentStart(fileString string, currPos int) bool {
	return fileString[currPos+1] == '*'
}

// Function that is called after a '/' rune is encountered.
// Check to see if the preceding rune in string is an asterisk. Returns true if so.
func peekForCommentEnd(fileString string, currPos int) bool {
	return fileString[currPos-1] == '*'
}

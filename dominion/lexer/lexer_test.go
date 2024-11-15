package lexer

import (
	"dominionlang/token"
	// "testing"
)

func GetLexer(input string) *Lexer {
	return NewLexer(input, GetLiteralReader())
}

func GetLiteralReader() *token.TokenLiteralReader {
	return token.NewLiteralReader(*token.GetDefaultReaders())
}

// func TestNextToken(t *testing.T) {
// 	input := `=+*-/%,(){};`

// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.ASSIGN, "="},
// 		{token.ADD, "+"},
// 		{token.MULT, "*"},
// 		{token.SUBTRACT, "-"},
// 		{token.DIVIDE, "/"},
// 		{token.MODULO, "%"},
// 		{token.COMMA, ","},
// 		{token.LPAREN, "("},
// 		{token.RPAREN, ")"},
// 		{token.LBRACE, "{"},
// 		{token.RBRACE, "}"},
// 		{token.SEMICOL, ";"},
// 		{token.EOF, ""},
// 	}

// 	lexer := GetLexer(input)
// 	for i, test := range tests {
// 		token := lexer.NextToken()

// 		if token.Type != test.expectedType {
// 			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
// 				i, test.expectedType, token.Type)
// 		}

// 		if token.Literal != test.expectedLiteral {
// 			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q", i, test.expectedLiteral, token.Literal)
// 		}
// 	}
// }

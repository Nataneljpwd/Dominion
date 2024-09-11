package lexer

import (
	"dominionlang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+*-/%,(){};`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.ADD, "+"},
		{token.SUBTRACT, "-"},
		{token.MULT, "*"},
		{token.DIVIDE, "/"},
		{token.MODULO, "%"},
		{token.COMMA, ","},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOL, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)
	for i, test := range tests {
		token := lexer.NextToken()

		if token.TYPE != test.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, test.expectedType, token.TYPE)
		}

		if token.LITERAL != test.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q", i, test.expectedLiteral, token.LITERAL)
		}
	}
}

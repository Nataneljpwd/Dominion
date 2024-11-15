package lexer

import (
	"testing"
)

func TestReaderSymbol(t *testing.T) {
	input := `=+*-/%,(){};`

	tests := []string{
		"=",
		"+",
		"*",
		"-",
		"/",
		"%",
		",",
		"(",
		")",
		"{",
		"}",
		";",
		"",
	}

	readers := GetLiteralReader()
	for i, test := range tests {
		char := readers.ReadLiteral(input[i : i+1])

		if char != string(test[i]) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test[i], char)
		}
	}
}

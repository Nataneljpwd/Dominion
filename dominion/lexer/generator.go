package lexer

import (
	"dominionlang/token"
)

type tokenGenerator func(nextChar byte) token.Token

var tokenMap = map[string]tokenGenerator{
	",": func(_ byte) token.Token { return newToken(token.COMMA, ",") },
	";": func(_ byte) token.Token { return newToken(token.SEMICOL, ";") },
	"+": func(_ byte) token.Token { return newToken(token.ADD, "+") },
	"-": func(_ byte) token.Token { return newToken(token.SUBTRACT, "-") },
	"/": func(_ byte) token.Token { return newToken(token.DIVIDE, "/") },
	"*": func(_ byte) token.Token { return newToken(token.MULT, "*") },
	"%": func(_ byte) token.Token { return newToken(token.MODULO, "%") },
	"(": func(_ byte) token.Token { return newToken(token.LPAREN, "(") },
	")": func(_ byte) token.Token { return newToken(token.RPAREN, ")") },
	"{": func(_ byte) token.Token { return newToken(token.LBRACE, "{") },
	"}": func(_ byte) token.Token { return newToken(token.RBRACE, "}") },
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

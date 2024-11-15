package lexer

import (
	"dominionlang/token"
	"dominionlang/utils"
	"errors"
	"fmt"
)

type TokenGenerator func(nextChar string) token.Token

var tokenMap = map[string]TokenGenerator{
	token.COMMA:    func(_ string) token.Token { return newToken(token.COMMA, ",") },
	token.ASSIGN:   func(_ string) token.Token { return newToken(token.ASSIGN, "=") },
	token.COMPARE:  func(_ string) token.Token { return newToken(token.COMPARE, "==") },
	token.SEMICOL:  func(_ string) token.Token { return newToken(token.SEMICOL, ";") },
	token.ADD:      func(_ string) token.Token { return newToken(token.ADD, "+") },
	token.SUBTRACT: func(_ string) token.Token { return newToken(token.SUBTRACT, "-") },
	token.DIVIDE:   func(_ string) token.Token { return newToken(token.DIVIDE, "/") },
	token.MULT:     func(_ string) token.Token { return newToken(token.MULT, "*") },
	token.MODULO:   func(_ string) token.Token { return newToken(token.MODULO, "%") },
	token.LPAREN:   func(_ string) token.Token { return newToken(token.LPAREN, "(") },
	token.RPAREN:   func(_ string) token.Token { return newToken(token.RPAREN, ")") },
	token.LBRACE:   func(_ string) token.Token { return newToken(token.LBRACE, "{") },
	token.RBRACE:   func(_ string) token.Token { return newToken(token.RBRACE, "}") },
	token.EOF:      func(_ string) token.Token { return newToken(token.EOF, "") },
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func GetTokenGenerator(literal string) (TokenGenerator, error) {
	if len(literal) == 0 {
		return tokenMap[token.EOF], nil
	}

	if utils.IsLetter(literal) {
		return newIdentToken, nil
	}

	tokenGenerator, ok := tokenMap[literal]
	fmt.Println("Testing ", literal)

	if !ok {
		for key := range tokenMap {
			fmt.Println(key, literal)
			fmt.Println(key == literal)
			fmt.Println(len(key), len(literal))
		}
		return nil, errors.New(fmt.Sprintf("Invalid token, Got '%s'", literal))
	}

	fmt.Println()
	return tokenGenerator, nil
}

func newIdentToken(literal string) token.Token {
	return newToken(token.IDENT, literal)
}

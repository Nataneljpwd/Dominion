package lexer

import (
	"dominionlang/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0 // terminating with an eof / null-byte
	} else {
		l.currentChar = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	literal := string(l.currentChar)
	tokenGenerator, ok := tokenMap[literal]

	if !ok {
		return token.Token{Type: token.ILLEGAL, Literal: literal}
	}

	if isLetter(l.currentChar) {
		literal = l.readUntilWhitespace()
	} else {
		l.readChar()
	}

	return tokenGenerator(l.nextChar())
}

func (l *Lexer) nextChar() byte {
	if l.readPosition+1 <= len(l.input) {
		return l.input[l.readPosition+1]
	}
	return 0
}

func (l *Lexer) readUntilWhitespace() string {
	var result strings.Builder

	for isWhitespace(l.currentChar) {
		result.WriteByte(l.currentChar)
		l.readChar()
	}

	return result.String()
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

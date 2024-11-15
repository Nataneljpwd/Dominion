package lexer

import (
	"dominionlang/token"
	"log"
	"os"
	"strings"
)

type Lexer struct {
	input         string
	position      int
	readPosition  int
	currentChar   byte
	literalReader LiteralReader
}

type LiteralReader interface {
	ReadLiteral(input string) string
}

func NewLexer(input string, literalReader LiteralReader) *Lexer {
	l := &Lexer{input: input, literalReader: literalReader}
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

	literal := l.readTokenLiteral()

	tokenGenerator, err := GetTokenGenerator(literal)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return tokenGenerator(literal)
}

func (l *Lexer) nextChar() byte {
	if l.readPosition+1 <= len(l.input) {
		return l.input[l.readPosition+1]
	}
	return 0
}

func (l *Lexer) readTokenLiteral() string {
	literal := l.literalReader.ReadLiteral(l.input[l.readPosition:len(l.input)])

	l.readPosition += len(literal)

	return literal
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

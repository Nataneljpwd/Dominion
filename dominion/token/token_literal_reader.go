package token

import (
	"dominionlang/utils"
	"fmt"
	"log"
)

type TokenLiteralReader struct {
	readers utils.PredicateMap[string, LiteralReaderFunc]
}

type LiteralReaderFunc func(string) string

func (r *TokenLiteralReader) ReadLiteral(input string) string {
	reader, ok := r.readers.Get(input)

	if !ok {
		log.Fatal(fmt.Sprintf("Invalid String Input, got '%s' expected valid ascii token", input))
	}

	return reader(input)
}

func NewLiteralReader(tokenReaders utils.PredicateMap[string, LiteralReaderFunc]) *TokenLiteralReader {
	t := &TokenLiteralReader{}
	return t
}

func GetDefaultReaders() *utils.PredicateMap[string, LiteralReaderFunc] {
	var defaultReaders *utils.PredicateMap[string, LiteralReaderFunc] = utils.NewPredicateMap[string, LiteralReaderFunc]()

	defaultReaders.Add(utils.IsLetter, utils.GetReadUntilPredicate(utils.IsSymbol))
	defaultReaders.Add(utils.IsSymbol, utils.GetReadUntilPredicate(utils.IsSymbol))
	defaultReaders.Add(utils.IsNumber, utils.GetReadUntilPredicate(utils.IsNumber))

	return defaultReaders
}

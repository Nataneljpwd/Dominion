package token

const (

	// Terminating
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	INT   = "INT"
	BOOL  = "BOOL"
	IDENT = "IDENT"

	// Operators
	ASSIGN   = "="
	ADD      = "+"
	SUBTRACT = "-"
	MULT     = "*"
	DIVIDE   = "/"
	MODULO   = "%"
	COMPARE  = "=="

	// Delimiters
	COMMA   = ","
	SEMICOL = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	EMPIRE   = "EMPIRE" // the keyword for variables that persist between runs (using the metadata of the file probably)
	MMS      = "MAMAS"  // the keyword for constant variables
)

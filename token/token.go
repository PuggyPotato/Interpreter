package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//IDENTIFIER + LITERAL
	IDENT = "IDENT"
	INT = "INT"

	//OPERATOR
	ASSIGN = "="
	PLUS = "+"

	//DELIMETERS
	COMMA = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	//KEYWORD
	FUNCTION = "FUNCTION"
	LET = "LET"

)
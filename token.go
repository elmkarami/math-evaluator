package main

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	// zero-based
	Column int
	Length int
}

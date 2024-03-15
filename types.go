package main

type TokenType int

const (
	// Single-character tokens
	LEFT_PAREN TokenType = iota + 1
	RIGHT_PAREN
	// Operators
	DOT
	MINUS
	PLUS
	SLASH
	STAR

	// types
	NUMBER
	FUNCCALL
	COMMA

	EOF
	// unknown
	UNKNOWN
)

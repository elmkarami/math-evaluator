package main

import (
	"fmt"
	"strings"
)

type ErrorType int

const (
	Syntax ErrorType = iota
	Unknown
)

var (
	SyntaxError  = EvaluatorError{ErrorKind: Syntax, Message: "Expected expression, found '%v'"}
	UnknownError = EvaluatorError{ErrorKind: Unknown, Message: "Undefined identifier '%v'"}
)

type EvaluatorError struct {
	ErrorKind ErrorType
	Token     Token
	Message   string
}

func (u EvaluatorError) withMessage(token Token, message string) EvaluatorError {
	e := u
	e.Token = token
	if message != "" {
		e.Message = message
	} else {
		e.Message = fmt.Sprintf(e.Message, token.Lexeme)
	}
	return e
}

func (u EvaluatorError) Error() string {
	return u.Message
}

func (err EvaluatorError) from(token Token) EvaluatorError {
	return err.withMessage(token, "")
}

func ReportErrorInSource(err error, source string) string {
	// we always expect an EvaluatorError, can be improved later
	evaluatorErr, _ := err.(EvaluatorError)
	columnIndex := evaluatorErr.Token.Column
	if columnIndex > len(source) {
		columnIndex = len(source)
	}
	pointerLine := strings.Repeat(" ", columnIndex) + strings.Repeat("^", evaluatorErr.Token.Length)
	return fmt.Sprintf("Error at column %d: %s:\n%s\n%s\n",
		columnIndex+1, evaluatorErr.Message, source, pointerLine)
}

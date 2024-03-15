package main

import (
	"fmt"
	"math"
)

type Interpreter struct {
}

func (i *Interpreter) Run(input string) (any, error) {
	scanner := NewScanner(input)
	err := scanner.ScanTokens()
	if err != nil {
		return nil, err
	}
	parser := NewParser(scanner.Tokens)
	expression, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return i.evaluate(expression), nil
}

func (i *Interpreter) evaluate(exp Exp) any {
	return exp.Accept(i)
}

func (i *Interpreter) VisitBinaryExp(exp Binary) any {
	left := i.evaluate(exp.Left).(float64)
	right := i.evaluate(exp.Right).(float64)

	switch exp.Operator.Type {
	case PLUS:
		return left + right
	case MINUS:
		return left - right
	case STAR:
		return left * right
	case SLASH:
		return left / right
	}
	return nil
}

func (i *Interpreter) VisitGroupingExp(exp Grouping) any {
	return i.evaluate(exp.Expression)
}

func (i *Interpreter) VisitLiteralExp(exp Literal) any {
	return exp.Value
}

func (i *Interpreter) VisitUnaryExp(exp Unary) any {
	switch exp.Operator.Type {
	case MINUS:
		return -i.evaluate(exp.Right).(float64)
	case PLUS:
		return i.evaluate(exp.Right).(float64)
	}
	// should never happen
	panic(fmt.Sprintf("Unknown unary operator %s", exp.Operator.Lexeme))
}

func (i *Interpreter) VisitPowCallExp(exp PowCall) any {
	base := i.evaluate(exp.Base).(float64)
	power := i.evaluate(exp.Exponent).(float64)
	return math.Pow(base, power)
}

func (i *Interpreter) VisitFuncCallExp(exp FuncCall) any {
	// values are in radians
	value := i.evaluate(exp.Expression).(float64)
	switch exp.Name.Lexeme {
	case "cos":
		return math.Cos(value)
	case "sin":
		return math.Sin(value)
	case "tan":
		return math.Tan(value)
	case "acos":
		return math.Acos(value)
	case "asin":
		return math.Asin(value)
	case "atan":
		return math.Atan(value)
	case "sqrt":
		return math.Sqrt(value)
	}
	// should never happen
	panic(fmt.Sprintf("Unknown function %s", exp.Name.Lexeme))
}

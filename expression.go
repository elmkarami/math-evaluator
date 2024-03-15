package main

import "fmt"

type Exp interface {
	Accept(Visitor) any
	String() string
}

type Visitor interface {
	VisitBinaryExp(Binary) any
	VisitGroupingExp(Grouping) any
	VisitLiteralExp(Literal) any
	VisitFuncCallExp(FuncCall) any
	VisitPowCallExp(PowCall) any
	VisitUnaryExp(Unary) any
}

type Binary struct {
	Left     Exp
	Operator Token
	Right    Exp
}

func (b Binary) Accept(visitor Visitor) any {
	return visitor.VisitBinaryExp(b)
}

func (b Binary) String() string {
	return "Binary(" + b.Left.String() + " " + "Operator(" + b.Operator.Lexeme + ")" + " " + b.Right.String() + ")"
}

type Grouping struct {
	Expression Exp
}

func (g Grouping) Accept(visitor Visitor) any {
	return visitor.VisitGroupingExp(g)
}

func (g Grouping) String() string {
	return "Grouping(" + g.Expression.String() + ")"
}

type Literal struct {
	Value any
}

func (l Literal) String() string {
	return fmt.Sprintf("Literal(%v)", l.Value)
}

func (l Literal) Accept(visitor Visitor) any {
	return visitor.VisitLiteralExp(l)
}

type FuncCall struct {
	Name       Token
	Expression Exp
}

func (f FuncCall) String() string {
	return "FunCall(" + f.Name.Lexeme + "(" + f.Expression.String() + ")" + ")"
}

func (f FuncCall) Accept(visitor Visitor) any {
	return visitor.VisitFuncCallExp(f)
}

type PowCall struct {
	Base, Exponent Exp
}

func (p PowCall) String() string {
	return "Pow(" + "Base=" + p.Base.String() + ", " + "Exponent=" + p.Exponent.String() + ")"
}

func (p PowCall) Accept(visitor Visitor) any {
	return visitor.VisitPowCallExp(p)
}

type Unary struct {
	Operator Token
	Right    Exp
}

func (u Unary) String() string {
	return "Unary(" + u.Operator.Lexeme + " " + u.Right.String() + ")"
}

func (u Unary) Accept(visitor Visitor) any {
	return visitor.VisitUnaryExp(u)
}

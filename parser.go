package main

type Parser struct {
	Tokens  []Token
	current int
	err     error
}

func NewParser(tokens []Token) *Parser {
	return &Parser{Tokens: tokens, current: 0}
}

func (p *Parser) Parse() (Exp, error) {
	exp := p.expression()
	if p.err != nil {
		return nil, p.err
	}
	if !p.isEnd() {
		return nil, SyntaxError.from(p.peek())
	}
	return exp, nil
}

func (p *Parser) isEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) peek() Token {
	return p.Tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.Tokens[p.current-1]
}

func (p *Parser) advance() Token {
	if !p.isEnd() {
		p.current += 1
	}
	return p.previous()
}

func (p *Parser) match(types ...TokenType) bool {
	for _, type_ := range types {
		if p.check(type_) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(type_ TokenType) bool {
	if p.isEnd() {
		return false
	}
	return p.peek().Type == type_
}

func (p *Parser) expect(type_ TokenType, message string) {
	if p.check(type_) {
		p.advance()
		return
	}
	previous := p.previous()
	p.err = SyntaxError.withMessage(previous, message)

}

func (p *Parser) expression() Exp {
	exp := p.factor()

	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.factor()
		exp = Binary{exp, operator, right}
	}
	return exp
}

func (p *Parser) factor() Exp {
	exp := p.primary()
	for p.match(SLASH, STAR) {
		operator := p.previous()
		right := p.factor()

		exp = Binary{exp, operator, right}
	}
	return exp
}

func (p *Parser) primary() Exp {
	if p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.primary()
		return Unary{operator, right}
	}

	if p.match(LEFT_PAREN) {
		exp := p.expression()
		p.expect(RIGHT_PAREN, "Expected ')' after expression")
		return Grouping{exp}
	}

	if p.match(NUMBER) {
		return Literal{p.previous().Literal}
	}

	if p.match(FUNCCALL) {
		funcName := p.previous()
		p.expect(LEFT_PAREN, "Expected '(' after function call")
		if funcName.Lexeme == "pow" {
			return p.powCall()
		}
		exp := p.expression()
		p.expect(RIGHT_PAREN, "Expected ')' after function params")
		return FuncCall{funcName, exp}
	}
	token := p.peek()
	p.err = SyntaxError.from(token)
	return nil
}

func (p *Parser) powCall() Exp {
	base := p.expression()
	p.expect(COMMA, "Expected ',' after first argument of pow")
	exponent := p.expression()
	p.expect(RIGHT_PAREN, "Expected ')' after function params")
	return PowCall{Base: base, Exponent: exponent}
}

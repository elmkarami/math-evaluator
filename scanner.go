package main

import (
	"math"
	"strconv"
	"unicode"
)

var (
	ReservedFuncCalls = [8]string{"cos", "acos", "sin", "asin", "tan", "atan", "sqrt", "pow"}
)

func isReservedFuncCall(name string) bool {
	for _, reserved := range ReservedFuncCalls {
		if name == reserved {
			return true
		}
	}
	return false
}

func getEOFToken(column int) Token {
	return Token{
		Type:    EOF,
		Lexeme:  "EOF",
		Literal: nil,
		Column:  column,
		Length:  1,
	}
}

type Scanner struct {
	Source  string
	Tokens  []Token
	start   int
	current int
	column  int
}

func (s *Scanner) ScanTokens() error {
	for {
		if s.isEnd() {
			s.Tokens = append(s.Tokens, getEOFToken(s.column+1))
			break
		}
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Scanner) scanToken() error {
	char := s.advance()
	var tokenType TokenType

	switch char {
	case "(":
		tokenType = LEFT_PAREN
	case ")":
		tokenType = RIGHT_PAREN
	case ".":
		tokenType = DOT
	case "-":
		tokenType = MINUS
	case "+":
		tokenType = PLUS
	case "*":
		tokenType = STAR
	case "/":
		tokenType = SLASH
	case ",":
		tokenType = COMMA
	case " ", "\t", "\r":
		break
	case "":
		break
	case "\n":
		return SyntaxError.withMessage(s.getToken(UNKNOWN, nil), "Line break not allowed")
	default:
		var err error
		if s.isDigit(char) {
			err = s.scanNumber()
		} else if s.isAlpha(char) {
			err = s.identifier()
		} else {
			err = SyntaxError.from(s.getToken(UNKNOWN, nil))
		}

		if err != nil {
			return err
		}

	}
	if tokenType != 0 {
		s.addToken(tokenType, nil)
	}
	return nil
}

func (s *Scanner) advance() string {
	char := string(s.Source[s.current])
	s.current += 1
	s.column += 1
	return char
}

func (s *Scanner) getToken(tt TokenType, literal any) Token {
	text := s.Source[s.start:s.current]
	return Token{
		Type:    tt,
		Lexeme:  text,
		Literal: literal,
		Length:  len(text),
		Column:  int(math.Abs(float64(s.column - len(text)))),
	}
}

func (s *Scanner) addToken(tt TokenType, literal any) {
	s.Tokens = append(
		s.Tokens,
		s.getToken(tt, literal),
	)
}

func (s *Scanner) peek() string {
	if s.isEnd() {
		return ""
	}
	return string(s.Source[s.current])
}

func (s *Scanner) peekNext() string {
	if (s.current + 1) >= len(s.Source) {
		return ""
	}
	return string([]rune(s.Source)[s.current+1])
}

func (s *Scanner) scanNumber() error {
	for {
		if !s.isDigit(s.peek()) {
			break
		}
		s.advance()
	}
	if s.peek() == "." && s.isDigit(s.peekNext()) {
		s.advance()
		for {
			if !s.isDigit(s.peek()) {
				break
			}
			s.advance()
		}
	}
	value := s.Source[s.start:s.current]
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	s.addToken(NUMBER, number)
	return nil
}

func (s *Scanner) isDigit(value string) bool {
	return value != "" && unicode.IsDigit(rune(value[0]))
}

func (s *Scanner) isAlpha(value string) bool {
	return value != "" && unicode.IsLetter(rune(value[0]))
}

func (s *Scanner) identifier() error {
	for {
		if !s.isAlphaNumeric(s.peek()) {
			break
		}
		s.advance()

	}
	value := s.Source[s.start:s.current]
	if !isReservedFuncCall(value) {
		return UnknownError.from(s.getToken(UNKNOWN, nil))
	}
	s.addToken(FUNCCALL, nil)
	return nil
}

func (s *Scanner) isAlphaNumeric(value string) bool {
	return s.isDigit(value) || s.isAlpha(value)
}

func (s *Scanner) isEnd() bool {
	return s.current >= len(s.Source)
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source: source,
	}
}

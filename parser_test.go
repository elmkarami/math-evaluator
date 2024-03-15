package main

import (
	"strings"
	"testing"
)

func TestParser_ParseToString(t *testing.T) {
	input := "-1+2-((1-cos(1.1))*3)+pow(2,4)"
	expected := `
Binary(
  Binary(
    Binary(
      Unary(- Literal(1))
      Operator(+)
      Literal(2)
    )
    Operator(-)
    Grouping(
      Binary(
        Grouping(
          Binary(
            Literal(1)
            Operator(-)
            FunCall(cos(Literal(1.1)))
          )
        )
        Operator(*)
        Literal(3)
      )
    )
  )
  Operator(+)
  Pow(
    Base=Literal(2),
    Exponent=Literal(4)
  )
)
`
	// Remove all spaces and newlines
	expected = strings.ReplaceAll(expected, "\n", " ")
	expected = strings.ReplaceAll(expected, " ", "")
	got, err := compileInput(input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if strings.ReplaceAll(got, " ", "") != expected {
		t.Errorf("got %v", got)
	}
}

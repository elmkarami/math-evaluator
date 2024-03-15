package main

import (
	"testing"
)

func TestReportErrorInSource(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			input:  "1 +1.",
			expect: "Error at column 5: Expected expression, found '.':\n1 +1.\n    ^\n",
		},
		{
			input:  "1 +",
			expect: "Error at column 4: Expected expression, found 'EOF':\n1 +\n   ^\n",
		},
		{
			input:  "1 + 2 * )",
			expect: "Error at column 9: Expected expression, found ')':\n1 + 2 * )\n        ^\n",
		},
		{
			input:  "2 - (1-4)*3(",
			expect: "Error at column 12: Expected expression, found '(':\n2 - (1-4)*3(\n           ^\n",
		},
		{
			input:  "sin(cos)12",
			expect: "Error at column 8: Expected ')' after function params:\nsin(cos)12\n       ^\n",
		},
		{
			input:  "sin(1 + 2 * tan(12)",
			expect: "Error at column 19: Expected ')' after function params:\nsin(1 + 2 * tan(12)\n                  ^\n",
		},
		{
			input:  "notknown(1 + 2 * tan(12)",
			expect: "Error at column 1: Undefined identifier 'notknown':\nnotknown(1 + 2 * tan(12)\n^^^^^^^^\n",
		},
	}
	Interpreter := Interpreter{}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := Interpreter.Run(tt.input)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			result := ReportErrorInSource(err, tt.input)
			if result != tt.expect {
				println(err)
				t.Errorf("Expected error: %s, got: %s", tt.expect, result)
			}
		})
	}
}

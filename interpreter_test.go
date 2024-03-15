package main

import (
	"testing"
)

func TestInterpreter_(t *testing.T) {
	var tests = []struct {
		input    string
		expected float64
	}{
		// basic arithmetic
		{"1 + 2", 3},
		{"1 + 2.2", 3.2},
		{"+1 - 2", -1},
		{"4 - 2", 2},
		{"2 * 3", 6},
		{"8 / 4", 2},
		{"1 + 2 * 3", 7},
		{"(1 + 2) * 3", 9},
		{"1 + 2 * 3 / 4", 2.5},
		{"-1 + 2 - 3", -2},
		{"2 + -2", 0},

		// Functions (values are in radians for trigonometric functions)
		{"sin(0)", 0},
		{"sin(90)", 0.893996663600558},
		{"sin(3.141592653589793)", 1.2246467991473515e-16},
		{"cos(0)", 1},
		{"cos(90)", -0.4480736161291701},
		{"cos(3.141592653589793)", -1},
		{"tan(0)", 0},
		{"tan(45)", 1.6197751905438615},
		{"tan(3.141592653589793)", -1.2246467991473515e-16},
		{"asin(0)", 0},
		{"acos(1)", 0},
		{"atan(0)", 0},
		{"sqrt(4)", 2},
		{"pow(2, 3)", 8},

		// Complex Expressions
		{"1 + 2 * cos(3 + 4 * 5)", -0.06566604066679504},
		{"1 + 2 * cos(sin(1))", 2.332733490785761},
		{"pow(2, 10)", 1024},
		{"sqrt(pow(2, 4)) + tan(1)", 5.557407724654902},

		// Expressions with Groupings
		{"(1 + 2) * (3 + 4)", 21},
		{"(1 + 2) * (cos(0) + sin(0))", 3},
		{"(1 + 2) * (cos(0) + sin(0)) / 2", 1.5},
		{"(-2 + 2) * cos(0) + sin(0) / 2", 0},
	}
	interpreter := Interpreter{}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := interpreter.Run(test.input)
			if err != nil {
				t.Fatalf("Unexpected error for input %s: %s", test.input, err.Error())
			}
			if result.(float64) != test.expected {
				t.Errorf("Expected %f, got %f for input: '%s'", test.expected, result.(float64), test.input)
			}
		})

	}
}

// test error handling
func TestInterpreter_ErrorHandling(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		// basic arithmetic
		{"+", "Expected expression, found 'EOF'"},
		{"1 +", "Expected expression, found 'EOF'"},
		{"1 + 2 * )", "Expected expression, found ')'"},
		{"2 - (1-4)*3(", "Expected expression, found '('"},
		// functions
		{"sin()", "Expected expression, found ')'"},
		{"sin(1", "Expected ')' after function params"},
		{"sin(1 + 2", "Expected ')' after function params"},
		{"sin(1 + 2 * tan(12)", "Expected ')' after function params"},
		{"notknown(1 + 2 * tan(12)", "Undefined identifier 'notknown'"},
		// grouping
		{"(1 + 2", "Expected ')' after expression"},
		{"(1 + 2) * (3 + 4", "Expected ')' after expression"},
		{"()", "Expected expression, found ')'"},
		// invalid tokens
		{"_", "Expected expression, found '_'"},
		{"1+1\n", "Line break not allowed"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			_, err := executeInput(test.input)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if err.Error() != test.expected {
				t.Errorf("Expected error: %s, got: %s", test.expected, err.Error())
			}
		})
	}
}

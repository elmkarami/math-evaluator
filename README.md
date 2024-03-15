# Simple Math Interpreter

A lightweight, versatile interpreter designed for evaluating mathematical expressions with various supported function. This project encompasses the core components of a traditional interpreter: a scanner (lexer), parser, and evaluator, adapted specifically for mathematical calculations.

Features

- Supported Operations: Includes basic arithmetic operations (addition, subtraction, multiplication, division) and unary operations (positive, negative).
- Function Calls: Supports a variety of mathematical functions such as trigonometric (sin, cos, tan, asin, acos, atan), square root (sqrt), and power (pow) functions.
- Error Reporting: Provides detailed syntax and runtime error messages to aid in debugging expressions.

## Getting Started

**Prerequisites**

Ensure you have Go installed on your machine. This project is built using Go, so it's necessary to run the interpreter. No other dependencies are required.

**Installation**

Clone the repository to your local machine:

```bash
git clone https://github.com/elmkarami/math-evaluator
cd math-evaluator
go build .
```

Usage
To use the interpreter, compile(as shown above) or run the main.go file with your Go environment. You can pass mathematical expressions directly into the command line.

**Examples**:

```bash
>>> ./math-evaluator 'sqrt(9) + pow(2, 3)'
11
>>> ./math-evaluator '-1.2 + 2 * (cos(sin(tan(1))))'
-0.11924455725586158
```

the program can also spot errors and print them nicely for an enhanced experience. Let's say that you miss-spelled the function `cos` in the previous example:

```bash
./math-evaluator '-1.2 + 2 * (coss(sin(tan(1))))'
Error at column 13: Undefined identifier 'coss':
-1.2 + 2 * (coss(sin(tan(1))))
            ^^^^
```

# Architecture Overview

This section aims to offer insight into the thoughtful design of the interpreter:

- Token: Represents the smallest unit of data (like numbers, operators, parentheses, identifiers) in the input.
- Scanner: Analyzes the input string and converts it into a list of tokens.
- Parser: Takes tokens from the Scanner and builds an AST representing the input expression using recursive descent parsing, a simple yet powerful method for many expression grammars. The Parser supports expressions with varying precedence levels and function calls..
- Interpreter: Once the AST is constructed, the Interpreter evaluates the tree to compute the expression's result. It implements the Visitor pattern, allowing it to perform different operations depending on the node type it visits, plus making it easy to support new expressions without modifying the existing code.
- Error Handling: Errors during scanning, parsing, or evaluation are captured and reported with helpful messages.
- Debug mode: provides essential insights for troubleshooting, displaying internal processes like tokenization and parsing. This feature is useful for quickly identifying and resolving errors.

usage for the debug mode:

```bash
>>> ./math-evaluator debug '-1.2 + 2 * tan(1)'
Binary(Unary(- Literal(1.2)) Operator(+) Binary(Literal(2) Operator(*) FunCall(tan(Literal(1)))))
```

**Supported Functions**

The interpreter currently supports the following functions:

- cos, acos, sin, asin, tan, atan: Trigonometric functions expecting arguments in radians.
- sqrt: Computes the square root of its argument.
- pow: Raises the first argument to the power of the second argument.

# Trade-offs and Considerations

- For now it does not support multi line expressions, but it's easier to add it, it should be as simple as a for loop that executes every expression on each line and return a list of results.
- If there are multiple errors, the program will highlight only the first error it faces.
- Can not deal with big numbers, since it's implemented in Go, it will rely on the underlying infrastructure Go offers for numbers (for now all numbers are float64 even ints!!), hence it can show results like: `-1.9999999999998e+33`, but since we're no launching any rocket to the moon, these precision quirks add a bit of charm rather than cause for alarm!

_Nitpicks_:

- The AST implementation does not play well with CPU cache optimization, nodes are allocated across the heap, these jumps over different RAM location are hundreds(even thousands) of times slower than CPU registers. For that, we might consider using a more dense data structure like a an array of instructions which offers an constant time index lookup and constant append to the end, and play nicely with the CPU since the instructions and operands are a few bytes apart.
- We can go even further by precomputing literal terms, at least the numbers, we can directly make our program spit out pre-calculated results into an assembly file, or an IR for LLVM(https://llvm.org/) or QBE(https://c9x.me/compile/)
- Make the debug mode returns a prettified version of the AST
- Accept an expression that spreads across multiple lines for easy reading.
- But wait.. it does not hurt to have a proper LSP :)

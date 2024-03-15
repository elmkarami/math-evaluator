package main

import (
	"fmt"
	"os"
	"strings"
)

func compileInput(input string) (string, error) {
	Scanner := Scanner{Source: input}
	err := Scanner.ScanTokens()
	if err != nil {
		return "", err
	}
	parser := Parser{Tokens: Scanner.Tokens}
	exp, err := parser.Parse()
	if err != nil {
		return "", err
	}
	return exp.String(), nil
}

func executeInput(input string) (any, error) {
	interpreter := Interpreter{}
	return interpreter.Run(input)
}

func PrintResultOrError(input string, result any, err error) {
	if err != nil {
		println(ReportErrorInSource(err, input))
		return
	}
	println(fmt.Sprintf("%v", result))
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		println("usage: ./program '1+ 2 - (1-4)*3'")
		println("usage: ./program debug '1+ 2 - (1-4)*3'")
		return
	}
	command := args[0]
	var result any
	var err error
	var input string
	if command == "debug" {
		input = strings.Join(args[1:], "")
		result, err = compileInput(input)
	} else {
		input = strings.Join(args, "")
		result, err = executeInput(input)
	}
	PrintResultOrError(input, result, err)

}

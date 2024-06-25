package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Polo123456789/go-lox/pkg/lox"
)

func testASTPrinter() {
	expr := lox.ExprBinary{
		Left: &lox.ExprUnary{
			Operator: lox.Token{Type: lox.MINUS, Lexeme: "-", Literal: nil, Line: 1},
			Right:    &lox.ExprLiteral{Value: 123},
		},
		Operator: lox.Token{Type: lox.STAR, Lexeme: "*", Literal: nil, Line: 1},
		Right: &lox.ExprBinary{
			Left:     &lox.ExprLiteral{Value: 45.67},
			Operator: lox.Token{Type: lox.PLUS, Lexeme: "+", Literal: nil, Line: 1},
			Right: &lox.ExprGrouping{
				Expr: &lox.ExprBinary{
					Left:     &lox.ExprLiteral{Value: 1},
					Operator: lox.Token{Type: lox.MINUS, Lexeme: "-", Literal: nil, Line: 1},
					Right:    &lox.ExprLiteral{Value: 2},
				},
			},
		},
	}

	fmt.Println(lox.PrintAST(&expr))
}

func main() {
	testASTPrinter()
	os.Exit(0)

	interpreter := lox.NewLox()
	reader := os.Stdin
	replMode := true

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		reader = file
		replMode = false
		defer reader.Close()
	}

	if replMode {
		fmt.Print(">> ")
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			errors := interpreter.Run(scanner.Text())
			if len(errors) > 0 {
				printErrors(errors)
			}
			fmt.Print(">> ")
		}
	} else {
		buff, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		errors := interpreter.Run(string(buff))
		if len(errors) > 0 {
			printErrors(errors)
		}
	}
}

func printErrors(errors []lox.Error) {
	for _, err := range errors {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

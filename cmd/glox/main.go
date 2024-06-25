package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Polo123456789/go-lox/pkg/lox"
)

func main() {
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

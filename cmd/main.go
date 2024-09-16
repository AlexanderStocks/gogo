package main

import (
	"fmt"
	"os"

	"github.com/AlexanderStocks/GoGo/internal/interpreter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: goi <sourcefile.go>")
		os.Exit(1)
	}

	filename := os.Args[1]

	err := interpreter.RunFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

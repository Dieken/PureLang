package main

import "os"

//go:generate peg grammar.peg

func main() {
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	parser := &PureLang{Buffer: string(buf)}
	err = parser.Init()
	if err != nil {
		panic(err)
	}

	err = parser.Parse()
	if err != nil {
		panic(err)
	}

	parser.PrintSyntaxTree()
}

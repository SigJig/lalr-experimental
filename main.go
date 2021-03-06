package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SigJig/lalr-experimental/internal/analyser"
	"github.com/SigJig/lalr-experimental/internal/lexer"
)

func testLex() {
	pattern := lexer.CompilePattern(analyser.Metapattern)
	f, err := os.Open("internal/analyser/grammar")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	lexer := lexer.NewLexer(reader, pattern, -1)
	ch, ech := lexer.Stream(-1)

	for {
		select {
		case tok, ok := <-ch:
			if !ok {
				return
			}

			if tok.Kind != "ignore" {
				fmt.Println(tok)
			}
		case err := <-ech:
			fmt.Println(err)

			return
		}
	}
}

func testParse() {
	fmt.Println(analyser.TestStates())
}

func main() {
	testParse()
}

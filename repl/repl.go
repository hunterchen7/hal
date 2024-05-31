package repl

import (
	"bufio"
	"fmt"
	"io"
	"waiig/lexer"
	"waiig/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text() // get the just read line of text
		l := lexer.New(line)   // pass the line to a new instance of the lexer

		// loop to get tokens from lexer
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok) // print each token until EOF is encountered
		}
	}
}

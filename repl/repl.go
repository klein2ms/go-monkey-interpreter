package repl

import (
	"bufio"
	"fmt"
	"github.com/klein2ms/go-monkey-interpreter/lexer"
	"github.com/klein2ms/go-monkey-interpreter/token"
	"io"
)

// PROMPT represents the repl prompt
const PROMPT = ">> "

// Start takes a provided io.Reader parsing the input and
// writing the output to the provided io.Writer
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

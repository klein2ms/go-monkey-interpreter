package repl

import (
	"bufio"
	"fmt"
	"github.com/klein2ms/go-monkey-interpreter/lexer"
	"github.com/klein2ms/go-monkey-interpreter/parser"
	"io"
	"github.com/klein2ms/go-monkey-interpreter/evaluator"
)

// PROMPT represents the repl prompt
const PROMPT = ">> "

// MONKEY is output to the console on repl start
const MONKEY = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
		   '-----'
`

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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY)
	io.WriteString(out, "Woopps! We ran into some monkey business here!\n")
	io.WriteString(out, " parser serrors :\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
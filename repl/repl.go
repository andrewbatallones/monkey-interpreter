package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/andrewbatallones/monkey_interpreter/lexer"
	"github.com/andrewbatallones/monkey_interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		ast := p.ParseProgram()

		for i := 0; i < len(ast.Statements); i++ {
			fmt.Fprintf(out, "%+v\n", ast.Statements[i])
		}

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok)
		// }
	}
}

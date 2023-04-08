package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ljesparis/monkey"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>> ")
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := monkey.NewLexer(line)

		for tok := l.NextToken(); tok.Type != monkey.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

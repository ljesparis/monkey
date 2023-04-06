package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ljesparis/monkey/gomonkey"
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
        l := gomonkey.NewLexer(line)

        for tok := l.NextToken(); tok.Type != gomonkey.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}

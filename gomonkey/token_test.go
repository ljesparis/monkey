package gomonkey_test

import (
	"testing"

	"github.com/ljesparis/monkey/gomonkey"
)


func TestLookupIden(t *testing.T) {
    tests := []struct {
        expectedLiteral string
        expectedType gomonkey.TokenType
    } {
        {"fn", gomonkey.FUNCTION},
        {"let", gomonkey.LET},
        {"_foo_bar", gomonkey.IDENTIFIER},
    }

    for i, tt := range tests {
        if tok := gomonkey.LookupIdent(tt.expectedLiteral); tok != tt.expectedType {
            t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok)
        }
    } 
}

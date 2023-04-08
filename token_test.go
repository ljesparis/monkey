package monkey_test

import (
	"testing"

	"github.com/ljesparis/monkey"
)

func TestLookupIden(t *testing.T) {
	tests := []struct {
		expectedLiteral string
		expectedType    monkey.TokenType
	}{
		{"fn", monkey.FUNCTION},
		{"let", monkey.LET},
		{"_foo_bar", monkey.IDENTIFIER},
	}

	for i, tt := range tests {
		if tok := monkey.LookupIdent(tt.expectedLiteral); tok != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok)
		}
	}
}

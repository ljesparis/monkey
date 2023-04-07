package gomonkey_test

import (
	"testing"

	"github.com/ljesparis/monkey/gomonkey"
)

func TestLetStatement(t *testing.T) {
	testCases := []struct {
		input    string
		expected []struct{ expectedIdentifier string }
	}{
		{
			`let x = 5;
            let y = 10;
            let foobar = 898989;`,
			[]struct {
				expectedIdentifier string
			}{
				{"x"},
				{"y"},
				{"foobar"},
			},
		},
	}

	for _, testCase := range testCases {
		l := gomonkey.NewLexer(testCase.input)
		p := gomonkey.NewParser(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if program == nil {
			t.Fatal("ParseProgram() returned nil")
		}

		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements does not contain 3 statements. Got=%d", len(program.Statements))
		}

		for i, ex := range testCase.expected {
			stmt := program.Statements[i]
			if !testLetStatement(t, stmt, ex.expectedIdentifier) {
				return
			}
		}
	}
}

func testLetStatement(t *testing.T, stmt gomonkey.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral() is not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*gomonkey.LetStatement)
	if !ok {
		t.Errorf("s is not LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *gomonkey.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

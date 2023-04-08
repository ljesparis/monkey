package gomonkey_test

import (
	"fmt"
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

func TestReturnStatement(t *testing.T) {
	testCases := []struct {
		input    string
		expected []struct{ expectedIdentifier string }
	}{
		{
			`return 5;
            return 10;
            return 898989;`,
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

		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements does not contain 3 statements. Got=%d", len(program.Statements))
		}

		for _, stmt := range program.Statements {
			returnStmt, ok := stmt.(*gomonkey.ReturnStatement)
			if !ok {
				fmt.Errorf("stmt not *ReturnStatement, got=%T", stmt)
				continue
			}

			if returnStmt.TokenLiteral() != "return" {
				t.Errorf("returnStmt not 'return'. got=%q", returnStmt.TokenLiteral())
			}
		}
	}
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

func TestIdentifiersExpression(t *testing.T) {
	input := "foobar"

	l := gomonkey.NewLexer(input)
	p := gomonkey.NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*gomonkey.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*gomonkey.Indentifier)
	if !ok {
		t.Fatalf("exp not *Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := gomonkey.NewLexer(input)
	p := gomonkey.NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*gomonkey.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*gomonkey.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *Identifier. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("ident.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", literal.TokenLiteral())
	}
}

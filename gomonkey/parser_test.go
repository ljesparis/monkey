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

	testIntegerLiteral(t, stmt.Expression, 5)
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := gomonkey.NewLexer(tt.input)
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

		exp, ok := stmt.Expression.(*gomonkey.PrefixExpression)
		if !ok {
			t.Fatalf("exp not *PrefixExpreesion. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Errorf("ident.Value not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}

	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := gomonkey.NewLexer(tt.input)
		p := gomonkey.NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements has not enough statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*gomonkey.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*gomonkey.InfixExpression)
		if !ok {
			t.Fatalf("exp not *InfixExpreesion. got=%T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Errorf("ident.Value not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il gomonkey.Expression, value int64) bool {
	integ, ok := il.(*gomonkey.IntegerLiteral)
	if !ok {
		t.Errorf("il not *IntegerLiteral. got=%T", il)
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false

	}

	return true
}

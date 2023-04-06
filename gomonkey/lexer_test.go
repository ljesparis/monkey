package gomonkey_test

import (
	"testing"

	"github.com/ljesparis/monkey/gomonkey"
)


func TestNextTokenShouldBeOk(t *testing.T) {

    tests := []struct {
        source string
        expected []gomonkey.Token
    } {
        {
            `=+(){},;-/*!<>==!=`,
            []gomonkey.Token{
                {gomonkey.ASSIGN, "="},
                {gomonkey.PLUS, "+"},
                {gomonkey.LPAREN, "("},
                {gomonkey.RPAREN, ")"},
                {gomonkey.LBRACE, "{"},
                {gomonkey.RBRACE, "}"},
                {gomonkey.COMMA, ","},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.MINUS, "-"},
                {gomonkey.SLASH, "/"},
                {gomonkey.ASTERISK, "*"},
                {gomonkey.BANG, "!"},
                {gomonkey.LESS_THAN, "<"},
                {gomonkey.GREAT_THAN, ">"},
                {gomonkey.EQUAL, "=="},             
                {gomonkey.NOT_EQUAL, "!="},             
                {gomonkey.EOF, ""},
            },
        },
        {
            `let _this_is_a_variable = 100000;`,
            []gomonkey.Token{
                {gomonkey.LET, "let"},
                {gomonkey.IDENTIFIER, "_this_is_a_variable"},
                {gomonkey.ASSIGN, "="},
                {gomonkey.INT, "100000"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.EOF, ""},
            },
        },
        {
            `$

            5 == 10;
            10 != 5;
            `,
            []gomonkey.Token{
                {gomonkey.ILLEGAL, "$"},
                {gomonkey.INT, "5"},             
                {gomonkey.EQUAL, "=="},             
                {gomonkey.INT, "10"},             
                {gomonkey.SEMICOLON, ";"},             
                {gomonkey.INT, "10"},             
                {gomonkey.NOT_EQUAL, "!="},             
                {gomonkey.INT, "5"},             
                {gomonkey.SEMICOLON, ";"},             
                {gomonkey.EOF, ""},
            },
        },
        {
            `let five = 5;
            let ten = 10;

            let add = fn(x,y) {
            x + y
            };

            let result = add(five, ten);
            */!5;

            5 < 10 > 5;

            if (5 < 10) {
                return false;
            } else {
                return true;
            }
            `,
            []gomonkey.Token{
                {gomonkey.LET, "let"},
                {gomonkey.IDENTIFIER, "five"},
                {gomonkey.ASSIGN, "="},
                {gomonkey.INT, "5"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.LET, "let"},
                {gomonkey.IDENTIFIER, "ten"},
                {gomonkey.ASSIGN, "="},
                {gomonkey.INT, "10"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.LET, "let"},
                {gomonkey.IDENTIFIER, "add"},
                {gomonkey.ASSIGN, "="},
                {gomonkey.FUNCTION, "fn"},
                {gomonkey.LPAREN, "("},
                {gomonkey.IDENTIFIER, "x"},
                {gomonkey.COMMA, ","},
                {gomonkey.IDENTIFIER, "y"},
                {gomonkey.RPAREN, ")"},
                {gomonkey.LBRACE, "{"},
                {gomonkey.IDENTIFIER, "x"},
                {gomonkey.PLUS, "+"},
                {gomonkey.IDENTIFIER, "y"},
                {gomonkey.RBRACE, "}"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.LET, "let"},
                {gomonkey.IDENTIFIER, "result"},
                {gomonkey.ASSIGN, "="},
                {gomonkey.IDENTIFIER, "add"},
                {gomonkey.LPAREN, "("},
                {gomonkey.IDENTIFIER, "five"},
                {gomonkey.COMMA, ","},
                {gomonkey.IDENTIFIER, "ten"},
                {gomonkey.RPAREN, ")"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.ASTERISK, "*"},
                {gomonkey.SLASH, "/"},
                {gomonkey.BANG, "!"},
                {gomonkey.INT, "5"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.INT, "5"},
                {gomonkey.LESS_THAN, "<"},
                {gomonkey.INT, "10"},
                {gomonkey.GREAT_THAN, ">"},
                {gomonkey.INT, "5"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.IF, "if"},
                {gomonkey.LPAREN, "("},
                {gomonkey.INT, "5"},
                {gomonkey.LESS_THAN, "<"},
                {gomonkey.INT, "10"},
                {gomonkey.RPAREN, ")"},
                {gomonkey.LBRACE, "{"},
                {gomonkey.RETURN, "return"},
                {gomonkey.FALSE, "false"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.RBRACE, "}"},
                {gomonkey.ELSE, "else"},
                {gomonkey.LBRACE, "{"},
                {gomonkey.RETURN, "return"},
                {gomonkey.TRUE, "true"},
                {gomonkey.SEMICOLON, ";"},
                {gomonkey.RBRACE, "}"},
                {gomonkey.EOF, ""},
            },
        },
    }

    for _, test := range tests {
        lexer := gomonkey.NewLexer(test.source)

        for i, tt := range test.expected {
            tok := lexer.NextToken()

            if tok.Type != tt.Type {
                t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.Type, tok.Type)
            }

            if tok.Literal != tt.Literal {
                t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.Literal, tok.Literal)
            }
        }
    }


}


package monkey_test

import (
	"testing"

	"github.com/ljesparis/monkey"
)

func TestNextTokenShouldBeOk(t *testing.T) {

	tests := []struct {
		source   string
		expected []monkey.Token
	}{
		{
			`=+(){},;-/*!<>==!=`,
			[]monkey.Token{
				{monkey.ASSIGN, "="},
				{monkey.PLUS, "+"},
				{monkey.LPAREN, "("},
				{monkey.RPAREN, ")"},
				{monkey.LBRACE, "{"},
				{monkey.RBRACE, "}"},
				{monkey.COMMA, ","},
				{monkey.SEMICOLON, ";"},
				{monkey.MINUS, "-"},
				{monkey.SLASH, "/"},
				{monkey.ASTERISK, "*"},
				{monkey.BANG, "!"},
				{monkey.LESS_THAN, "<"},
				{monkey.GREAT_THAN, ">"},
				{monkey.EQUAL, "=="},
				{monkey.NOT_EQUAL, "!="},
				{monkey.EOF, ""},
			},
		},
		{
			`let _this_is_a_variable = 100000;`,
			[]monkey.Token{
				{monkey.LET, "let"},
				{monkey.IDENTIFIER, "_this_is_a_variable"},
				{monkey.ASSIGN, "="},
				{monkey.INT, "100000"},
				{monkey.SEMICOLON, ";"},
				{monkey.EOF, ""},
			},
		},
		{
			`$

            5 == 10;
            10 != 5;
            `,
			[]monkey.Token{
				{monkey.ILLEGAL, "$"},
				{monkey.INT, "5"},
				{monkey.EQUAL, "=="},
				{monkey.INT, "10"},
				{monkey.SEMICOLON, ";"},
				{monkey.INT, "10"},
				{monkey.NOT_EQUAL, "!="},
				{monkey.INT, "5"},
				{monkey.SEMICOLON, ";"},
				{monkey.EOF, ""},
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
			[]monkey.Token{
				{monkey.LET, "let"},
				{monkey.IDENTIFIER, "five"},
				{monkey.ASSIGN, "="},
				{monkey.INT, "5"},
				{monkey.SEMICOLON, ";"},
				{monkey.LET, "let"},
				{monkey.IDENTIFIER, "ten"},
				{monkey.ASSIGN, "="},
				{monkey.INT, "10"},
				{monkey.SEMICOLON, ";"},
				{monkey.LET, "let"},
				{monkey.IDENTIFIER, "add"},
				{monkey.ASSIGN, "="},
				{monkey.FUNCTION, "fn"},
				{monkey.LPAREN, "("},
				{monkey.IDENTIFIER, "x"},
				{monkey.COMMA, ","},
				{monkey.IDENTIFIER, "y"},
				{monkey.RPAREN, ")"},
				{monkey.LBRACE, "{"},
				{monkey.IDENTIFIER, "x"},
				{monkey.PLUS, "+"},
				{monkey.IDENTIFIER, "y"},
				{monkey.RBRACE, "}"},
				{monkey.SEMICOLON, ";"},
				{monkey.LET, "let"},
				{monkey.IDENTIFIER, "result"},
				{monkey.ASSIGN, "="},
				{monkey.IDENTIFIER, "add"},
				{monkey.LPAREN, "("},
				{monkey.IDENTIFIER, "five"},
				{monkey.COMMA, ","},
				{monkey.IDENTIFIER, "ten"},
				{monkey.RPAREN, ")"},
				{monkey.SEMICOLON, ";"},
				{monkey.ASTERISK, "*"},
				{monkey.SLASH, "/"},
				{monkey.BANG, "!"},
				{monkey.INT, "5"},
				{monkey.SEMICOLON, ";"},
				{monkey.INT, "5"},
				{monkey.LESS_THAN, "<"},
				{monkey.INT, "10"},
				{monkey.GREAT_THAN, ">"},
				{monkey.INT, "5"},
				{monkey.SEMICOLON, ";"},
				{monkey.IF, "if"},
				{monkey.LPAREN, "("},
				{monkey.INT, "5"},
				{monkey.LESS_THAN, "<"},
				{monkey.INT, "10"},
				{monkey.RPAREN, ")"},
				{monkey.LBRACE, "{"},
				{monkey.RETURN, "return"},
				{monkey.FALSE, "false"},
				{monkey.SEMICOLON, ";"},
				{monkey.RBRACE, "}"},
				{monkey.ELSE, "else"},
				{monkey.LBRACE, "{"},
				{monkey.RETURN, "return"},
				{monkey.TRUE, "true"},
				{monkey.SEMICOLON, ";"},
				{monkey.RBRACE, "}"},
				{monkey.EOF, ""},
			},
		},
	}

	for _, test := range tests {
		lexer := monkey.NewLexer(test.source)

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

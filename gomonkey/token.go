package gomonkey


const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    IDENTIFIER = "INDENTIFIER"
    INT = "INT"

    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    DIV = "/"
    MULT = "*"

    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    FUNCTION = "FUNCTION"
    LET = "LET"
)


type TokenType string

type Token struct {
    Type TokenType
    Literal string
}


func newToken(t TokenType, l byte) Token {
    return Token{
        Type: t,
        Literal: string(l),
    }
}


var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENTIFIER
}


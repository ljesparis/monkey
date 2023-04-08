package monkey

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "INDENTIFIER"
	INT   = "INT"

	BANG     = "!"
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	LESS_THAN  = "<"
	GREAT_THAN = ">"
	LEQ_THAN   = "<="
	GEQ_THAN   = ">="
	EQ         = "=="
	NOT_EQ     = "!="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

func newToken(t TokenType, l byte) Token {
	return Token{
		Type:    t,
		Literal: string(l),
	}
}

func eofToken() Token {
	return Token{Type: EOF, Literal: ""}
}

var (
	keywords = map[string]TokenType{
		"fn":     FUNCTION,
		"let":    LET,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
		"true":   TRUE,
		"false":  FALSE,
	}

	operators = map[string]TokenType{
		"+":  PLUS,
		"-":  MINUS,
		"/":  SLASH,
		"*":  ASTERISK,
		"<":  LESS_THAN,
		">":  GREAT_THAN,
		"!":  BANG,
		"=":  ASSIGN,
		"==": EQ,
		"!=": NOT_EQ,
		"<=": LEQ_THAN,
		">=": GEQ_THAN,
		";":  SEMICOLON,
		"(":  LPAREN,
		")":  RPAREN,
		"{":  LBRACE,
		"}":  RBRACE,
		",":  COMMA,
	}
)

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func lookupOperator(op string) TokenType {
	if tok, ok := operators[op]; ok {
		return tok
	}
	return ILLEGAL
}

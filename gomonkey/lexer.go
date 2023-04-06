package gomonkey

type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
}


func NewLexer(in string) *Lexer {
    lex := &Lexer {input: in}
    lex.readChar()
    return lex
}

func (l* Lexer) NextToken() Token {
    var tok Token

    l.skipWhitespace()

    switch l.ch {
    case '=':
        tok = newToken(ASSIGN, l.ch)
    case ';':
        tok = newToken(SEMICOLON, l.ch)
    case '(':
        tok = newToken(LPAREN, l.ch)
    case ')':
        tok = newToken(RPAREN, l.ch)
    case '{':
        tok = newToken(LBRACE, l.ch)
    case '}':
        tok = newToken(RBRACE, l.ch)
    case ',':
        tok = newToken(COMMA, l.ch)
    case '+':
        tok = newToken(PLUS, l.ch)
    case '-':
        tok = newToken(MINUS, l.ch)
    case '/':
        tok = newToken(DIV, l.ch)
    case '*':
        tok = newToken(MULT, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = EOF
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier() 
            tok.Type = LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Type = INT
            tok.Literal = l.readNumber()
            return tok
        } else {
            tok = newToken(ILLEGAL, l.ch)
        }
    }

    l.readChar()

    return tok
}

func (l* Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }

    l.position = l.readPosition
    l.readPosition += 1
}

func (l* Lexer) readNumber() string {
    pos := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[pos:l.position]
}

func (l* Lexer) readIdentifier() string {
    pos := l.position
    for isLetter(l.ch) {
        l.readChar()
    }

    return l.input[pos:l.position]
}


func (l* Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

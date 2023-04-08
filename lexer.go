package monkey

type Lexer struct {
	input            string
	currenctPosition int
	nextPosition     int
	ch               byte
}

func NewLexer(in string) *Lexer {
	lex := &Lexer{input: in}
	lex.readChar()
	return lex
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case 0:
		tok = eofToken()
	case '!', '=':
		if l.peakChar() == '=' {
			last_ch := l.ch
			l.readChar()
			tok.Literal = string(last_ch) + string(l.ch)
			tok.Type = lookupOperator(tok.Literal)
		} else {
			tt := lookupOperator(string(l.ch))
			tok = newToken(tt, l.ch)
		}
	case '+', '-', '/', '*', '<', '>', ';', '(', ')', '{', '}', ',':
		tt := lookupOperator(string(l.ch))
		tok = newToken(tt, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
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

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}

	l.currenctPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) readNumber() string {
	pos := l.currenctPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.currenctPosition]
}

func (l *Lexer) readIdentifier() string {
	pos := l.currenctPosition
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.currenctPosition]
}

func (l *Lexer) peakChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) skipWhitespace() {
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

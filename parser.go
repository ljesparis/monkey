package monkey

import (
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)

var precedences = map[TokenType]int{
	EQ:         EQUALS,
	NOT_EQ:     EQUALS,
	LESS_THAN:  LESSGREATER,
	GREAT_THAN: LESSGREATER,
	PLUS:       SUM,
	MINUS:      SUM,
	SLASH:      PRODUCT,
	ASTERISK:   PRODUCT,
}

type Parser struct {
	l      *Lexer
	errors []string

	curToken  Token
	peekToken Token

	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[TokenType]prefixParseFn)

	p.registerPrefix(IDENT, p.parseIdentifier)
	p.registerPrefix(INT, p.parseIntegerLiteral)
	p.registerPrefix(BANG, p.parsePrefixExpression)
	p.registerPrefix(MINUS, p.parsePrefixExpression)

	p.infixParseFns = make(map[TokenType]infixParseFn)
	p.registerInfix(MINUS, p.parseInfixExpression)
	p.registerInfix(PLUS, p.parseInfixExpression)
	p.registerInfix(ASTERISK, p.parseInfixExpression)
	p.registerInfix(SLASH, p.parseInfixExpression)
	p.registerInfix(LESS_THAN, p.parseInfixExpression)
	p.registerInfix(GREAT_THAN, p.parseInfixExpression)
	p.registerInfix(EQ, p.parseInfixExpression)
	p.registerInfix(NOT_EQ, p.parseInfixExpression)

	return p
}

func (p *Parser) registerPrefix(tt TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}

func (p *Parser) registerInfix(tt TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}

func (p *Parser) noPrefixParseFnError(t TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIdentifier() Expression {
	return &Indentifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekErrors(tt TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tt, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case LET:
		return p.parseLetStatement()
	case RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() Statement {
	stmt := &ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedense int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(SEMICOLON) && precedense < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{Token: p.curToken, Operator: p.curToken.Literal}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseLetStatement() Statement {
	stmt := &LetStatement{Token: p.curToken}

	if !p.expectPeek(IDENT) {
		return nil
	}

	stmt.Name = &Indentifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(ASSIGN) {
		return nil
	}

	// skip tokens until semicolon
	for !p.curTokenIs(SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() Statement {
	stmt := ReturnStatement{Token: p.curToken}

	p.nextToken()

	// skip tokens until semicolon
	for !p.curTokenIs(SEMICOLON) {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) curTokenIs(tt TokenType) bool {
	return p.curToken.Type == tt
}

func (p *Parser) peekTokenIs(tt TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) expectPeek(tt TokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken()
		return true
	}

	p.peekErrors(tt)
	return true
}

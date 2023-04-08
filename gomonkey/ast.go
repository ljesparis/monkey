package gomonkey

import "bytes"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

type LetStatement struct {
	Token Token
	Name  *Indentifier
	Value Expression
}

func (lt *LetStatement) statementNode() {}
func (lt *LetStatement) TokenLiteral() string {
	return lt.Token.Literal
}

func (lt *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(lt.TokenLiteral() + " ")
	out.WriteString(lt.Name.String())
	out.WriteString(" = ")

	if lt.Value != nil {
		out.WriteString(lt.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Indentifier struct {
	Token Token
	Value string
}

func (i *Indentifier) expressionNode() {}
func (i *Indentifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Indentifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

func (i *PrefixExpression) expressionNode() {}
func (i *PrefixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Operator)
	out.WriteString(i.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    Token
	Operator string
	Left     Expression
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")

	return out.String()
}

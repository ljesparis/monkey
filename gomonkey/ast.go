package gomonkey

type Node interface {
	TokenLiteral() string
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

type LetStatement struct {
	Token Token
	Name  *Indentifier
	Value Expression
}

func (lt *LetStatement) statementNode() {}
func (lt *LetStatement) TokenLiteral() string {
	return lt.Token.Literal
}

type Indentifier struct {
	Token Token
	Value string
}

func (i *Indentifier) expressionNode() {}
func (i *Indentifier) TokenLiteral() string {
	return i.Token.Literal
}

type RunStatement struct {
	Tk          Token
	ReturnValue Expression
}

func (rs *RunStatement) statementNode() {}
func (rs *RunStatement) TokenLiteral() string {
	return rs.Tk.Literal
}

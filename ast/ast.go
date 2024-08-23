package ast

import "bytes"

// base nodei interface
type NodeI interface {
	TokenLiteral() string
	String() string
}

// all statements implement this interface
type StatementI interface {
	NodeI
	statementNode()
}

// all expressions implement this interface
type ExpressionI interface {
	NodeI
	expressionNode()
}

type Program struct {
	Statements []StatementI
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}

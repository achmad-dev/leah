package ast

import (
	"bytes"

	"github.com/achmad-dev/leah/token"
)

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

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type Program struct {
	Statements []StatementI
}

func (p *Program) TokenLiteral() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.TokenLiteral())
		out.WriteString("\n")
	}
	return out.String()
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}

type VarStatement struct {
	Value ExpressionI
	Name  *Identifier
	Token token.Token
}

func (v *VarStatement) TokenLiteral() string {
	return v.Token.Literal
}

func (v *VarStatement) statementNode() {}

func (v *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(v.TokenLiteral() + " ")
	out.WriteString(v.Name.String())
	out.WriteString(" = ")
	if v.Value != nil {
		out.WriteString(v.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// return statement
type ReturnStatement struct {
	Value ExpressionI
	Token token.Token
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(r.TokenLiteral() + " ")
	if r.Value != nil {
		out.WriteString(r.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// expression statement
type ExpressionStatement struct {
	Expression ExpressionI
	Token      token.Token
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

// block statement
type BlockStatement struct {
	Statements []StatementI
	Token      token.Token
}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BlockStatement) statementNode() {}

func (b *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}

// boolean

type Boolean struct {
	Value bool
	Token token.Token
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

// integer

type Integer struct {
	Value int64
	Token token.Token
}

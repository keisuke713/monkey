package parser

import (
	"github.com/keisuke713/monkey/ast"
	"github.com/keisuke713/monkey/lexer"
	"github.com/keisuke713/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken  = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}
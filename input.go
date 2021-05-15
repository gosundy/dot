package dot

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gosundy/dot/parser"
)

func Input(input []byte) (*DotVisitor, func() *Value) {
	is := antlr.NewInputStream(string(input))
	// Create the Lexer
	lexer := parser.NewDotLexer(is)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewDotParser(tokenStream)
	p.BuildParseTrees = true
	v := NewDotVisitor()
	//Start is rule name of Calc.g4
	calcfunc := p.Start().Accept(v)
	return v, calcfunc.(func() *Value)
}
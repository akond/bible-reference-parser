package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ParseBibleText(s string, walker Walker, disableErrors bool) string {
	input := antlr.NewInputStream(s)
	lexer := NewBibleLexer(input)
	//lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewBibleParser(stream)
	if disableErrors {
		//p.RemoveErrorListeners()
	}

	//errorListener := NewRecoveringErrorListener()
	//p.AddErrorListener(errorListener)
	//p.BuildParseTrees = true
	tree := p.R()
	listener := NewTreeShapeListener(walker, nil)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.collector
}

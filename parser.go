package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)


func ParseBibleText(text string, walker Walker, disableErrors bool) string {
	input := antlr.NewInputStream(text)
	lexer := NewBibleLexer(input)
	//lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, 0)
	bibleParser := NewBibleParser(stream)
	if disableErrors {
		//bibleParser.RemoveErrorListeners()
	}

	//errorListener := NewRecoveringErrorListener()
	//bibleParser.AddErrorListener(errorListener)
	//bibleParser.BuildParseTrees = true
	tree := bibleParser.R()
	listener := NewTreeShapeListener(walker, nil)
	//listener := NewFixerListener(bibleParser)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.collector
}

package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)



func createTokenStream(text string) *antlr.CommonTokenStream {
	input := antlr.NewInputStream(text)
	lexer := NewBibleLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return stream
}

func createBibleParser(text string) *BibleParser {
	stream := createTokenStream(text)
	bibleParser := NewBibleParser(stream)
	return bibleParser
}

func SubstituteBibleRefWithStar(text string) string {
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, createBibleParser(text).R())
	return listener.collector
}


func SubstituteBibleRefWithXml(text string) string {
	listener := NewReferenceSubstitutionListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, createBibleParser(text).R())
	return listener.collector
}

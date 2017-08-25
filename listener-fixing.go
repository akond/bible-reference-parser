package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"regexp"
)

type FixerListener struct {
	*BaseBibleListener
	parser    antlr.Parser
	collector string
	walker    Walker
	//errorListener *RecoveringErrorListener
	errorListener *antlr.DefaultErrorListener
	errorContext  antlr.ParserRuleContext
}

func NewFixerListener(parser antlr.Parser) *FixerListener {
	listener := new(FixerListener)
	listener.parser = parser
	return listener
}

func (this *FixerListener) EnterVersespan(ctx *VersespanContext) {
	s := regexp.MustCompile("\\D").Split(ctx.GetText(), 2)
	if s[0] >= s[1] {
		panic("a < b: " + ctx.GetText())
	}
}

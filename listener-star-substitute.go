package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*BaseBibleListener
	collector     string
	//errorListener *RecoveringErrorListener
	errorListener *antlr.DefaultErrorListener
	errorContext  antlr.ParserRuleContext
}

func NewTreeShapeListener() *TreeShapeListener {
	listener := new(TreeShapeListener)
	listener.errorListener = antlr.NewDefaultErrorListener()
	listener.errorContext = nil
	return listener
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	ctx.EnterRule(this)
	switch ctx.GetRuleIndex() {
	case BibleParserRULE_reference:
		this.collector += "*"

	case BibleParserRULE_text:
		this.collector += ctx.GetText()
	}
}

func (this *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	ctx.ExitRule(this)
}


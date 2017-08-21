package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Walker func(this *TreeShapeListener, ctx antlr.ParserRuleContext);

type TreeShapeListener struct {
	*BaseBibleListener
	collector     string
	walker        Walker
	//errorListener *RecoveringErrorListener
	errorListener *antlr.DefaultErrorListener
	errorContext  antlr.ParserRuleContext
}

func NewTreeShapeListener(walker Walker, errorListener *RecoveringErrorListener) *TreeShapeListener {
	listener := new(TreeShapeListener)
	listener.walker = walker
	//listener.errorListener = errorListener
	listener.errorListener = antlr.NewDefaultErrorListener()
	listener.errorContext = nil
	return listener
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(this.errorListener.InErrorMode);

	this.walker(this, ctx)

	//if this.errorListener.InErrorMode == 0 {
	//	//if "parser.ReferenceContext" ==
	//	this.walker(this, ctx)
	//	//fmt.Println("EnterEveryRule >> ", ctx.GetRuleIndex(), this.errorListener.InErrorMode, ctx.GetText())
	//	//fmt.Println(ctx.GetChildCount())
	//
	//} else if this.errorListener.InErrorMode == 1 {
	//	//fmt.Println(">> ", ctx.GetText());
	//	this.errorListener.InErrorMode ++
	//	this.collector += ctx.GetText()
	//	this.errorContext = ctx
	//}
}

func (this *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("ExitEveryRule >-> ", ctx.GetText())
	//fmt.Println("UT", &ctx)
	//fmt.Println(reflect.TypeOf(this.errorContext).String())

	//fmt.Println(this.errorContext.GetRuleIndex())
	//fmt.Println("------", this.errorContext.GetRuleIndex())
	//fmt.Println("------", ctx.GetRuleIndex())

	if this.errorContext != nil && ctx.GetRuleIndex() == this.errorContext.GetRuleIndex() {
		//fmt.Println("-> reset to no error mode")
		//this.errorListener.InErrorMode = 0
	}
}


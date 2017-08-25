package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	//"fmt"
	"fmt"
	"reflect"
)

type ActualVisitor struct {
	BaseBibleVisitor
}

func (this *ActualVisitor) contextName(ctx interface{}) string {
	name := reflect.ValueOf(ctx).Type().String()[8:]
	return name[:len(name)-7]
}

func (this *ActualVisitor) VisitChildren(node antlr.RuleNode) {
	for _, n := range node.GetChildren() {
		//name := reflect.ValueOf(n).Type().String()[8:]
		//fmt.Println(name[:len(name)-7])
		methodName := "Visit" + this.contextName(n)
		//methodName := "Visit" + name[:len(name)-7]
		fmt.Println(methodName)
		reflect.ValueOf(this).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(n)})
		//fmt.Println(reflect.ValueOf(n.GetPayload()).Type().String())
		//fmt.Println(reflect.ValueOf(this).MethodByName("VisitReference1Context"))
	}
}

func (this *ActualVisitor) VisitReference(ctx *ReferenceContext) {
	fmt.Println("!Reference")
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitReference1(ctx *Reference1Context) {
	fmt.Println("!Reference 1")
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitManychaptersbook(ctx *ManychaptersbookContext) {
	fmt.Println(this.contextName(ctx.GetChild(0)))
	//this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitBook1(ctx *Book1Context) {
	fmt.Println("!BOOK1")
	fmt.Println("-> " + ctx.GetText() + ".")

	this.BaseBibleVisitor.VisitBook1(ctx)
}

func (this *ActualVisitor) VisitBook2(ctx *Book2Context) {
	fmt.Println("!BOOK2")
	fmt.Println("-> " + ctx.GetText() + ".")
	fmt.Println(ctx.GetStart())
	this.BaseBibleVisitor.VisitBook2(ctx)
}

type ReferenceSubstitutionListener struct {
	*BaseBibleListener
	collector string
	//errorListener *RecoveringErrorListener
	errorListener *antlr.DefaultErrorListener
	errorContext  antlr.ParserRuleContext
}

func NewReferenceSubstitutionListener() *ReferenceSubstitutionListener {
	listener := new(ReferenceSubstitutionListener)
	listener.errorListener = antlr.NewDefaultErrorListener()
	listener.errorContext = nil
	return listener
}

func (this *ReferenceSubstitutionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	ctx.EnterRule(this)
	switch ctx.GetRuleIndex() {
	case BibleParserRULE_reference:
		visitor := new(ActualVisitor)
		visitor.VisitReference(ctx.(*ReferenceContext))
		this.collector += "~"

	case BibleParserRULE_text:
		this.collector += ctx.GetText()
	}
}

func (this *ReferenceSubstitutionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	ctx.ExitRule(this)
}

package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"reflect"
	"strconv"
	"regexp"
)

var parserBookRe *regexp.Regexp

func init() {
	parserBookRe = regexp.MustCompile("\\*parser\\.Book(\\d+)")
}

type Part struct {
	ChapterStart int
	VerseStart   int
	ChapterEnd   int
	VerseEnd     int
	Text         string
}

type IHasChapterValue interface {
	SetChapterValue(int)
}

func (p Part) String() string {
	if p.Text != "" {
		return p.Text
	} else if p.ChapterEnd != 0 || p.VerseEnd != 0 {
		return fmt.Sprintf("{%d:%d %d:%d}", p.ChapterStart, p.VerseStart, p.ChapterEnd, p.VerseEnd)
	} else {
		return fmt.Sprintf("{%d:%d}", p.ChapterStart, p.VerseStart)
	}
}

type Reference struct {
	Book  int
	Parts []*Part
}

func (this *Reference) IsEmpty() bool {
	return len(this.Parts) == 0
}

func (this *Reference) String() string {
	return fmt.Sprintf("{%d %s}", this.Book, this.Parts)
}

func (this *Reference) AppendPart(c1, v1, c2, v2 int) {
	p := new(Part)
	p.ChapterStart = c1
	p.VerseStart = v1
	p.ChapterEnd = c2
	p.VerseEnd = v2
	this.Parts = append(this.Parts, p)
}

func (this *Reference) AppendPartSingle(c1, v1 int) {
	p := new(Part)
	p.ChapterStart = c1
	p.VerseStart = v1
	this.Parts = append(this.Parts, p)
}

func (this *Reference) AppendText(s string) {
	p := new(Part)
	p.Text = s
	this.Parts = append(this.Parts, p)
}

var TokenToBook = map[int]int{BibleParserRULE_book1: 1,
	BibleParserRULE_book2: 2,
	BibleParserRULE_book3: 3,
	BibleParserRULE_book4: 4,
	BibleParserRULE_book5: 5,
	BibleParserRULE_book6: 6,
	BibleParserRULE_book7: 7,
	BibleParserRULE_book8: 8,
	BibleParserRULE_book9: 9,
	BibleParserRULE_book10: 10,
	BibleParserRULE_book11: 11,
	BibleParserRULE_book12: 12,
	BibleParserRULE_book13: 13,
	BibleParserRULE_book14: 14,
	BibleParserRULE_book15: 15,
	BibleParserRULE_book16: 16,
	BibleParserRULE_book17: 17,
	BibleParserRULE_book18: 18,
	BibleParserRULE_book19: 19,
	BibleParserRULE_book20: 20,
	BibleParserRULE_book21: 21,
	BibleParserRULE_book22: 22,
	BibleParserRULE_book23: 23,
	BibleParserRULE_book24: 24,
	BibleParserRULE_book25: 25,
	BibleParserRULE_book26: 26,
	BibleParserRULE_book27: 27,
	BibleParserRULE_book28: 28,
	BibleParserRULE_book29: 29,
	BibleParserRULE_book30: 30,
	BibleParserRULE_book31: 31,
	BibleParserRULE_book32: 32,
	BibleParserRULE_book33: 33,
	BibleParserRULE_book34: 34,
	BibleParserRULE_book35: 35,
	BibleParserRULE_book36: 36,
	BibleParserRULE_book37: 37,
	BibleParserRULE_book38: 38,
	BibleParserRULE_book39: 39,
	BibleParserRULE_book40: 40,
	BibleParserRULE_book41: 41,
	BibleParserRULE_book42: 42,
	BibleParserRULE_book43: 43,
	BibleParserRULE_book44: 44,
	BibleParserRULE_book45: 45,
	BibleParserRULE_book46: 46,
	BibleParserRULE_book47: 47,
	BibleParserRULE_book48: 48,
	BibleParserRULE_book49: 49,
	BibleParserRULE_book50: 50,
	BibleParserRULE_book51: 51,
	BibleParserRULE_book52: 52,
	BibleParserRULE_book53: 53,
	BibleParserRULE_book54: 54,
	BibleParserRULE_book55: 55,
	BibleParserRULE_book56: 56,
	BibleParserRULE_book57: 57,
	BibleParserRULE_book58: 58,
	BibleParserRULE_book59: 59,
	BibleParserRULE_book60: 60,
	BibleParserRULE_book61: 61,
	BibleParserRULE_book62: 62,
	BibleParserRULE_book63: 63,
	BibleParserRULE_book64: 64,
	BibleParserRULE_book65: 65,
	BibleParserRULE_book66: 66,
}

type ChildCallback func(ctx antlr.ParserRuleContext);

func loopChildren(f ChildCallback, node antlr.RuleNode) {
	if f == nil {
		return
	}

	for _, child := range node.GetChildren() {
		if ctx, ok := child.(antlr.ParserRuleContext); ok {
			f(ctx)
		}
	}
}

type ReferenceNode struct {
	Type string
	Val  int
}

func NewRefNode(t string, v int) *ReferenceNode {
	var r = new(ReferenceNode)
	r.Val = v
	r.Type = t
	return r
}

type ActualVisitor struct {
	BaseBibleVisitor
	stack  []*ReferenceNode
	Result *Reference
}

func (this *ActualVisitor) contextName(ctx interface{}) string {
	name := reflect.ValueOf(ctx).Type().String()[8:]
	return name[:len(name)-7]
}

func (this *ActualVisitor) VisitChildren(node antlr.RuleNode) {
	for _, n := range node.GetChildren() {
		i := reflect.ValueOf(n).Type().String()
		if i[len(i)-7:] == "Context" {
			methodName := "Visit" + this.contextName(n)
			//fmt.Println(" -> " + methodName)
			reflect.ValueOf(this).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(n)})
		}
	}
}

func (this *ActualVisitor) VisitReference(ctx *ReferenceContext) {
	this.VisitChildren(ctx)

	//for _, n := range this.stack {
	//	switch n.Type {
	//	case "book":
	//		this.Result.Book = n.Val
	//	}
	//}

	this.stack = make([]*ReferenceNode, 0)
}

func (this *ActualVisitor) VisitManychaptersbook(ctx *ManychaptersbookContext) {
	for _, n := range ctx.GetChildren() {
		typeName := reflect.ValueOf(n).Type().String()
		if match := parserBookRe.FindStringSubmatch(typeName); match != nil {
			book, _ := strconv.Atoi(match[1])
			this.Result.Book = book
			//this.stack = append(this.stack, NewRefNode("book", book))
		}
	}
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitSinglechapterbook(ctx *SinglechapterbookContext) {
	for _, n := range ctx.GetChildren() {
		typeName := reflect.ValueOf(n).Type().String()
		if match := parserBookRe.FindStringSubmatch(typeName); match != nil {
			book, _ := strconv.Atoi(match[1])
			this.Result.Book = book
			//this.stack = append(this.stack, NewRefNode("book", book))
			this.stack = append(this.stack, NewRefNode("chapter", 1))
		}
	}
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitParts(ctx *PartsContext) {
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitPart(ctx *PartContext) {
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitChapter(ctx *ChapterContext) {
	i, _ := strconv.Atoi(ctx.GetText())
	this.stack = append(this.stack, NewRefNode("chapter", i))
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitSpanlist(ctx *SpanlistContext) {
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitChapterverse(ctx *ChapterverseContext) {
	this.VisitChildren(ctx)
}

func (this *ActualVisitor) VisitVerse(ctx *VerseContext) {
	this.VisitChildren(ctx)
}

/**
 * ref1: chapterverse MINUS chapterverse;
 */
func (this *ActualVisitor) VisitRef1(ctx *Ref1Context) {
	this.VisitChildren(ctx)
	slice := this.stack[len(this.stack)-4:]

	chapter1 := slice[0].Val
	verse1 := slice[1].Val

	chapter2 := slice[2].Val
	verse2 := slice[3].Val

	this.Result.AppendPart(chapter1, verse1, chapter1, maxVerseNumber(this.Result.Book, chapter1))
	for chapter:= chapter1 + 1; chapter < chapter2; chapter ++ {
		this.Result.AppendPart(chapter, 1, chapter, maxVerseNumber(this.Result.Book, chapter))
	}
	this.Result.AppendPart(chapter2, 1, chapter2, verse2)
	
	this.stack = this.stack[0:len(this.stack)-4]
}

// ref2: chapter ':' spanlist;
func (this *ActualVisitor) VisitRef2(ctx *Ref2Context) {
	this.VisitChildren(ctx)

	var chapter, verse int
	for _, n := range this.stack {
		switch n.Type {
		case "chapter":
			chapter = n.Val

		case "verse":
			this.Result.AppendPartSingle(chapter, n.Val)
			n.Type = "void"

		case "verse1":
			verse = n.Val
			n.Type = "void"

		case "verse2":
			this.Result.AppendPart(chapter, verse, chapter, n.Val)
			n.Type = "void"
		}
	}

}

// ref3: chapter MINUS chapter;
func (this *ActualVisitor) VisitRef3(ctx *Ref3Context) {
	this.VisitChildren(ctx)

	slice := this.stack[len(this.stack)-2:]
	for chapter := slice[0].Val; chapter <= slice[1].Val; chapter++ {
		this.Result.AppendPart(chapter, 1, chapter, maxVerseNumber(this.Result.Book, chapter))
	}
	this.stack = this.stack[0:len(this.stack)-2]
}

// ref4: chapter;
func (this *ActualVisitor) VisitRef4(ctx *Ref4Context) {
	this.VisitChildren(ctx)

	slice := this.stack[len(this.stack)-1:]
	chapter := slice[0].Val
	this.Result.AppendPart(chapter, 1, chapter, maxVerseNumber(this.Result.Book, chapter))
	this.stack = this.stack[0:len(this.stack)-1]
}

func (this *ActualVisitor) VisitRef5(ctx *Ref5Context) {
	this.VisitChildren(ctx)

	var verse int
	for _, n := range this.stack {
		switch n.Type {
		case "verse":
			this.Result.AppendPartSingle(1, n.Val)

		case "verse1":
			verse = n.Val

		case "verse2":
			this.Result.AppendPart(1, verse, 1, n.Val)
		}
	}
}

func (this *ActualVisitor) VisitTerminator(ctx *TerminatorContext) {
	this.Result.AppendText(ctx.GetText())
}

func (this *ActualVisitor) VisitVersespan(ctx *VersespanContext) {
	this.VisitChildren(ctx)

	verses := ctx.GetTypedRuleContexts(reflect.TypeOf(new(VerseContext)))
	i1, _ := strconv.Atoi(verses[0].GetText())
	i2, _ := strconv.Atoi(verses[1].GetText())

	this.stack = append(this.stack, NewRefNode("verse1", i1), NewRefNode("verse2", i2))
}

func (this *ActualVisitor) VisitSingleverse(ctx *SingleverseContext) {
	i, _ := strconv.Atoi(ctx.GetText())
	this.stack = append(this.stack, NewRefNode("verse", i))
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
		visitor.Result = new(Reference)
		visitor.VisitReference(ctx.(*ReferenceContext))

		//fmt.Println(visitor.Result)
		if visitor.Result.IsEmpty() {
			this.collector += ctx.GetText()
		} else {
			this.collector += fmt.Sprint(visitor.Result)
		}

	case BibleParserRULE_text:
		this.collector += ctx.GetText()
	}
}

func (this *ReferenceSubstitutionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	ctx.ExitRule(this)
}

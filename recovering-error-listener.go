package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type RecoveringErrorListener struct {
	*antlr.DefaultErrorListener
	InErrorMode int
}

func (d *RecoveringErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//fmt.Println(">> 1>> ", reflect.TypeOf(e).String())
	if e != nil {
		switch e.(type) {
		case *antlr.InputMisMatchException, *antlr.BaseRecognitionException:
			d.InErrorMode = 1
		}
	}
}

func NewRecoveringErrorListener() *RecoveringErrorListener {
	//fmt.Println(">> 2 >>")
	return new(RecoveringErrorListener)
}

func (d *RecoveringErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAmbiguity")
}

func (d *RecoveringErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAttemptingFullContext")
}

func (d *RecoveringErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportContextSensitivity")
}

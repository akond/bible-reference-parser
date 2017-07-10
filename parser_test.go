package parser

import (
	"testing"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func referenceSubstitutionParseWalker(this *TreeShapeListener, ctx antlr.ParserRuleContext) {
	//fmt.Println("****", ctx.GetText());
	//fmt.Println("referenceSubstitutionParseWalker >> ", ctx.GetText())
	switch ctx.GetRuleIndex() {
	case BibleParserRULE_reference:
		this.collector += "*"

	case BibleParserRULE_text, BibleParserRULE_terminator:
		this.collector += ctx.GetText()
	}

	//
	//fmt.Println(ctx.GetRuleIndex(), ctx.GetText());
	//fmt.Printf("-> %T\n", this.errorListener);
	//if this.errorListener.InErrorMode {
	//
	//} else {
	//	switch ctx.GetRuleIndex() {
	//	case BibleParserRULE_reference:
	//		this.collector += "*"
	//
	//	case BibleParserRULE_text:
	//		this.collector += ctx.GetText()
	//	}
	//}
}

func TestParsing(t *testing.T) {
	var references = map[string]string{
		"Ин. 5:1":                                                                             "",
		"Быт. 39:6—12":                                                                        "",
		"1 Ин. 4:18, 19":                                                                      "",
		"Ис. Нав. 3:9—17":                                                                     "",
		"Ин. 9:2, 3":                                                                          "",
		"1 Цар. 15:2, 8, 9":                                                                   "",
		"Быт. 6:4,5,6,10-13":                                                                  "",
		"Быт. 6:4":                                                                            "",
		"1 Цар. 3:4,5,7-11":                                                                   "",
		"Чис. 13:26-34; 14:1-2":                                                               "*; *",
		"Иов 38 ; 42:5, 6;":                                                                   "* ; *;",
		"3 Цар. 17 ; Иов 38 ; 42:5, 6; Лк. 4:24–28 ; Евр. 11:1 ; Откр. 1:17 .":                "* ; * ; *; * ; * ; * .",
		"Прочтите 1Цар.16:7; Матф.7:1 и 1Кор.4:5.":                                            "Прочтите *; * и *.",
		"Иуд. 3,5":                                                                            "",
		"Что общего между 3 Цар. 17:3, 4 и 17:8, 9?":                                          "Что общего между * и *?",
		"4:24-1":                                                                              "4:24-1",
		"басни (см. 2 Петр. 1:16), но содержит ":                                              "басни (см. *), но содержит ",
		"2 Цар. 15–18":                                                                        "*–18",
		"Отрывок, записанный во 2 Цар. 15,18":                                                 "Отрывок, записанный во *,18",
		"Отрывок, записанный во 2 Цар. 15–18":                                                 "Отрывок, записанный во *–18",
		"Иуд. 3,5 не Иуды 4—16 прокол":                                                        "* не * прокол",
		"Почему 1 Цар. 3:4,5,7-11 не соответствует.":                                          "Почему * не соответствует.",
		"Прочитайте Быт. 6:4,5,6,10-13 и ответье на вопрос.":                                  "Прочитайте * и ответье на вопрос.",
		"Быт. 6:4; 7:9":                                                                       "*; *",
		", что слепота была следствием греха человека или его родителей (Ин. 9:2, 3). Должны": ", что слепота была следствием греха человека или его родителей (*). Должны",
		"Ис. Нав. 3:9—17;":                                                                    "*;",
		"1Кор.1:18-2:2":                                                                       "",
		"Деян.23:1-6;25:23-26:29":                                                             "*;*",
	}

	for k, v := range references {
		text := ParseBibleText(k, referenceSubstitutionParseWalker, false)
		expecting := "*"
		if v != "" {
			expecting = v
		}

		if text != expecting {
			t.Fatalf("%q is not recognized as a reference.\nGot\n\t%q\ninstead of\n\t%q", k, text, expecting)
		}
	}
}

func TestTexts(t *testing.T) {
	var texts = []string{

	}
	for _, s := range texts {
		text := ParseBibleText(s, referenceSubstitutionParseWalker, false)
		if text != s {
			t.Fatalf("%q is recognized as a reference. Got\n%q", s, text)
		}
	}
}

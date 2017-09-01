package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"crypto/sha1"
)

type Tags []interface{}

func (this Tags ) String() string  {
	str := ""
	for _, tag := range this {
		str += fmt.Sprintf("%s", tag)
	}
	return str
}

func createTokenStream(text string) *antlr.CommonTokenStream {
	input := antlr.NewInputStream(text)
	lexer := NewBibleLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return stream
}

func createBibleParser(text string) *BibleParser {
	stream := createTokenStream(text)
	bibleParser := NewBibleParser(stream)
	bibleParser.RemoveErrorListeners()

	return bibleParser
}

func SubstituteBibleRefWithStar(text string) string {
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, createBibleParser(text).R())
	return listener.collector
}

func SubstituteBibleRefWithXml(text string, f ReferenceCallback) Tags {
	listener := NewReferenceSubstitutionListener(f)
	antlr.ParseTreeWalkerDefault.Walk(listener, createBibleParser(text).R())
	return listener.collector
}

func biemHash(biem string) string {
	return "bt-" + fmt.Sprintf("%x\n", sha1.Sum([]byte(biem)))[:7]
}

type Ulink struct {
	XMLName struct{} `xml:"ulink"`
	Url string `xml:"url,attr"`
	Id string `xml:"bibletext:id,attr"`
	Data string `xml:",chardata"`
}

func (this *Ulink) String() string  {
	return fmt.Sprintf(`<ulink url="%s" bibletext:id="%s">%s</ulink>`, this.Url, this.Id, this.Data)
}

func InjectBibleReferences(s string) Tags {
	return SubstituteBibleRefWithXml(s, func(result *Reference, s string) *Ulink {
		biem := referenceToBiem(*result)
		return &Ulink{Url:"biem://"+ biem, Id:biemHash(biem), Data:s}
	})
}

func referenceToBiem(r Reference) string {
	str := ""

	for _, part := range r.Parts {
		if part.Chapter == 0 {
			continue
		}

		if 0 < len(str) {
			str += ";"
		}
		str += fmt.Sprintf("%d:", part.Chapter)

		for i, verse := range part.Verses {
			if 0 < verse && i != 0{
				str += ","
			}
			str += fmt.Sprintf("%d", verse)
		}
	}
	return fmt.Sprintf("%d:", r.Book) + str
}

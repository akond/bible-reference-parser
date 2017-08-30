package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"crypto/sha1"
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

func SubstituteBibleRefWithXml(text string, f ReferenceCallback) string {
	listener := NewReferenceSubstitutionListener(f)
	antlr.ParseTreeWalkerDefault.Walk(listener, createBibleParser(text).R())
	return listener.collector
}

func biemHash(biem string) string {
	return "bt-" + fmt.Sprintf("%x\n", sha1.Sum([]byte(biem)))[:7]
}

func InjectBibleReferences(s string) string {
	return SubstituteBibleRefWithXml(s, func(result *Reference, s string) string {
		biem := referenceToBiem(*result)
		hash := biemHash(biem)
		return fmt.Sprintf(`<ulink url="biem://%v" bibletext:id="%s">`, biem, hash) + s + "</ulink>"
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

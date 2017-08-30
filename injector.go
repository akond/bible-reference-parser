package parser

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var (
	bibleBucket = []byte ("bible")
)

type BibleInjector struct {
	in, out   chan string
	collector string
}

func NewBibleInjector(dbLocation string) *BibleInjector {
	db, err := bolt.Open(dbLocation, 0644, nil)
	handleErr(err)

	bibleInjector := new(BibleInjector)
	bibleInjector.in = make(chan string)
	bibleInjector.out = make(chan string)

	go func() {
		db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(bibleBucket)
			for text := range bibleInjector.in {
				bibleInjector.out <- SubstituteBibleRefWithXml(text, func(ref *Reference, s string) string {
					biem := referenceToBiem(*ref)
					hash := biemHash(biem)
					bibleInjector.collectText(bucket, ref, hash)

					return fmt.Sprintf(`<ulink url="biem://%v" bibletext:id="%s">`, biem, hash) + s + "</ulink>"
				})
			}
			return nil
		})
		defer db.Close()
	}()

	return bibleInjector
}

func formatText(text, hash string) string {
	return fmt.Sprintf(`<bibletext:text id="%s">%s</bibletext:text>`, hash, text) + "\n"
}

func formatPart(part Part, text []string) string {
	if part.Text != "" {
		return ""
	}

	verses := ""
	i := 0
	part.Visit(func(verse int) {
		verses += fmt.Sprintf(`<dt>%d</dt><dd>%s</dd>`, verse, text[i]) + "\n"
		i++
	})
	return fmt.Sprintf(`<strong>Глава %d</strong> <dl>%s</dl>`, part.Chapter, verses) + "\n"
}

func (this *BibleInjector) collectText(bucket *bolt.Bucket, ref *Reference, hash string) {
	var texts []string
	var text string

	for _, part := range ref.Parts {
		part.Visit(func(verse int) {
			texts = append(texts, string(bucket.Get(verseKey(ref.Book, part.Chapter, verse))))
		})
		text += formatPart(*part, texts)
	}

	this.collector += formatText(text, hash)
}

func verseKey(book, chapter, verse int) []byte {
	return []byte(fmt.Sprintf("%d-%d-%d", book, chapter, verse))
}

func (this BibleInjector) Inject(s string) string {
	this.in <- s
	return <-this.out
}

func (this BibleInjector) Text() string {
	return `<bibletext:container>` + this.collector + `</bibletext:container>`
}

func (this BibleInjector) Close() {
	close(this.in)
	close(this.out)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("Unable to proceed [%s]", err)
	}
}

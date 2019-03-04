package parser

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strings"
)

var (
	bibleBucket = []byte ("bible")
)

type BibleInjector struct {
	in                  chan string
	out                 chan Tags
	collector           string
	collectorCheckpoint map[string]struct{}
}

func NewBibleInjector(dbLocation string) *BibleInjector {
	db, err := bolt.Open(dbLocation, 0644, nil)
	handleErr(err)

	bibleInjector := new(BibleInjector)
	bibleInjector.in = make(chan string)
	bibleInjector.out = make(chan Tags)
	bibleInjector.collectorCheckpoint = make(map[string]struct{}, 0)

	go func() {
		db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(bibleBucket)
			for text := range bibleInjector.in {
				bibleInjector.out <- SubstituteBibleRefWithXml(text, func(ref *Reference, s string) *Ulink {
					biem := referenceToBiem(*ref)
					hash := biemHash(biem)
					bibleInjector.collectText(bucket, ref, hash)
					return &Ulink{Url: "biem://" + biem, Id: hash, Data: s}
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

func cleanText(string string) string {
	result := string
	result = strings.Replace(result, "<", "[", -1)
	result = strings.Replace(result, ">", "]", -1)
	return result
}

func formatPart(part Part, text []string) string {
	if part.Text != "" {
		return ""
	}

	verses := ""
	i := 0
	part.Visit(func(verse int) {
		verses += fmt.Sprintf(`<dt>%d</dt><dd>%s</dd>`, verse, cleanText(text[i])) + "\n"
		i++
	})
	return fmt.Sprintf(`<strong>Глава %d</strong> <dl>%s</dl>`, part.Chapter, verses) + "\n"
}

type BibleTextCollectorFunc func(this *BibleInjector) (bucket *bolt.Bucket, ref *Reference, hash string);

func (this *BibleInjector) collectText(bucket *bolt.Bucket, ref *Reference, hash string) {
	_, ok := this.collectorCheckpoint[hash]
	if ok {
		return
	}

	var texts []string
	var text string

	for _, part := range ref.Parts {
		texts = make([]string, 0)
		part.Visit(func(verse int) {
			texts = append(texts, string(bucket.Get(verseKey(ref.Book, part.Chapter, verse))))
		})
		text += formatPart(*part, texts)
	}

	this.collector += formatText(text, hash)
	this.collectorCheckpoint[hash] = struct{}{}
}

func verseKey(book, chapter, verse int) []byte {
	return []byte(fmt.Sprintf("%d-%d-%d", book, chapter, verse))
}

func (this BibleInjector) Inject(s string) Tags {
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

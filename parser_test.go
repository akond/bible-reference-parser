package parser

import (
	"testing"
	"regexp"
)

func assertEqual(t *testing.T, expecting, got, message string) {
	if expecting != got {
		t.Fatalf(message+"\nGot\n\t%q\ninstead of\n\t%q", got, expecting)
	}
}

func TestReference_AppendPart(t *testing.T) {
	r := new(Reference)

	r.AppendPartSingle(3, 7)
	assertEqual(t, "{0 [3:7]}", r.String(), "")

	r.AppendPartSingle(3, 8)
	assertEqual(t, "{0 [3:7-8]}", r.String(), "")

	r.AppendPartSingle(3, 9)
	assertEqual(t, "{0 [3:7-9]}", r.String(), "")

	r.AppendPartSingle(3, 22)
	assertEqual(t, "{0 [3:7-9,22]}", r.String(), "")
}

func TestBiemLink(t *testing.T) {
	var texts = map[string]string{
		"Прочитайте Быт. 6:4,5,6,10-13 и ответье на вопрос.":                   `Прочитайте <ulink url="biem://1:6:4-6,10-13">Быт. 6:4,5,6,10-13</ulink> и ответье на вопрос.`,
		"Почему 1 Цар. 3:4,5,7-11 не соответствует.":                           `Почему <ulink url="biem://9:3:4-5,7-11">1 Цар. 3:4,5,7-11</ulink> не соответствует.`,
		"Иуд. 3,5 не Иуды 4—16 прокол":                                         `<ulink url="biem://51:1:3,5">Иуд. 3,5</ulink> не <ulink url="biem://51:1:4-16">Иуды 4—16</ulink> прокол`,
		"Отрывок, записанный во 2 Цар. 15,18":                                  `Отрывок, записанный во <ulink url="biem://10:15:1-37;18:1-33">2 Цар. 15,18</ulink>`,
		"Иов 38 ; 42:5, 6;":                                                    `<ulink url="biem://18:38:1-41;42:5-6">Иов 38 ; 42:5, 6</ulink>;`,
		"Чис. 13:26-34; 14:1-2":                                                `<ulink url="biem://4:13:26-34;14:1-2">Чис. 13:26-34; 14:1-2</ulink>`,
		"басни (см. 2 Петр. 1:16), но содержит ":                               `басни (см. <ulink url="biem://47:1:16">2 Петр. 1:16</ulink>), но содержит `,
		"Что общего между 3 Цар. 17:3, 4 и 17:8, 9?":                           `Что общего между <ulink url="biem://11:17:3-4;17:8-9">3 Цар. 17:3, 4 и 17:8, 9</ulink>?`,
		"3 Цар. 17 ; Иов 38 ; 42:5, 6; Лк. 4:24–28 ; Евр. 11:1 ; Откр. 1:17 .": `<ulink url="biem://11:17:1-24">3 Цар. 17</ulink> ; <ulink url="biem://18:38:1-41;42:5-6">Иов 38 ; 42:5, 6</ulink>; <ulink url="biem://42:4:24-28">Лк. 4:24–28</ulink> ; <ulink url="biem://65:11:1">Евр. 11:1</ulink> ; <ulink url="biem://66:1:17">Откр. 1:17</ulink> .`,
		"Прочтите 1Цар.16:7; Матф.7:1 и 1Кор.4:5.":                             `Прочтите <ulink url="biem://9:16:7">1Цар.16:7</ulink>; <ulink url="biem://40:7:1">Матф.7:1</ulink> и <ulink url="biem://53:4:5">1Кор.4:5</ulink>.`,

		"1 Ин. 4:18, 19":          `<ulink url="biem://48:4:18-19">1 Ин. 4:18, 19</ulink>`,
		"Ис. Нав. 3:9—17;":        `<ulink url="biem://6:3:9-17">Ис. Нав. 3:9—17</ulink>;`,
		"1 Цар. 15:2, 8, 9":       `<ulink url="biem://9:15:2,8-9">1 Цар. 15:2, 8, 9</ulink>`,
		"Быт. 39:6—12,18":         "<ulink url=\"biem://1:39:6-12,18\">Быт. 39:6—12,18</ulink>",
		"Быт. 39:6—12":            `<ulink url="biem://1:39:6-12">Быт. 39:6—12</ulink>`,
		"Быт. 6:4; 7:9":           `<ulink url="biem://1:6:4;7:9">Быт. 6:4; 7:9</ulink>`,
		"1Кор.1:18-2:2":           `<ulink url="biem://53:1:18-31;2:1-2">1Кор.1:18-2:2</ulink>`,
		"Деян.23:1-6;25:23-26:29": `<ulink url="biem://44:23:1-6;25:23-27;26:1-29">Деян.23:1-6;25:23-26:29</ulink>`,
		"Иуд. 7":                  `<ulink url="biem://51:1:7">Иуд. 7</ulink>`,
		"Иуд. 7,8":                `<ulink url="biem://51:1:7-8">Иуд. 7,8</ulink>`,
		"Иуд. 7,9":                `<ulink url="biem://51:1:7,9">Иуд. 7,9</ulink>`,
		"Иуд. 7-9":                `<ulink url="biem://51:1:7-9">Иуд. 7-9</ulink>`,
		"Иуд. 1,7-9":              `<ulink url="biem://51:1:1,7-9">Иуд. 1,7-9</ulink>`,
		"Иуд. 1,2-9":              `<ulink url="biem://51:1:1-9">Иуд. 1,2-9</ulink>`,
		"Иуд. 1-5,6-9":            `<ulink url="biem://51:1:1-9">Иуд. 1-5,6-9</ulink>`,
		"Неем.3:6":                `<ulink url="biem://16:3:6">Неем.3:6</ulink>`,
		"Неем.3:6,8,19-20":        `<ulink url="biem://16:3:6,8,19-20">Неем.3:6,8,19-20</ulink>`,
		"Евр.9:14-18":             `<ulink url="biem://65:9:14-18">Евр.9:14-18</ulink>`,
		"Евр.9:14-10:8":           `<ulink url="biem://65:9:14-28;10:1-8">Евр.9:14-10:8</ulink>`,
		"1 Цар.23:6-1":            `<ulink url="biem://9:23:6">1 Цар.23:6</ulink>-1`,
		"1 Пет.1":                 `<ulink url="biem://46:1:1-25">1 Пет.1</ulink>`,
		"1 Пет.1-2":               `<ulink url="biem://46:1:1-25;2:1-25">1 Пет.1-2</ulink>`,
		"1 Пет.1-3":               `<ulink url="biem://46:1:1-25;2:1-25;3:1-22">1 Пет.1-3</ulink>`,
		"1Кор.1:18-3:2":           `<ulink url="biem://53:1:18-31;2:1-16;3:1-2">1Кор.1:18-3:2</ulink>`,

		"Иуд. cc":        "Иуд. cc",
		"Быт. 39:6—1211": "Быт. 39:6—1211 [[ERROR: 6—1211 has invalid verse number]]",
		"Иуд. 211":       "Иуд. 211 [[ERROR: 211 has invalid verse number]]",

		"Ин. 5:1":                                                                             `<ulink url="biem://43:5:1">Ин. 5:1</ulink>`,
		"Ис. Нав. 3:9—17":                                                                     `<ulink url="biem://6:3:9-17">Ис. Нав. 3:9—17</ulink>`,
		"Ин. 9:2, 3":                                                                          `<ulink url="biem://43:9:2-3">Ин. 9:2, 3</ulink>`,
		"Быт. 6:4,5,6,10-13":                                                                  `<ulink url="biem://1:6:4-6,10-13">Быт. 6:4,5,6,10-13</ulink>`,
		"Быт. 6:4":                                                                            `<ulink url="biem://1:6:4">Быт. 6:4</ulink>`,
		"1 Цар. 3:4,5,7-11":                                                                   `<ulink url="biem://9:3:4-5,7-11">1 Цар. 3:4,5,7-11</ulink>`,
		"Иуд. 3,5":                                                                            `<ulink url="biem://51:1:3,5">Иуд. 3,5</ulink>`,
		"4:24-1":                                                                              "4:24-1",
		"2 Цар. 15–18":                                                                        `<ulink url="biem://10:15:1-37;16:1-23;17:1-29;18:1-33">2 Цар. 15–18</ulink>`,
		"Отрывок, записанный во 2 Цар. 15–18":                                                 `Отрывок, записанный во <ulink url="biem://10:15:1-37;16:1-23;17:1-29;18:1-33">2 Цар. 15–18</ulink>`,
		", что слепота была следствием греха человека или его родителей (Ин. 9:2, 3). Должны": `, что слепота была следствием греха человека или его родителей (<ulink url="biem://43:9:2-3">Ин. 9:2, 3</ulink>). Должны`,
		"Деян.23:6-1":                                                                         `<ulink url="biem://44:23:6">Деян.23:6</ulink>-1`,
		"Деян.23:6-24:1":                                                                      `<ulink url="biem://44:23:6-35;24:1-1">Деян.23:6-24:1</ulink>`,
		"нашуго Авв.2-3 текста":                                                               `нашуго <ulink url="biem://35:2:1-20;3:1-19">Авв.2-3</ulink> текста`,
		"нашуго Авв.2 - 3 текста":                                                             `нашуго <ulink url="biem://35:2:1-20">Авв.2</ulink> - 3 текста`,

		"": ``,
	}

	removeTextId := regexp.MustCompile(`\sbibletext:id.*?\>`)
	for text, expecting := range texts {
		got := SubstituteBibleRefWithXml(text, UlinkFactory()).String()
		got = removeTextId.ReplaceAllString(got, ">")

		assertEqual(t, expecting, got, text)
	}
}

func UlinkFactory() ReferenceCallback {
	return func(ref *Reference, s string) *Ulink {
		biem := referenceToBiem(*ref)
		hash := biemHash(biem)
		return &Ulink{Url: "biem://" + biem, Id: hash, Data: s}
	}
}

func TestHashFunction(t *testing.T) {
	var texts = map[string]string{
		"Быт. 39:6—12": `bt-e90e1cf`,
	}

	for text, expecting := range texts {
		got := SubstituteBibleRefWithXml(text, UlinkFactory())

		if got[0].(*Ulink).Id != expecting {
			t.Fatalf("hashes do not match for %q", text)
		}
	}
}

func TestSomeThing(t *testing.T) {
	//injector := NewBibleInjector("bible.boltdb")
	//fmt.Println(injector.Inject("Быт. 39:6—12,18;40:1. Ин.3:16\n\n"))
	//fmt.Println(injector.Text())
}

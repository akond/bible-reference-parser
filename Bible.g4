grammar Bible;

@lexer::header {
    //import "strconv"
}

@parser::header {
}

@lexer::members {
}

@parser::members {

    func (p *BibleParser) Find(k int) antlr.Token {
    	stream := p.GetTokenStream()
    	i := p.GetCurrentToken().GetTokenIndex() + 1

    	for stream.Get(i).GetTokenType() != k {
			i++
    	}

    	return stream.Get(i)
    }

	func IsValidVersespan(p *BibleParser) bool {
		first, _ := strconv.Atoi(p.GetCurrentToken().GetText())
		last, _ := strconv.Atoi(p.Find(BibleParserNUMBER).GetText())
		return first < last;
	}

	func IsValidChapterSpan(p *BibleParser, localctx antlr.RuleContext) bool {
		first, _ := strconv.Atoi(p.GetCurrentToken().GetText())
		last, _ := strconv.Atoi(p.Find(BibleParserNUMBER).GetText())
		return first < last;
	}
}

// --------------------------------------------------------------------------------------------------------------------
// Правила
// --------------------------------------------------------------------------------------------------------------------
r: ( reference | text)*;

reference: (reference1 | reference2 | reference3);

reference1: book2 whitespace parts;
reference2: book2 whitespace chapter (':' spanlist )? ;
reference3: book1 whitespace spanlist;

parts: part (terminator part)*;
part: barereference1 | barereference2 | barereference3 | barereference4;

barereference1: chapterverse MINUS chapterverse;

barereference2: chapter ':' spanlist ;

barereference3:
	{IsValidChapterSpan(p, localctx)}?
	chapter MINUS chapter;

barereference4: chapter;


chapterverse: chapter ':' verse ;
terminator: WS* (';' | '.' | 'и') WS*;
verselist: versespan | verse;
spanlist: verselist (',' WS* verselist)*;

versespan:
	{IsValidVersespan(p)}?
	verse MINUS verse;

chapter: NUMBER;
verse: NUMBER;
whitespace: WS*;
text: (ANY | MINUS | NUMBER | WS | DOT | ';' | ':' | ',' | 'и')+?;

book1: SINGLECHAPTERBOOK DOT?;
book2: MANYCHAPTERSBOOK DOT?;


// --------------------------------------------------------------------------------------------------------------------
// Лексемы
// --------------------------------------------------------------------------------------------------------------------
MANYCHAPTERSBOOK: OLDTESTAMENT | NEWTESTAMENT;
OLDTESTAMENT: MOSES | HISTORY | TEACHINGS | BIGPROPHETS | SMALLPROPHETS1 | SMALLPROPHETS2;
NEWTESTAMENT: EVANGELY | JOHN | CASUAL | CORYNTH;

MOSES: 'Бытие' | 'Быт' | 'Исх' | 'Левит' | 'Лев' | 'Числ' | 'Чис' | 'Второзаконие' | 'Втор' ;
HISTORY: 'Ис' DOT? WS* 'Нав' | 'Нав' | 'Суд' | 'Руфь' | 'Руф' | KINGS | 'Езд' | 'Неем' | 'Есф';
KINGS: ('1' | '2' | '3' | '4') WS* 'Цар';
TEACHINGS: 'Иова' | 'Иов' | 'Псал' | 'Пс' | 'Притчи' | 'Притч' | 'Прит' | 'Еккл' | 'Екк' | 'Песн';
BIGPROPHETS: 'Исаии' | 'Ис' | ('Пл' | 'Плач') DOT? WS* ('Иеремии' | 'Иерем' | 'Иер') | 'Плач' | 'Иерем' | 'Иер' | 'Иезек' | 'Иез';
SMALLPROPHETS1: 'Дан' | 'Осии' | 'Ос' | 'Иоил' | 'Иоиля' | 'Иоиль' | 'Амоса' | 'Амос' | 'Ам' | 'Ионы' | 'Иона' | 'Ион';
SMALLPROPHETS2: 'Михею' | 'Михея' | 'Мих' | 'Наума' | 'Наум' | 'Авв' | 'Соф' | 'Агг' | 'Зах' | 'Мал';
EVANGELY: 'Матфея' | 'Матф' | 'Мат' | 'Мф' | 'Марка' | 'Марк' | 'Мар' | 'Мк' | 'Луки' | 'Лук' | 'Лк' | 'Иоанна' | 'Иоан' | 'Ин';
JOHN: ('1' | '2' | '3') WS* ('Иоан' | 'Ин');
CASUAL: 'Иак' | 'Римл' | 'Рим' | 'Деяни' [я.] WS* 'апостолов' | 'Деян' | 'Гал' | 'Ефес' | 'Еф' | 'Филип' | 'Филп' | 'Фил' | 'Флп' | 'Колос' | 'Кол' | 'Титу' | 'Тит' | 'Флм' | 'Евреям' | 'Евр' | 'Откровении' | 'Откр' | 'Отк'  ;
CORYNTH: ('1' | '2') WS* ('Пет' | 'Петра' | 'Петр' | 'Кор' | 'Тим' | 'Фес');
SINGLECHAPTERBOOK: 'Иуды' | 'Иуд' | 'Авд' ;

DOT: '.';
NUMBER: ('0'..'9')+;
MINUS: [\p{Pd}];
WS: [\p{Z}];
ANY: .;

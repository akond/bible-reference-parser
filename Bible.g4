grammar Bible;

@lexer::header {
//    import "strconv"
}

@parser::header {
}


@lexer::members {
	var n int
//	func abc (i int) {}
}

@parser::members {
func isOk (a int) bool {
	fmt.Println ("first ->", a, ".")
	return true
}
}

// --------------------------------------------------------------------------------------------------------------------
// Правила
// --------------------------------------------------------------------------------------------------------------------
r: ( reference | text)*;

reference: (reference1 | reference2 | reference3);
//continuation: terminator (chapter ':' spanlist | spanlist ) ;

reference1: book2 whitespace parts;
reference2: book2 whitespace chapter (':' spanlist )? ;
reference3: book1 whitespace spanlist;

parts: part (terminator part)*;
part: barereference1 | barereference2 | barereference3;
barereference1: chapterverse MINUS chapterverse;
barereference2: chapter ':' spanlist ;
barereference3: chapter;

chapterverse: chapter ':' verse ;
terminator: WS* (';' | '.' | 'и') WS*;
verselist: versespan | verse;
spanlist: verselist (',' WS* verselist)*;

versespan:
	A = verse

	MINUS
		{
//			length := len (localctx.GetText())
		}

	verse
	 ;

chapter: NUMBER;
verse: NUMBER;
whitespace: WS*;
text: (ANY | MINUS | NUMBER | WS | DOT | ';' | ':' | ',' | 'и')+?;

book2: (OLDTESTAMENT | NEWTESTAMENT) DOT?;
book1: SINGLECHAPTERBOOK DOT?;


// --------------------------------------------------------------------------------------------------------------------
// Лексемы
// --------------------------------------------------------------------------------------------------------------------
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

grammar Bible;

@parser::header {
//    import "strconv"
}


// --------------------------------------------------------------------------------------------------------------------
// Правила
// --------------------------------------------------------------------------------------------------------------------

r: ( reference | text)*;

reference: (reference3 | reference1 | reference2 | barereference);
continuation: terminator (chapter ':' spanlist | spanlist ) ;

reference1: book1 whitespace chapter (':' spanlist )? ;
reference2: book2 whitespace spanlist;
reference3: book1 whitespace chapterverse MINUS chapterverse ;

barereference: barereference2 | barereference1 ;
barereference1: chapter ':' spanlist ;
barereference2: chapterverse MINUS chapterverse;

chapterverse: chapter ':' verse ;
terminator: WS* (';' | '.') WS*;

spanlist: (versespan | verse) (',' WS* (versespan | verse))*;

versespan:
	NUMBER
		{
			firstVerse, _ := strconv.Atoi(localctx.GetText())
		}

	MINUS
		{
			length := len (localctx.GetText())
		}

	NUMBER
		{
			secondVerse, _ := strconv.Atoi(localctx.GetText()[length:])
			if secondVerse <= firstVerse {
				panic(antlr.NewBaseRecognitionException("Invalid verse span", p, p.GetInputStream(), localctx))
			}
		}
	 ;
chapter: NUMBER;
verse: NUMBER;
whitespace: WS*;
text: (ANY | MINUS | NUMBER | WS | DOT | ';' | ':' | ',')+?;

book1: (OLDTESTAMENT | NEWTESTAMENT) DOT?;
book2: SINGLECHAPTERBOOK DOT?;


// --------------------------------------------------------------------------------------------------------------------
// Лексемы
// --------------------------------------------------------------------------------------------------------------------
MINUS: '-' | '—' | '–';

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
WS: [\p{Z}];
ANY: .;

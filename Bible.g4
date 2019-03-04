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
		defer func() bool {
			recover()
			return false
		}()
		first, _ := strconv.Atoi(p.GetCurrentToken().GetText())
		last, _ := strconv.Atoi(p.Find(BibleParserNUMBER).GetText())
		return first < last;
	}

	func IsValidChapterSpan(p *BibleParser) bool {
		defer func() bool {
			recover()
			return false
		}()
		first, _ := strconv.Atoi(p.GetCurrentToken().GetText())
		last, _ := strconv.Atoi(p.Find(BibleParserNUMBER).GetText())
		return first < last
	}
}

// --------------------------------------------------------------------------------------------------------------------
// Правила
// --------------------------------------------------------------------------------------------------------------------
r: ( reference | text)*;

reference:
	manychaptersbook whitespace parts
	| manychaptersbook whitespace chapter (':' spanlist )?
	| singlechapterbook whitespace ref5;

parts: part (terminator part)*;
part: ref1
	| ref2
	| ref3
    | ref4;

ref1: chapterverse MINUS chapterverse;
ref2: chapter ':' spanlist;
ref3: {IsValidChapterSpan(p)}?
           	chapter MINUS chapter;
ref4: chapter;
ref5: spanlist;

chapterverse: chapter ':' singleverse;
terminator: WS* (';' | '.' | 'и' | ',') WS*;
spanlist: (versespan | singleverse) (',' WS* (versespan | singleverse))*;

versespan:
	{IsValidVersespan(p)}?
	verse MINUS verse;

chapter: NUMBER;
verse: NUMBER;
singleverse: NUMBER;
whitespace: WS*;
text: (ANY | MINUS | NUMBER | WS | DOT | ';' | ':' | ',' | 'и');

singlechapterbook: (book31 | book51 | book64) DOT?;
manychaptersbook: (book1 |  book2 |  book3 |  book4 |  book5 |  book6 |  book7 |  book8 |
                   book9 |  book10 |  book11 |  book12 |  book13 |  book14 |  book15 |
                   book16 |  book17 |  book18 |  book19 |  book20 |  book21 |  book22 |
                   book23 |  book24 |  book25 |  book26 |  book27 |  book28 |  book29 |
                   book30 |  book32 |  book33 |  book34 |  book35 |  book36 |
                   book37 |  book38 |  book39 |  book40 |  book41 |  book42 |  book43 |
                   book44 |  book45 |  book46 |  book47 |  book48 |  book49 |  book50 |
                   book52 |  book53 |  book54 |  book55 |  book56 |  book57 |
                   book58 |  book59 |  book60 |  book61 |  book62 |  book63 |  
                   book65 |  book66) DOT?;

book1: BOOK1;
book2: BOOK2;
book3: BOOK3;
book4: BOOK4;
book5: BOOK5;
book6: BOOK6;
book7: BOOK7;
book8: BOOK8;
book9: BOOK9;
book10: BOOK10;
book11: BOOK11;
book12: BOOK12;
book13: BOOK13;
book14: BOOK14;
book15: BOOK15;
book16: BOOK16;
book17: BOOK17;
book18: BOOK18;
book19: BOOK19;
book20: BOOK20;
book21: BOOK21;
book22: BOOK22;
book23: BOOK23;
book24: BOOK24;
book25: BOOK25;
book26: BOOK26;
book27: BOOK27;
book28: BOOK28;
book29: BOOK29;
book30: BOOK30;
book31: BOOK31;
book32: BOOK32;
book33: BOOK33;
book34: BOOK34;
book35: BOOK35;
book36: BOOK36;
book37: BOOK37;
book38: BOOK38;
book39: BOOK39;
book40: BOOK40;
book41: BOOK41;
book42: BOOK42;
book43: BOOK43;
book44: BOOK44;
book45: BOOK45;
book46: BOOK46;
book47: BOOK47;
book48: BOOK48;
book49: BOOK49;
book50: BOOK50;
book51: BOOK51;
book52: BOOK52;
book53: BOOK53;
book54: BOOK54;
book55: BOOK55;
book56: BOOK56;
book57: BOOK57;
book58: BOOK58;
book59: BOOK59;
book60: BOOK60;
book61: BOOK61;
book62: BOOK62;
book63: BOOK63;
book64: BOOK64;
book65: BOOK65;
book66: BOOK66;

// --------------------------------------------------------------------------------------------------------------------
// Лексемы
// --------------------------------------------------------------------------------------------------------------------
BOOK1: 'Бытие' | 'Быт';
BOOK2: 'Исх';
BOOK3: 'Левит' | 'Лев';
BOOK4: 'Числ' | 'Чис';
BOOK5: 'Втор';
BOOK6: 'Ис' DOT? WS* 'Нав' | 'Нав';
BOOK7: 'Суд';
BOOK8: 'Руфь' | 'Руф';
BOOK9: '1' WS* 'Цар';
BOOK10: '2' WS* 'Цар';
BOOK11: '3' WS* 'Цар';
BOOK12: '4' WS* 'Цар';
BOOK13: '1' WS* 'Пар';
BOOK14: '2' WS* 'Пар';
BOOK15: 'Езд';
BOOK16: 'Неем';
BOOK17: 'Есф';
BOOK18: 'Иова' | 'Иов' ;
BOOK19: 'Псал' | 'Пс';
BOOK20: 'Притчи' | 'Притч' | 'Прит';
BOOK21: 'Еккл' | 'Екк';
BOOK22: 'Пес' DOT? WS* 'П' DOT | 'Песн';
BOOK23: 'Исаии' | 'Ис';
BOOK24: 'Иеремии' | 'Иерем' | 'Иер';
BOOK25: ('Пл' | 'Плач') DOT? WS* ('Иеремии' | 'Иерем' | 'Иер');
BOOK26: 'Иезек' | 'Иез';
BOOK27: 'Дан';
BOOK28: 'Осии' | 'Ос';
BOOK29: 'Иоил' | 'Иоиля' | 'Иоиль';
BOOK30: 'Амоса' | 'Амос' | 'Ам';
BOOK31: 'Авд';																			// 1
BOOK32: 'Ионы' | 'Иона' | 'Ион';
BOOK33: 'Михею' | 'Михея' | 'Мих';
BOOK34: 'Наума' | 'Наум';
BOOK35: 'Авв';
BOOK36: 'Соф';
BOOK37: 'Агг';
BOOK38: 'Зах';
BOOK39: 'Мал';
BOOK40: 'Матфея' | 'Матф' | 'Мат' | 'Мф';
BOOK41: 'Марка' | 'Марк' | 'Мар' | 'Мк';
BOOK42: 'Луки' | 'Лук' | 'Лк';
BOOK43: 'Иоанна' | 'Иоан' | 'Ин';
BOOK44: 'Деяни' [я.] WS* 'апостолов' | 'Деян';
BOOK45: 'Иак';
BOOK46: '1' WS* ('Пет' | 'Петра' | 'Петр');
BOOK47: '2' WS* ('Пет' | 'Петра' | 'Петр');
BOOK48: '1' WS* ('Иоан' | 'Ин');
BOOK49: '2' WS* ('Иоан' | 'Ин');
BOOK50: '3' WS* ('Иоан' | 'Ин');
BOOK51: 'Иуды' | 'Иуд';																	// 1
BOOK52: 'Римл' | 'Рим';
BOOK53: '1' WS* 'Кор';
BOOK54: '2' WS* 'Кор';
BOOK55: 'Гал';
BOOK56: 'Ефес' | 'Еф';
BOOK57: 'Филип' | 'Филп' | 'Фил' | 'Флп';
BOOK58: 'Колос' | 'Кол';
BOOK59: '1' WS* 'Фес';
BOOK60: '2' WS* 'Фес';
BOOK61: '1' WS* 'Тим';
BOOK62: '2' WS* 'Тим';
BOOK63: 'Титу' | 'Тит';
BOOK64: 'Флм';
BOOK65: 'Евреям' | 'Евр';
BOOK66: 'Откровении' | 'Откр' | 'Отк';

DOT: '.';
NUMBER: ('0'..'9')+;
MINUS: [\p{Pd}];
WS: [\p{Z}];
ANY: .;

//line ./timeparse.y:2
package parser

import __yyfmt__ "fmt"

//line ./timeparse.y:2
import (
	tr "github.com/mz-eco/timerange"
	"time"
)

//line ./timeparse.y:11
type yySymType struct {
	yys      int
	stmt     int
	tm       DateTime
	date     Date
	time     Time
	cb       func(int) tr.Interval
	interval tr.Interval
}

const num = 57346
const big = 57347
const interval = 57348
const next = 57349
const stmt = 57350

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"num",
	"big",
	"interval",
	"next",
	"stmt",
	"' '",
	"'T'",
	"','",
	"'['",
	"')'",
	"'+'",
	"'-'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./timeparse.y:65

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 43

var yyAct = [...]int{

	2, 6, 26, 32, 9, 7, 13, 14, 14, 16,
	33, 18, 19, 15, 36, 6, 20, 11, 12, 7,
	8, 31, 30, 3, 5, 27, 29, 6, 35, 28,
	34, 7, 38, 37, 25, 21, 24, 23, 22, 10,
	1, 17, 4,
}
var yyPact = [...]int{

	11, -1000, 9, 23, 8, -1000, -9, -2, -3, 5,
	34, -1000, -1000, 33, 32, 30, -1000, -1000, 21, 21,
	23, -1000, -8, -1000, -13, -5, 24, -1000, -1000, -1000,
	22, 1, 29, 28, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 42, 24, 0, 41, 2, 40, 39,
}
var yyR1 = [...]int{

	0, 7, 7, 6, 6, 6, 5, 5, 5, 4,
	4, 3, 3, 3, 1, 1, 1, 1, 2, 2,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 3, 5, 1, 1, 1, 3,
	3, 3, 1, 1, 1, 3, 5, 1, 3, 5,
}
var yyChk = [...]int{

	-1000, -6, -3, 12, -1, -2, 4, 8, 11, -3,
	-7, 9, 10, 15, 16, 15, -3, -4, 14, 15,
	11, -2, 4, 4, 4, 4, -5, 4, 8, 5,
	-5, -3, 16, 15, 6, 6, 13, 4, 4,
}
var yyDef = [...]int{

	0, -2, 0, 0, 12, 13, 14, 17, 0, 0,
	0, 1, 2, 0, 0, 0, 3, 4, 0, 0,
	0, 11, 0, 15, 18, 0, 0, 6, 7, 8,
	0, 0, 0, 0, 9, 10, 5, 19, 16,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 9, 3, 3, 3, 3, 3, 3, 3,
	3, 13, 3, 14, 11, 15, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 16, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 10, 3, 3, 3, 3, 3,
	3, 12,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./timeparse.y:37
		{
			yylex.(*Lex).tr = makeRange(yyDollar[1].tm, yyDollar[3].tm)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./timeparse.y:38
		{
			yylex.(*Lex).tr = tr.RangeTo(makeTime(yyDollar[1].tm, time.Now()), yyDollar[3].interval)
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./timeparse.y:39
		{
			yylex.(*Lex).tr = makeRange(yyDollar[2].tm, yyDollar[4].tm)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./timeparse.y:43
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./timeparse.y:44
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./timeparse.y:45
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./timeparse.y:47
		{
			yyVAL.interval = yyDollar[3].cb(1 * yyDollar[2].stmt)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./timeparse.y:48
		{
			yyVAL.interval = yyDollar[3].cb(-1 * yyDollar[2].stmt)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./timeparse.y:50
		{
			yyVAL.tm = NewDateTime(yyDollar[1].date, yyDollar[3].time)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./timeparse.y:51
		{
			yyVAL.tm = NewDateTime(yyDollar[1].date, Time{})
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./timeparse.y:52
		{
			yyVAL.tm = NewDateTime(NewDate(-1, 0, 0), yyDollar[1].time)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./timeparse.y:56
		{
			yyVAL.date = NewDate(-1, 0, yyDollar[1].stmt)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./timeparse.y:57
		{
			yyVAL.date = NewDate(-1, yyDollar[1].stmt, yyDollar[3].stmt)
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./timeparse.y:58
		{
			yyVAL.date = NewDate(yyDollar[1].stmt, yyDollar[3].stmt, yyDollar[5].stmt)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./timeparse.y:59
		{
			yyVAL.date = NewDate(yyDollar[1].stmt, 0, 0)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./timeparse.y:62
		{
			yyVAL.time = NewTime(yyDollar[1].stmt, yyDollar[3].stmt, 0)
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./timeparse.y:63
		{
			yyVAL.time = NewTime(yyDollar[1].stmt, yyDollar[3].stmt, yyDollar[5].stmt)
		}
	}
	goto yystack /* stack new state and value */
}

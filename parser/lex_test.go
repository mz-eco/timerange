package parser

import (
	"fmt"
	"strings"
	"testing"
	"text/scanner"
)

func TestLex_Error(t *testing.T) {

	lex := NewLex("[10-12,10-16)")



	yyParse(lex)

	fmt.Println(lex.tr)
}

func TestName(t *testing.T) {

	lex := new(Lex)
	lex.scanner.Init(strings.NewReader("2015-12-22 14:22:22,-12days"))
	lex.scanner.Whitespace = 0
	lex.scanner.Mode = scanner.ScanInts | scanner.ScanIdents

	for tok := lex.scanner.Scan();tok != scanner.EOF;tok = lex.scanner.Scan() {
		fmt.Println(lex.scanner.TokenText())
	}
}
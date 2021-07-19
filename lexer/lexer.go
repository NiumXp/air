package lexer

import (
	"errors"
	"unicode"

	t "github.com/NiumXp/air/lexer/tokens"
)

type Lexer struct {
	Input []rune

	index int
}

func NewLexer(input string) Lexer {
	return Lexer{
		Input: []rune(input),

		index: 0, // zero value explicit
	}
}

func (l *Lexer) atEOF() bool {
	return !(l.index < len(l.Input))
}

func (l *Lexer) nextRune(consume bool) rune {
	value := l.Input[l.index]

	if !l.atEOF() && consume {
		l.index += 1
	}

	return value
}

func (l *Lexer) doubleSymbol(s rune, p rune) t.Token {
	if l.nextRune(false) == p {
		l.nextRune(true)
		return t.Symbol(string(s) + string(p))
	}
	return t.Symbol(string(s))
}

func (l *Lexer) getString(mark rune) (string, error) {
	start := l.index

	// We can't use `last := ''`.
	var last rune

	for {
		if last == '\n' || l.atEOF() {
			return "", errors.New("not finished string")
		}

		actual := l.nextRune(true)

		if (actual == mark) && (last != '\\') {
			// The second expression allows to escape `"`.
			break
		}

		last = actual
	}

	// The magic number is there because `actual == mark` (`"` or `'`).
	return string(l.Input[start : l.index-1]), nil
}

func (l *Lexer) getDigits(initial rune) string {
	start := l.index

	for {
		rune_ := l.nextRune(false)

		if !unicode.IsDigit(rune_) {
			break
		}

		l.nextRune(true)
	}

	return string(initial) + string(l.Input[start:l.index])
}

func (l *Lexer) NextToken() (t.Token, error) {
	if l.atEOF() {
		return t.EOF, nil
	}

	rune_ := l.nextRune(true)
	if unicode.IsSpace(rune_) {
		return l.NextToken() // recursion :tada: :cry:
	}

	switch rune_ {
	case '(', ')', '+', '-', '*', '/', '^', ',', '!', '=':
		return t.Symbol(string(rune_)), nil
	case '>':
		return l.doubleSymbol('>', '='), nil
	case '<':
		return l.doubleSymbol('<', '='), nil
	case '"', '\'':
		str, err := l.getString(rune_)
		return t.Literal(str), err
	}

	if unicode.IsDigit(rune_) {
		digits := l.getDigits(rune_)
		return t.Literal(digits), nil
	}

	return t.Unknown(string(rune_)), nil
}

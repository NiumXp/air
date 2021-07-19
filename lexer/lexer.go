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
	return l.index == len(l.Input)
}

func (l *Lexer) inLastRune() bool {
	return l.index == len(l.Input)-1
}

func (l *Lexer) previousRune() rune {
	return l.Input[l.index-1]
}

func (l *Lexer) actualRune() rune {
	return l.Input[l.index]
}

func (l *Lexer) nextRune() rune {
	return l.Input[l.index+1]
}

func (l *Lexer) walkRune() {
	l.index += 1
}

func (l *Lexer) canUseDoubleSymbol(symbol rune) t.Token {
	actual := string(l.actualRune())

	if l.inLastRune() {
		l.walkRune()
		return t.Symbol(actual)
	}

	if l.nextRune() == symbol {
		l.walkRune()
		l.walkRune()
		return t.Symbol(actual + string(symbol))
	}

	l.walkRune()
	return t.Symbol(actual)
}

func (l *Lexer) getString() (string, error) {
	mark := l.actualRune()
	l.walkRune()

	start := l.index

	for {
		if l.atEOF() || l.actualRune() == '\n' {
			return "", errors.New("not finished string")
		}

		if l.inLastRune() {
			l.walkRune()
			break
		}

		if (l.nextRune() == mark) && (l.previousRune() != '\\') {
			l.walkRune()
			l.walkRune()
			break
		}

		l.walkRune()
	}

	return string(l.Input[start : l.index-1]), nil
}

func (l *Lexer) getDigits() string {
	start := l.index

	for {
		if l.inLastRune() {
			l.walkRune()
		}

		if l.atEOF() {
			break
		}

		if !unicode.IsDigit(l.nextRune()) {
			l.walkRune()
			break
		}

		l.walkRune()
	}

	return string(l.Input[start:l.index])
}

func (l *Lexer) getLetters() string {
	start := l.index

	for {
		if l.inLastRune() {
			l.walkRune()
		}

		if l.atEOF() {
			break
		}

		if !unicode.IsLetter(l.nextRune()) {
			l.walkRune()
			break
		}

		l.walkRune()
	}

	return string(l.Input[start:l.index])
}

func (l *Lexer) removeComment() {
	for {
		if l.atEOF() || l.actualRune() == '\n' {
			break
		}

		l.walkRune()
	}
}

func (l *Lexer) NextToken() (t.Token, error) {
	if l.atEOF() {
		return t.EOF, nil
	}

	rune_ := l.actualRune()

	if unicode.IsSpace(rune_) {
		l.walkRune()
		return l.NextToken() // recursion :tada: :cry:
	}

	switch rune_ {
	case '#':
		l.removeComment()
		return l.NextToken() // recursion again :yayy:
	case '(', ')', '+', '-', '*', '/', '^', ',', '!', '=':
		l.walkRune()
		return t.Symbol(string(rune_)), nil
	case '>', '<':
		return l.canUseDoubleSymbol('='), nil
	case '"', '\'':
		str, err := l.getString()
		return t.Literal(str), err
	}

	if unicode.IsDigit(rune_) {
		return t.Literal(l.getDigits()), nil
	}

	if unicode.IsLetter(rune_) {
		return t.Identifier(l.getLetters()), nil
	}

	l.walkRune()
	return t.Unknown(string(rune_)), nil
}

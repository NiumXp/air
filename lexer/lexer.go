package lexer

import (
	t "github.com/NiumXp/air/lexer/tokens"
)

type Lexer struct {
	Input string

	line   int
	column int
}

func (l *Lexer) NextToken() (t.Token, error) {
	return t.EOF, nil
}

package lexer

const (
	_IDENTIFIER = iota
	_KEYWORD
	_LITERAL
	_SYMBOL
	_EOF
)

var (
	EOF = Token{_EOF, "EOF"}

	KEYWORDS = [...]string{
		"package",
		"use",
		"blow",
		"return",
	}
)

type Token struct {
	Type  int8
	Value string
}

func (t *Token) IsIdentifier() bool {
	return t.Type == _IDENTIFIER
}

func (t *Token) IsKeyword() bool {
	if t.Type == _KEYWORD {
		for _, k := range KEYWORDS {
			if k == t.Value {
				return true
			}
		}
	}
	return false
}

func (t *Token) IsLiteral() bool {
	return t.Type == _LITERAL
}

func (t *Token) IsSymbol() bool {
	return t.Type == _SYMBOL
}

func (t *Token) IsEOF() bool {
	return t.Type == _EOF
}

func Identifier(name string) Token {
	return Token{_IDENTIFIER, name}
}

func Literal(value string) Token {
	return Token{_LITERAL, value}
}

func Symbol(value string) Token {
	return Token{_SYMBOL, value}
}

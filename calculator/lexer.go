package calculator

import (
	"math/big"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	expression []rune // Pointer to the expression
	position   int    // rune position
	length     int    // length of the expression

	buffer []Token // allows for peeking ahead
}

type TokenKind string

const (
	TokenNumber     TokenKind = "Number"
	TokenWhitespace TokenKind = "Whitespace"
	TokenOperator   TokenKind = "Operator"
	TokenText       TokenKind = "Text"
	TokenCurrency   TokenKind = "Currency"
	TokenEnd        TokenKind = "End"
)

type Token struct {
	kind  TokenKind
	value string
	n     big.Float
}

var currencies = []string{"$", "€", "¥", "£", "₹", "₽"}

func NewLexer(input *string) *Lexer {
	runeInput := []rune(*input)

	l := &Lexer{
		expression: runeInput,
		position:   0,
		length:     utf8.RuneCountInString(*input),
	}

	return l
}

func (l *Lexer) Next() Token {
	if len(l.buffer) > 0 {
		token := l.buffer[0]
		l.buffer = l.buffer[1:]
		return token
	}

	for {
		ch := l.nextRune()
		if ch == 0 {
			return Token{kind: TokenEnd}
		}

		if unicode.IsSpace(ch) {
			l.position--
			return l.readWhitespace()
		}

		if unicode.IsDigit(ch) || ch == '.' || ch == ',' {
			l.position--
			return l.readNumber()
		}

		if slices.Contains(currencies, string(ch)) {
			return Token{kind: TokenCurrency, value: string(ch)}
		}

		if unicode.IsLetter(ch) {
			l.position--
			return l.readText()
		}

		switch ch {
		case '+', '-', '/', '^', '(', ')', '%':
			return Token{kind: TokenOperator, value: string(ch)}
		case '*':
			// Peek ahead to see if this is a ** operator (aka ^)
			if l.position < l.length {
				if l.expression[l.position] == '*' {
					l.position++
					return Token{kind: TokenOperator, value: "^"}
				} else {
					return Token{kind: TokenOperator, value: "*"}
				}
			} else {
				return Token{kind: TokenOperator, value: "*"}
			}
		}
	}
}

func (l *Lexer) HasNext() bool {
	return l.Peek().kind != TokenEnd
}

func (l *Lexer) NextNonWs() Token {
	for {
		token := l.Next()
		if token.kind != TokenWhitespace {
			return token
		}
	}
}

func (l *Lexer) Peek() Token {
	if len(l.buffer) == 0 {
		l.buffer = append(l.buffer, l.Next())
	}

	return l.buffer[0]
}

func (l *Lexer) PeekN(n int) Token {
	if len(l.buffer) < n {
		newBuffer := make([]Token, n)
		for i := 0; i < n; i++ {
			newBuffer[i] = l.Next()
		}
		l.buffer = newBuffer
	}

	return l.buffer[n-1]
}

func (l *Lexer) PeekNonWs() Token {
	n := 1
	for {
		token := l.PeekN(n)
		if token.kind != TokenWhitespace {
			return token
		}

		n++
	}
}

func (l *Lexer) TokenizeAll() []Token {
	var tokens []Token
	for {
		token := l.Next()
		if token.kind == TokenEnd {
			break
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func (l *Lexer) nextRune() rune {
	if l.position >= l.length {
		return 0
	}

	ch := l.expression[l.position]
	l.position++
	return ch
}

func (l *Lexer) readNumber() Token {
	start := l.position
	for {
		ch := l.nextRune()
		if ch == 0 {
			break
		}

		if !unicode.IsDigit(ch) && ch != '.' && ch != ',' {
			l.position--
			break
		}
	}

	clean := cleanupSeparators(string(l.expression[start:l.position]))
	valBig := big.NewFloat(0)
	valBig.SetString(clean)

	return Token{
		kind:  TokenNumber,
		value: string(l.expression[start:l.position]),
		n:     *valBig,
	}
}

func (l *Lexer) readWhitespace() Token {
	for {
		ch := l.nextRune()
		if ch == 0 {
			break
		}

		if !unicode.IsSpace(ch) {
			l.position--
			break
		}
	}

	return Token{kind: TokenWhitespace}
}

func cleanupSeparators(input string) string {
	// Determine what was used as a decimal separator, find the last , or . in the input
	var decimalSeparator rune = '.'

	for _, ch := range input {
		if ch == '.' {
			decimalSeparator = '.'
		} else if ch == ',' {
			decimalSeparator = ','
		}
	}

	// Remove all separators that are not the decimal separator
	cleaned := ""
	for _, ch := range input {
		if ch == '.' || ch == ',' {
			if ch == decimalSeparator {
				cleaned += string(ch)
			}
		} else {
			cleaned += string(ch)
		}
	}

	// Replace the decimal separator with a dot
	if decimalSeparator != '.' {
		cleaned = strings.Replace(cleaned, string(decimalSeparator), ".", 1)
	}

	return cleaned
}

func (l *Lexer) readText() Token {
	start := l.position
	for {
		ch := l.nextRune()
		if ch == 0 {
			break
		}

		if !unicode.IsLetter(ch) {
			l.position--
			break
		}
	}

	return Token{
		kind:  TokenText,
		value: string(l.expression[start:l.position]),
	}
}

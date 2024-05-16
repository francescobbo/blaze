package calculator

import (
	"math/big"
	"testing"
)

type testCase struct {
	expr   string
	tokens []Token
}

func TestLexerBasic(t *testing.T) {
	cases := []testCase{
		{
			"2",
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			" (",
			[]Token{
				{TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "(", *big.NewFloat(0)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"hello world",
			[]Token{
				{TokenText, "hello", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "world", *big.NewFloat(0)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"hello 3",
			[]Token{
				{TokenText, "hello", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "3", *big.NewFloat(3)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"pi   is  3.14", // Whitespace is combined
			[]Token{
				{TokenText, "pi", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "is", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "3.14", *big.NewFloat(3.14)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"bad end   ",
			[]Token{
				{TokenText, "bad", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "end", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
	}

	testCases(t, cases)
}

func TestLexerBasicMath(t *testing.T) {
	cases := []testCase{
		{
			"2+3* 4",
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenOperator, "+", *big.NewFloat(0)},
				{TokenNumber, "3", *big.NewFloat(3)}, {TokenOperator, "*", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "4", *big.NewFloat(4)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"(1+2)*3",
			[]Token{
				{TokenOperator, "(", *big.NewFloat(0)}, {TokenNumber, "1", *big.NewFloat(1)},
				{TokenOperator, "+", *big.NewFloat(0)}, {TokenNumber, "2", *big.NewFloat(2)},
				{TokenOperator, ")", *big.NewFloat(0)}, {TokenOperator, "*", *big.NewFloat(0)},
				{TokenNumber, "3", *big.NewFloat(3)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
	}

	testCases(t, cases)
}

func TestLexerSeparators(t *testing.T) {
	cases := []testCase{
		{
			"2,150,000.15",
			[]Token{
				{TokenNumber, "2,150,000.15", *big.NewFloat(2150000.15)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"2.150.000,15",
			[]Token{
				{TokenNumber, "2.150.000,15", *big.NewFloat(2150000.15)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			".15+,15",
			[]Token{
				{TokenNumber, ".15", *big.NewFloat(0.15)}, {TokenOperator, "+", *big.NewFloat(0)},
				{TokenNumber, ",15", *big.NewFloat(0.15)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
	}

	testCases(t, cases)
}

func TestPowers(t *testing.T) {
	cases := []testCase{
		{
			"4 ^6",
			[]Token{
				{TokenNumber, "4", *big.NewFloat(4)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "^", *big.NewFloat(0)}, {TokenNumber, "6", *big.NewFloat(6)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"4** 6",
			[]Token{
				{TokenNumber, "4", *big.NewFloat(4)}, {TokenOperator, "^", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "6", *big.NewFloat(6)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
	}

	testCases(t, cases)
}

func TestPercentMod(t *testing.T) {
	cases := []testCase{
		{
			"15% of 5",
			[]Token{
				{TokenNumber, "15", *big.NewFloat(15)}, {TokenOperator, "%", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "of", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "5", *big.NewFloat(5)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"15 % 5",
			[]Token{
				{TokenNumber, "15", *big.NewFloat(15)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "%", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"2+5% of 10",
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenOperator, "+", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenOperator, "%", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "of", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "10", *big.NewFloat(10)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"2 + 5%",
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenOperator, "%", *big.NewFloat(0)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"2 + 5 %", // Semantic error (% after whitespace is a mod operator, but with a missing operand)
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "%", *big.NewFloat(0)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"2+5 % of 10", // Semantic error (% after whitespace is a mod operator, but with a missing operand)
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenOperator, "+", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "%", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "of", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
	}

	testCases(t, cases)
}

func TestUnits(t *testing.T) {
	cases := []testCase{
		{
			"1 liter + 2 liters",
			[]Token{
				{TokenNumber, "1", *big.NewFloat(1)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "liter", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "liters", *big.NewFloat(0)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"1 liter + 12 oz",
			[]Token{
				{TokenNumber, "1", *big.NewFloat(1)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "liter", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenOperator, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "12", *big.NewFloat(12)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "oz", *big.NewFloat(0)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"10 usd in gbp",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "usd", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "in", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "gbp", *big.NewFloat(0)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"-3c to F",
			[]Token{
				{TokenOperator, "-", *big.NewFloat(0)}, {TokenNumber, "3", *big.NewFloat(3)},
				{TokenText, "c", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "to", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "F", *big.NewFloat(0)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"10m in ft",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenText, "m", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "in", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "ft", *big.NewFloat(0)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"10% of 100$ in eur",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenOperator, "%", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "of", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "100", *big.NewFloat(100)},
				{TokenCurrency, "$", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "in", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "eur", *big.NewFloat(0)}, {TokenEnd, "", *big.NewFloat(0)},
			},
		},
	}

	testCases(t, cases)
}

func TestCurrencySymbols(t *testing.T) {
	cases := []testCase{
		{
			"$10",
			[]Token{
				{TokenCurrency, "$", *big.NewFloat(0)}, {TokenNumber, "10", *big.NewFloat(10)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"10$",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenCurrency, "$", *big.NewFloat(0)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
		{
			"10$ + 20",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenCurrency, "$", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenOperator, "+", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "20", *big.NewFloat(20)},
				{TokenEnd, "", *big.NewFloat(0)},
			},
		},
	}

	testCases(t, cases)
}

func testCases(t *testing.T, cases []testCase) {
	for _, c := range cases {
		t.Run(c.expr, func(t *testing.T) {
			l := NewLexer(&c.expr)

			for _, token := range c.tokens {
				tok := l.Next()
				if tok.kind != token.kind {
					t.Errorf("Expected kind %v, got %v", token.kind, tok.kind)
				}

				if tok.value != token.value {
					t.Errorf("Expected value %v, got %v", token.value, tok.value)
				}

				if tok.n.Cmp(&token.n) != 0 {
					t.Errorf("Expected num %v, got %v", token.n, tok.n)
				}
			}
		})
	}
}

func TestPeeking(t *testing.T) {
	expr := "2+3"

	l := NewLexer(&expr)
	tok := l.Peek()
	if tok.kind != TokenNumber {
		t.Errorf("Expected number, got %v", tok.kind)
	}

	tok = l.Next()
	if tok.kind != TokenNumber {
		t.Errorf("Expected number, got %v", tok.kind)
	}

	tok = l.PeekN(2)
	if tok.kind != TokenNumber {
		t.Errorf("Expected number, got %v", tok.kind)
	}

	tok = l.Next()
	if tok.kind != TokenOperator {
		t.Errorf("Expected operator, got %v", tok.kind)
	}

	tok = l.Next()
	if tok.kind != TokenNumber {
		t.Errorf("Expected number, got %v", tok.kind)
	}
}

func TestWhitespaceHandling(t *testing.T) {
	expr := "2 + 3"

	l := NewLexer(&expr)
	l.Next()

	tok := l.Peek()
	if tok.kind != TokenWhitespace {
		t.Errorf("Expected whitespace, got %v", tok.kind)
	}

	tok = l.NextNonWs()
	if tok.kind != TokenOperator {
		t.Errorf("Expected operator, got %v", tok.kind)
	}

	tok = l.NextNonWs()
	if tok.kind != TokenNumber {
		t.Errorf("Expected number, got %v", tok.kind)
	}

	expr = "    +3"
	l = NewLexer(&expr)
	tok = l.PeekNonWs()
	if tok.kind != TokenOperator {
		t.Errorf("Expected operator, got %v", tok.kind)
	}

	tok = l.Next()
	if tok.kind != TokenWhitespace {
		t.Errorf("Expected whitespace, got %v", tok.kind)
	}
}

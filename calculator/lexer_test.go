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
			},
		},
		{
			" (",
			[]Token{
				{TokenWhitespace, "", *big.NewFloat(0)},
				{TokenLParen, "(", *big.NewFloat(0)},
			},
		},
		{
			"hello world",
			[]Token{
				{TokenText, "hello", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "world", *big.NewFloat(0)},
			},
		},
		{
			"hello 3",
			[]Token{
				{TokenText, "hello", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "3", *big.NewFloat(3)},
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
			},
		},
		{
			"bad end   ",
			[]Token{
				{TokenText, "bad", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "end", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
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
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenPlus, "+", *big.NewFloat(0)},
				{TokenNumber, "3", *big.NewFloat(3)}, {TokenTimes, "*", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "4", *big.NewFloat(4)},
			},
		},
		{
			"(1+2)*3",
			[]Token{
				{TokenLParen, "(", *big.NewFloat(0)}, {TokenNumber, "1", *big.NewFloat(1)},
				{TokenPlus, "+", *big.NewFloat(0)}, {TokenNumber, "2", *big.NewFloat(2)},
				{TokenRParen, ")", *big.NewFloat(0)}, {TokenTimes, "*", *big.NewFloat(0)},
				{TokenNumber, "3", *big.NewFloat(3)},
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
			},
		},
		{
			"2.150.000,15",
			[]Token{
				{TokenNumber, "2.150.000,15", *big.NewFloat(2150000.15)},
			},
		},
		{
			".15+,15",
			[]Token{
				{TokenNumber, ".15", *big.NewFloat(0.15)}, {TokenPlus, "+", *big.NewFloat(0)},
				{TokenNumber, ",15", *big.NewFloat(0.15)},
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
				{TokenPow, "^", *big.NewFloat(0)}, {TokenNumber, "6", *big.NewFloat(6)},
			},
		},
		{
			"4** 6",
			[]Token{
				{TokenNumber, "4", *big.NewFloat(4)}, {TokenPow, "^", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "6", *big.NewFloat(6)},
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
				{TokenNumber, "15", *big.NewFloat(15)}, {TokenPercent, "%", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "of", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "5", *big.NewFloat(5)},
			},
		},
		{
			"15 % 5",
			[]Token{
				{TokenNumber, "15", *big.NewFloat(15)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenPercent, "%", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)},
			},
		},
		{
			"2+5% of 10",
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenPlus, "+", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenPercent, "%", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "of", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "10", *big.NewFloat(10)},
			},
		},
		{
			"2 + 5%",
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenPlus, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenPercent, "%", *big.NewFloat(0)},
			},
		},
		{
			"2 + 5 %", // Semantic error (% after whitespace is a mod operator, but with a missing operand)
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenPlus, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenPercent, "%", *big.NewFloat(0)},
			},
		},
		{
			"2+5 % of 10", // Semantic error (% after whitespace is a mod operator, but with a missing operand)
			[]Token{
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenPlus, "+", *big.NewFloat(0)},
				{TokenNumber, "5", *big.NewFloat(5)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenPercent, "%", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "of", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "10", *big.NewFloat(10)},
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
				{TokenPlus, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "2", *big.NewFloat(2)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "liters", *big.NewFloat(0)},
			},
		},
		{
			"1 liter + 12 oz",
			[]Token{
				{TokenNumber, "1", *big.NewFloat(1)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "liter", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenPlus, "+", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenNumber, "12", *big.NewFloat(12)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "oz", *big.NewFloat(0)},
			},
		},
		{
			"10 usd in gbp",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "usd", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "in", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "gbp", *big.NewFloat(0)},
			},
		},
		{
			"-3c to F",
			[]Token{
				{TokenMinus, "-", *big.NewFloat(0)}, {TokenNumber, "3", *big.NewFloat(3)},
				{TokenText, "c", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "to", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "F", *big.NewFloat(0)},
			},
		},
		{
			"10m in ft",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenText, "m", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "in", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "ft", *big.NewFloat(0)},
			},
		},
		{
			"10% of 100$ in eur",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenPercent, "%", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenText, "of", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "100", *big.NewFloat(100)},
				{TokenCurrency, "$", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "in", *big.NewFloat(0)}, {TokenWhitespace, "", *big.NewFloat(0)},
				{TokenText, "eur", *big.NewFloat(0)},
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
			},
		},
		{
			"10$",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenCurrency, "$", *big.NewFloat(0)},
			},
		},
		{
			"10$ + 20",
			[]Token{
				{TokenNumber, "10", *big.NewFloat(10)}, {TokenCurrency, "$", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenPlus, "+", *big.NewFloat(0)},
				{TokenWhitespace, "", *big.NewFloat(0)}, {TokenNumber, "20", *big.NewFloat(20)},
			},
		},
	}

	testCases(t, cases)
}

func testCases(t *testing.T, cases []testCase) {
	for _, c := range cases {
		t.Run(c.expr, func(t *testing.T) {
			l := NewLexer(&c.expr)
			tokens := l.TokenizeAll()

			if len(tokens) != len(c.tokens) {
				t.Fatalf("Expected %d tokens, got %d", len(c.tokens), len(tokens))
			}

			i := 0
			for _, token := range c.tokens {
				tok := tokens[i]
				if tok.kind != token.kind {
					t.Errorf("Expected kind %v, got %v", token.kind, tok.kind)
				}

				if tok.value != token.value {
					t.Errorf("Expected value %v, got %v", token.value, tok.value)
				}

				if tok.n.Cmp(&token.n) != 0 {
					expf, _ := token.n.Float64()
					gotf, _ := tok.n.Float64()

					t.Errorf("Expected num %f, got %f", expf, gotf)
				}

				i++
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
	if tok.kind != TokenPlus {
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
	if tok.kind != TokenPlus {
		t.Errorf("Expected operator, got %v", tok.kind)
	}

	tok = l.NextNonWs()
	if tok.kind != TokenNumber {
		t.Errorf("Expected number, got %v", tok.kind)
	}

	expr = "    +3"
	l = NewLexer(&expr)
	tok = l.PeekNonWs()
	if tok.kind != TokenPlus {
		t.Errorf("Expected operator, got %v", tok.kind)
	}

	tok = l.Next()
	if tok.kind != TokenWhitespace {
		t.Errorf("Expected whitespace, got %v", tok.kind)
	}
}

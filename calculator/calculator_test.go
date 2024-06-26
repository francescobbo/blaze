package calculator

import (
	"testing"
)

func TestAntlrCalculator(t *testing.T) {
	tests := []struct {
		expr string
		n    float64
		unit string
	}{
		{"2 + 3 * 4", 14, ""},
		{"(1 + 2) * 3", 9, ""},
		{"2 ^ 3", 8, ""},
		{"4 ** 2", 16, ""},
		{"3 + 5 * (2 - 8)", -27, ""},
		{"-3 + 4", 1, ""},
		{"5 * -2", -10, ""},
		{"-(1 + 2) * 3", -9, ""},
		{"10m + 5m", 15, "m"},
		{"10m + 5ft", 11.524, "m"},
		{"5l + 1cup", 5.236588, "l"},
		{"(5l + 1cup) - 1 + 1 tsp", 4.241517, "l"},
		{"20% * 50", 10, ""},
		{"20% of 50", 10, ""},
		{"20 mod 3", 2, ""},
		{"20 % 3", 2, ""},
		{"20 % 3 + 5", 7, ""},
		{"20% + 3", 3.2, ""},
		{"20%", 0.2, ""},
		{"20 %", 0.2, ""},
		{"20 * 50%", 10, ""},
		{"10% of (3m + 2m)", 0.5, "m"},
		{"2 * (3m to ft)", 19.68504, "ft"},
		{"(5ft + 2m) to m", 3.524, "m"},
		{"0.1m + 0.2", 0.3, "m"},
		{"1 / 10", 0.1, ""},
		{"$10 + 20", 30, "USD"},
		{"10 + $20", 30, "USD"},
		{"10$ + 20", 30, "USD"},
		{"1/3 of 100$", 33.333333, "USD"},
		{"€10", 10, "EUR"},
		{"30 + 50¥", 80, "JPY"},
		{"40 - 10£", 30, "GBP"},
		{"sqrt 4", 2, ""},
		{"sqrt(9)", 3, ""},
		{"log(100)", 2, ""},
		{"ln(100)", 4.60517, ""},
		{"log(100 USD)", 2, "USD"},
		{"pi", 3.141593, ""},
		{"pi * 2", 6.283185, ""},
		{"pi * 2 USD", 6.283185, "USD"},
		{"2 * pi", 6.283185, ""},
		{"e * pi", 8.539734, ""},
		{"phi", 1.618034, ""},
		{".15+,15", 0.30, ""},
		{"3sin(pi/3 km/h)", 2.598076, "km/h"},
		// {"80 eur - 20%", 64, "EUR"},
		// {"80 eur - 20% + 10%", 72, "EUR"},
		// {"80 eur + 22%", 97.6, "EUR"},
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			res, err := AntlrEvaluate(test.expr)
			if err != nil {
				t.Fatalf("Evaluate(%q) returned error: %v", test.expr, err)
			}

			if res.Number-test.n > 1e-6 {
				t.Errorf("Evaluate(%q) = %v, want %v", test.expr, res.Number, test.n)
			}

			if res.Unit != test.unit {
				t.Errorf("Evaluate(%q) = %v, want %v", test.expr, res.Unit, test.unit)
			}
		})
	}
}

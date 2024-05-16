package calculator

import (
	"fmt"
	"math"
	"slices"
)

// Value represents a number with an optional unit.
type Value struct {
	Number float64
	Unit   string
}

// Map for currency symbol to ISO code
var currencySymbols = map[string]string{
	"$": "USD",
	"€": "EUR",
	"¥": "JPY",
	"£": "GBP",
}

// Function to evaluate an expression
func Evaluate(expression string) (Value, error) {
	lexer := NewLexer(&expression)
	return parseExpression(lexer)
}

// Function to parse and evaluate an expression
func parseExpression(lexer *Lexer) (Value, error) {
	result, err := parseTerm(lexer)
	if err != nil {
		return Value{}, err
	}

	for {
		token := lexer.PeekNonWs()

		if token.kind == TokenOperator && (token.value == "+" || token.value == "-") {
			lexer.NextNonWs()
			nextTerm, err := parseTerm(lexer)
			if err != nil {
				return Value{}, err
			}
			if nextTerm.Unit != "" && result.Unit == "" {
				result.Unit = nextTerm.Unit
			}
			if result.Unit != "" && nextTerm.Unit != result.Unit {
				nextTerm, err = convert(nextTerm, result.Unit)
				if err != nil {
					return Value{}, err
				}
			}
			if token.value == "+" {
				result.Number += nextTerm.Number
			} else {
				result.Number -= nextTerm.Number
			}
		} else {
			break
		}
	}

	return result, nil
}

// Function to parse and evaluate a term
func parseTerm(lexer *Lexer) (Value, error) {
	result, err := parseFactor(lexer)
	if err != nil {
		return Value{}, err
	}

	for {
		token := lexer.PeekNonWs()

		if token.value == "*" || token.value == "/" || token.value == "of" {
			lexer.NextNonWs()
			nextFactor, err := parseFactor(lexer)
			if err != nil {
				return Value{}, err
			}
			if nextFactor.Unit != "" && result.Unit == "" {
				result.Unit = nextFactor.Unit
			}
			if result.Unit != "" && nextFactor.Unit != result.Unit {
				nextFactor, err = convert(nextFactor, result.Unit)
				if err != nil {
					return Value{}, err
				}
			}
			if token.value == "*" {
				result.Number *= nextFactor.Number
			} else if token.value == "/" {
				result.Number /= nextFactor.Number
			}
		} else if token.value == "mod" {
			lexer.NextNonWs()
			nextFactor, err := parseFactor(lexer)
			if err != nil {
				return Value{}, err
			}
			result.Number = math.Mod(result.Number, nextFactor.Number)
		} else if token.value == "%" {
			lexer.NextNonWs()
			if lexer.PeekNonWs().value == "of" || lexer.PeekNonWs().value == "*" {
				lexer.NextNonWs()
				nextFactor, err := parseFactor(lexer)
				if err != nil {
					return Value{}, err
				}
				result.Number = (result.Number / 100) * nextFactor.Number
				result.Unit = nextFactor.Unit
			} else {
				// If this is an end of expression, then we need to divide by 100
				if lexer.PeekNonWs().kind == TokenEnd {
					result.Number /= 100
				} else {
					// otherwise if this is followed by an operator, this is still a percentage and needs to be divided by 100
					if lexer.PeekNonWs().value == "+" || lexer.PeekNonWs().value == "-" || lexer.PeekNonWs().value == "*" || lexer.PeekNonWs().value == "/" || lexer.PeekNonWs().value == "mod" {
						result.Number /= 100
					} else {
						// this is a modulo operation
						nextFactor, err := parseFactor(lexer)
						if err != nil {
							return Value{}, err
						}
						result.Number = math.Mod(result.Number, nextFactor.Number)
					}
				}
			}
		} else {
			break
		}
	}

	return result, nil
}

// Function to parse and evaluate a factor
func parseFactor(lexer *Lexer) (Value, error) {
	result, err := parsePrimary(lexer)
	if err != nil {
		return Value{}, err
	}

	for {
		token := lexer.PeekNonWs()

		if token.value == "^" {
			lexer.NextNonWs()
			nextPrimary, err := parsePrimary(lexer)
			if err != nil {
				return Value{}, err
			}
			result.Number = math.Pow(result.Number, nextPrimary.Number)
		} else if token.value == "to" {
			lexer.NextNonWs()
			if lexer.HasNext() {
				toUnit := lexer.PeekNonWs()
				lexer.NextNonWs()
				converted, err := convert(result, toUnit.value)
				if err != nil {
					return Value{}, err
				}
				result = converted
			} else {
				return Value{}, fmt.Errorf("missing target unit for conversion")
			}
		} else {
			break
		}
	}

	return result, nil
}

// Function to parse and evaluate a primary (number, negation, or expression in brackets)
func parsePrimary(lexer *Lexer) (Value, error) {
	if !lexer.HasNext() {
		return Value{}, fmt.Errorf("unexpected end of expression")
	}

	token := lexer.NextNonWs()

	functions := []string{"sqrt", "log", "ln", "sin", "cos", "tan", "asin", "acos", "atan"}

	if token.value == "(" {
		result, err := parseExpression(lexer)
		if err != nil {
			return Value{}, err
		}
		if lexer.NextNonWs().value != ")" {
			return Value{}, fmt.Errorf("missing closing bracket")
		}

		return result, nil
	} else if token.value == "-" {
		primary, err := parsePrimary(lexer)
		if err != nil {
			return Value{}, err
		}
		primary.Number = -primary.Number
		return primary, nil
	} else if slices.Contains(functions, token.value) {
		if !lexer.HasNext() {
			return Value{}, fmt.Errorf("expected argument for %v", token)
		}
		arg, err := parsePrimary(lexer)
		if err != nil {
			return Value{}, err
		}
		switch token.value {
		case "sqrt":
			arg.Number = math.Sqrt(arg.Number)
		case "ln":
			arg.Number = math.Log(arg.Number)
		case "log":
			arg.Number = math.Log10(arg.Number)
		case "sin":
			arg.Number = math.Sin(arg.Number)
		case "cos":
			arg.Number = math.Cos(arg.Number)
		case "tan":
			arg.Number = math.Tan(arg.Number)
		case "asin":
			arg.Number = math.Asin(arg.Number)
		case "acos":
			arg.Number = math.Acos(arg.Number)
		case "atan":
			arg.Number = math.Atan(arg.Number)
		}
		return arg, nil
	} else if token.value == "pi" || token.value == "π" || token.value == "e" || token.value == "phi" {
		var number float64
		switch token.value {
		case "pi", "π":
			number = math.Pi
		case "e":
			number = math.E
		case "phi":
			number = math.Phi
		}

		return Value{Number: number}, nil
	} else {
		// Check if token is a currency symbol
		unit := ""
		if token.kind == TokenCurrency {
			unit = currencySymbols[token.value]
			token = lexer.NextNonWs()
		}

		number := token.n

		if unit == "" {
			switch lexer.PeekNonWs().kind {
			case TokenCurrency:
				unit = currencySymbols[lexer.NextNonWs().value]
			case TokenText:
				val := lexer.PeekNonWs().value
				if val != "mod" && val != "of" && val != "to" {
					unit = lexer.NextNonWs().value
				}
			}
		}

		asFloat, _ := number.Float64()

		return Value{Number: asFloat, Unit: unit}, nil
	}
}

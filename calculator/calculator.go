package calculator

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
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
	tokens := tokenize(expression)
	return parseExpression(&tokens)
}

// Function to tokenize the expression
func tokenize(expression string) []string {
	var tokens []string
	var numberBuilder strings.Builder
	var wordBuilder strings.Builder

	var asRune = []rune(expression)

	for i := 0; i < len(asRune); i++ {
		ch := asRune[i]

		if _, ok := currencySymbols[string(ch)]; ok {
			if numberBuilder.Len() > 0 {
				tokens = append(tokens, numberBuilder.String())
				numberBuilder.Reset()
			}
			if wordBuilder.Len() > 0 {
				tokens = append(tokens, wordBuilder.String())
				wordBuilder.Reset()
			}
			tokens = append(tokens, string(ch))
		} else if unicode.IsDigit(ch) || ch == '.' {
			if wordBuilder.Len() > 0 {
				tokens = append(tokens, wordBuilder.String())
				wordBuilder.Reset()
			}
			numberBuilder.WriteRune(ch)
		} else if unicode.IsLetter(ch) {
			if numberBuilder.Len() > 0 {
				ttokens := numberBuilder.String()
				if ttokens[len(ttokens)-1] == ' ' {
					ttokens = ttokens[:len(ttokens)-1]
				}
				tokens = append(tokens, ttokens)
				numberBuilder.Reset()
			}
			wordBuilder.WriteRune(ch)
			if wordBuilder.String() == "mod" || wordBuilder.String() == "of" || wordBuilder.String() == "to" {
				tokens = append(tokens, wordBuilder.String())
				wordBuilder.Reset()
			}
		} else if ch == ' ' {
			if numberBuilder.Len() > 0 {
				tokens = append(tokens, numberBuilder.String())
				numberBuilder.Reset()
			}
			if wordBuilder.Len() > 0 {
				tokens = append(tokens, wordBuilder.String())
				wordBuilder.Reset()
			}
		} else {
			if numberBuilder.Len() > 0 {
				tokens = append(tokens, numberBuilder.String())
				numberBuilder.Reset()
			}
			if wordBuilder.Len() > 0 {
				tokens = append(tokens, wordBuilder.String())
				wordBuilder.Reset()
			}

			if ch == '*' && i+1 < len(asRune) && expression[i+1] == '*' {
				tokens = append(tokens, "**")
				i++
			} else {
				tokens = append(tokens, string(ch))
			}
		}
	}
	if numberBuilder.Len() > 0 {
		tokens = append(tokens, numberBuilder.String())
	}
	if wordBuilder.Len() > 0 {
		tokens = append(tokens, wordBuilder.String())
	}

	return tokens
}

// Function to parse and evaluate an expression
func parseExpression(tokens *[]string) (Value, error) {
	result, err := parseTerm(tokens)
	if err != nil {
		return Value{}, err
	}

	for len(*tokens) > 0 {
		token := (*tokens)[0]

		if token == "+" || token == "-" {
			*tokens = (*tokens)[1:]
			nextTerm, err := parseTerm(tokens)
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
			if token == "+" {
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
func parseTerm(tokens *[]string) (Value, error) {
	result, err := parseFactor(tokens)
	if err != nil {
		return Value{}, err
	}

	for len(*tokens) > 0 {
		token := (*tokens)[0]

		if token == "*" || token == "/" || token == "of" {
			*tokens = (*tokens)[1:]
			nextFactor, err := parseFactor(tokens)
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
			if token == "*" {
				result.Number *= nextFactor.Number
			} else if token == "/" {
				result.Number /= nextFactor.Number
			}
		} else if token == "mod" {
			*tokens = (*tokens)[1:]
			nextFactor, err := parseFactor(tokens)
			if err != nil {
				return Value{}, err
			}
			result.Number = math.Mod(result.Number, nextFactor.Number)
		} else if token == "%" {
			*tokens = (*tokens)[1:]
			if len(*tokens) > 0 && ((*tokens)[0] == "of" || (*tokens)[0] == "*") {
				*tokens = (*tokens)[1:]
				nextFactor, err := parseFactor(tokens)
				if err != nil {
					return Value{}, err
				}
				result.Number = (result.Number / 100) * nextFactor.Number
				result.Unit = nextFactor.Unit
			} else {
				// If this is an end of expression, then we need to divide by 100
				if len(*tokens) == 0 {
					result.Number /= 100
				} else {
					// otherwise if this is followed by an operator, this is still a percentage and needs to be divided by 100
					if (*tokens)[0] == "+" || (*tokens)[0] == "-" || (*tokens)[0] == "*" || (*tokens)[0] == "/" || (*tokens)[0] == "mod" {
						result.Number /= 100
					} else {
						// this is a modulo operation
						nextFactor, err := parseFactor(tokens)
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
func parseFactor(tokens *[]string) (Value, error) {
	result, err := parsePrimary(tokens)
	if err != nil {
		return Value{}, err
	}

	for len(*tokens) > 0 {
		token := (*tokens)[0]

		if token == "^" || token == "**" {
			*tokens = (*tokens)[1:]
			nextPrimary, err := parsePrimary(tokens)
			if err != nil {
				return Value{}, err
			}
			result.Number = math.Pow(result.Number, nextPrimary.Number)
		} else if token == "to" {
			*tokens = (*tokens)[1:]
			if len(*tokens) > 0 {
				toUnit := (*tokens)[0]
				*tokens = (*tokens)[1:]
				converted, err := convert(result, toUnit)
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
func parsePrimary(tokens *[]string) (Value, error) {
	if len(*tokens) == 0 {
		return Value{}, fmt.Errorf("unexpected end of expression")
	}

	token := (*tokens)[0]
	*tokens = (*tokens)[1:]

	functions := []string{"sqrt", "log", "ln", "sin", "cos", "tan", "asin", "acos", "atan"}

	if token == "(" {
		result, err := parseExpression(tokens)
		if err != nil {
			return Value{}, err
		}
		if len(*tokens) == 0 || (*tokens)[0] != ")" {
			return Value{}, fmt.Errorf("missing closing bracket")
		}
		*tokens = (*tokens)[1:]
		return result, nil
	} else if token == "-" {
		primary, err := parsePrimary(tokens)
		if err != nil {
			return Value{}, err
		}
		primary.Number = -primary.Number
		return primary, nil
	} else if slices.Contains(functions, token) {
		if len(*tokens) == 0 {
			return Value{}, fmt.Errorf("expected argument for %s", token)
		}
		arg, err := parsePrimary(tokens)
		if err != nil {
			return Value{}, err
		}
		switch token {
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
	} else if token == "pi" || token == "π" || token == "e" || token == "phi" {
		var number float64
		switch token {
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
		if isoCode, ok := currencySymbols[token]; ok {
			unit = isoCode
			token = (*tokens)[0]
			*tokens = (*tokens)[1:]
		}

		number, err := strconv.ParseFloat(token, 64)
		if err != nil {
			return Value{}, err
		}

		if unit == "" {
			if len(*tokens) > 0 && (unicode.IsLetter(rune((*tokens)[0][0])) || (*tokens)[0][0] == '$') {
				if isoCode, ok := currencySymbols[(*tokens)[0]]; ok {
					unit = isoCode
					*tokens = (*tokens)[1:]
				} else if (*tokens)[0] != "mod" && (*tokens)[0] != "of" && (*tokens)[0] != "to" {
					unit = (*tokens)[0]
					*tokens = (*tokens)[1:]
				}
			}
		}

		return Value{Number: number, Unit: unit}, nil
	}
}

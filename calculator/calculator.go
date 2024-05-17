package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/francescobbo/blaze/calculator/parser"
)

// Value represents a number with an optional unit.
type Value struct {
	Number float64
	Unit   string
}

type CalculatorVisitor struct {
	antlr.BaseParseTreeVisitor

	err error
}

func (v *CalculatorVisitor) VisitExpression(tree parser.IExpressionContext) Value {
	if tree.EOF() != nil {
		v.err = errors.New("empty expression")
		return Value{}
	}

	if tree.GetUnary() != nil {
		factor := v.VisitFactor(tree.GetUnary())
		if tree.GetUop() != nil && tree.GetUop().GetText() == "-" {
			return Value{-factor.Number, factor.Unit}
		}

		return factor
	}

	if tree.Unit() != nil {
		// This is a unit conversion
		lhs := v.VisitExpression(tree.GetLhs())
		unit := Value{1, tree.Unit().GetText()}

		v.fixUnits(&unit, &lhs) // First argument is the target unit

		return lhs
	}

	if tree.GetOp() == nil {
		v.err = errors.New("missing operator")
		return Value{}
	}

	if tree.GetOp().GetText() == "%" {
		lhs := v.VisitExpression(tree.GetLhs())
		rhs := v.VisitFactor(tree.GetRhss())

		v.fixUnits(&lhs, &rhs)

		return Value{math.Mod(lhs.Number, rhs.Number), rhs.Unit}
	} else {
		lhs := v.VisitExpression(tree.GetLhs())
		rhs := v.VisitExpression(tree.GetRhs())

		v.fixUnits(&lhs, &rhs)

		switch tree.GetOp().GetText() {
		case "+":
			return Value{lhs.Number + rhs.Number, lhs.Unit}
		case "-":
			return Value{lhs.Number - rhs.Number, lhs.Unit}
		case "*", "of":
			return Value{lhs.Number * rhs.Number, lhs.Unit}
		case "/":
			return Value{lhs.Number / rhs.Number, lhs.Unit}
		case "^", "**":
			return Value{math.Pow(lhs.Number, rhs.Number), lhs.Unit}
		case "mod":
			return Value{math.Mod(lhs.Number, rhs.Number), rhs.Unit}
		default:
			v.err = errors.New("unknown operator: " + tree.GetOp().GetText())
			return Value{}
		}
	}
}

func (v *CalculatorVisitor) VisitFactor(tree parser.IFactorContext) Value {
	value := v.VisitPrimary(tree.Primary())
	if tree.Unit() != nil {
		if value.Unit == "" {
			value.Unit = tree.Unit().GetText()
		} else {
			v.fixUnits(&Value{1, tree.Unit().GetText()}, &value)
		}
	} else if tree.CURRENCY_SYMBOL() != nil {
		value.Unit = symbolToIso(tree.CURRENCY_SYMBOL().GetText())
	}

	return value
}

func (v *CalculatorVisitor) VisitPrimary(tree parser.IPrimaryContext) Value {
	if tree == nil {
		v.err = errors.New("empty primary")
		return Value{}
	}

	if tree.NUMBER() != nil {
		factor := v.VisitNumber(tree.NUMBER())
		if tree.Constant() != nil {
			switch tree.Constant().CONSTANT().GetText() {
			case "pi", "π":
				return Value{factor * math.Pi, ""}
			case "e":
				return Value{factor * math.E, ""}
			case "phi":
				return Value{factor * math.Phi, ""}
			}
		} else if tree.Fn() != nil {
			value := v.VisitFunc(tree.Fn())

			lhs := Value{factor, ""}

			v.fixUnits(&lhs, &value)
			return Value{lhs.Number * value.Number, value.Unit}
		}

		return Value{factor, ""}
	} else if tree.GetSub() != nil {
		return v.VisitExpression(tree.GetSub())
	} else if tree.Percentage() != nil {
		value := v.VisitNumber(tree.Percentage().NUMBER())
		return Value{value / 100, ""}
	} else if tree.Fn() != nil {
		return v.VisitFunc(tree.Fn())
	} else if tree.Constant() != nil {
		name := tree.Constant().CONSTANT().GetText()
		switch name {
		case "pi", "π":
			return Value{math.Pi, ""}
		case "e":
			return Value{math.E, ""}
		case "phi":
			return Value{math.Phi, ""}
		}
	}

	v.err = errors.New("unknown primary")
	return Value{}
}

func (v *CalculatorVisitor) VisitNumber(tree antlr.TerminalNode) float64 {
	str := cleanupSeparators(tree.GetText())

	// Parse the number
	number, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}

	return number
}

func (v *CalculatorVisitor) VisitFunc(tree parser.IFnContext) Value {
	name := tree.FUNC_NAME().GetText()
	argument := v.VisitFactor(tree.Factor())
	switch name {
	case "sqrt":
		return Value{math.Sqrt(argument.Number), argument.Unit}
	case "log":
		return Value{math.Log10(argument.Number), argument.Unit}
	case "ln":
		return Value{math.Log(argument.Number), argument.Unit}
	case "sin":
		return Value{math.Sin(argument.Number), argument.Unit}
	case "cos":
		return Value{math.Cos(argument.Number), argument.Unit}
	case "tan":
		return Value{math.Tan(argument.Number), argument.Unit}
	case "asin":
		return Value{math.Asin(argument.Number), argument.Unit}
	case "acos":
		return Value{math.Acos(argument.Number), argument.Unit}
	case "atan":
		return Value{math.Atan(argument.Number), argument.Unit}
	}

	v.err = errors.New("unknown function: " + name)
	return Value{}
}

func AntlrEvaluate(expression string) (Value, error) {
	is := antlr.NewInputStream(expression)
	lexer := parser.NewCalculatorLexer(is)
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalculatorParser(stream)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	tree := p.Root()

	visitor := &CalculatorVisitor{}
	result := visitor.VisitExpression(tree.Expression())

	return result, nil
}

func (v *CalculatorVisitor) fixUnits(lhs, rhs *Value) {
	if lhs.Unit == rhs.Unit {
		return
	}

	if lhs.Unit == "" {
		lhs.Unit = rhs.Unit
		return
	}

	if rhs.Unit == "" {
		rhs.Unit = lhs.Unit
		return
	}

	// Try to keep the left hand side unit
	val, err := convert(*rhs, lhs.Unit)
	if err != nil {
		v.err = err
		return
	}

	rhs.Number = val.Number
	rhs.Unit = lhs.Unit
}

func symbolToIso(symbol string) string {
	switch symbol {
	case "$":
		return "USD"
	case "€":
		return "EUR"
	case "¥":
		return "JPY"
	case "£":
		return "GBP"
	case "₽":
		return "RUB"
	case "₹":
		return "INR"
	}

	return ""
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

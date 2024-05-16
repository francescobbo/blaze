package calculator

import (
	"math"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/francescobbo/blaze/calculator/parser"
)

type CalculatorVisitor struct {
	antlr.BaseParseTreeVisitor
}

func (v *CalculatorVisitor) VisitExpression(tree parser.IExpressionContext) float64 {
	if tree.EOF() != nil {
		return 0
	}

	if tree.GetUnary() != nil {
		factor := v.VisitFactor(tree.GetUnary())
		if tree.GetUop() != nil && tree.GetUop().GetText() == "-" {
			return -factor
		}

		return factor
	}

	// Disambiguate the modulo % from the percentage
	if tree.GetPct() != nil {
		lhs := v.VisitNumber(tree.GetPct().NUMBER()) / 100
		rhs := v.VisitExpression(tree.GetRhs())

		return lhs * rhs
	} else if tree.GetOp().GetText() == "%" {
		lhs := v.VisitExpression(tree.GetLhs())
		rhs := v.VisitFactor(tree.GetRhss())

		return math.Mod(lhs, rhs)
	} else {
		lhs := v.VisitExpression(tree.GetLhs())
		rhs := v.VisitExpression(tree.GetRhs())

		switch tree.GetOp().GetText() {
		case "+":
			return lhs + rhs
		case "-":
			return lhs - rhs
		case "*":
			return lhs * rhs
		case "/":
			return lhs / rhs
		case "^", "**":
			return math.Pow(lhs, rhs)
		case "mod":
			return math.Mod(lhs, rhs)
		default:
			// log.Println("Unknown operator:", tree.GetOp().GetText())
		}
	}

	return 0
}

func (v *CalculatorVisitor) VisitFactor(tree parser.IFactorContext) float64 {
	return v.VisitPrimary(tree.Primary())
}

func (v *CalculatorVisitor) VisitPrimary(tree parser.IPrimaryContext) float64 {
	if tree.NUMBER() != nil {
		return v.VisitNumber(tree.NUMBER())
	} else if tree.GetSub() != nil {
		return v.VisitExpression(tree.GetSub())
	} else if tree.Percentage() != nil {
		value := v.VisitNumber(tree.Percentage().NUMBER())
		return value / 100
	} else if tree.Currency() != nil {
		return v.VisitNumber(tree.Currency().NUMBER())
	} else if tree.Fn() != nil {
		name := tree.Fn().FUNC_NAME().GetText()
		argument := v.VisitFactor(tree.Fn().Factor())
		switch name {
		case "sqrt":
			return math.Sqrt(argument)
		case "log":
			return math.Log10(argument)
		case "ln":
			return math.Log(argument)
		case "sin":
			return math.Sin(argument)
		case "cos":
			return math.Cos(argument)
		case "tan":
			return math.Tan(argument)
		case "asin":
			return math.Asin(argument)
		case "acos":
			return math.Acos(argument)
		case "atan":
			return math.Atan(argument)
		}
	} else if tree.Constant() != nil {
		name := tree.Constant().CONSTANT().GetText()
		switch name {
		case "pi", "π":
			return math.Pi
		case "e":
			return math.E
		case "phi":
			return math.Phi
		}
	}
	return 0
}

func (v *CalculatorVisitor) VisitNumber(tree antlr.TerminalNode) float64 {
	// log.Println("VisitNumber", tree.GetText())

	str := cleanupSeparators(tree.GetText())

	// Parse the number
	number, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}

	return number
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

	// Dump the tree

	return Value{result, ""}, nil
}

// func cleanupSeparators(input string) string {
// 	// Determine what was used as a decimal separator, find the last , or . in the input
// 	var decimalSeparator rune = '.'

// 	for _, ch := range input {
// 		if ch == '.' {
// 			decimalSeparator = '.'
// 		} else if ch == ',' {
// 			decimalSeparator = ','
// 		}
// 	}

// 	// Remove all separators that are not the decimal separator
// 	cleaned := ""
// 	for _, ch := range input {
// 		if ch == '.' || ch == ',' {
// 			if ch == decimalSeparator {
// 				cleaned += string(ch)
// 			}
// 		} else {
// 			cleaned += string(ch)
// 		}
// 	}

// 	// Replace the decimal separator with a dot
// 	if decimalSeparator != '.' {
// 		cleaned = strings.Replace(cleaned, string(decimalSeparator), ".", 1)
// 	}

// 	return cleaned
// }

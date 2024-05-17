// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calculator
import "github.com/antlr4-go/antlr/v4"

// CalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type CalculatorListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterFn is called when entering the fn production.
	EnterFn(c *FnContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterPercentage is called when entering the percentage production.
	EnterPercentage(c *PercentageContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitFn is called when exiting the fn production.
	ExitFn(c *FnContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitPercentage is called when exiting the percentage production.
	ExitPercentage(c *PercentageContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)
}

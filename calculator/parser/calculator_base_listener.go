// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calculator
import "github.com/antlr4-go/antlr/v4"

// BaseCalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type BaseCalculatorListener struct{}

var _ CalculatorListener = &BaseCalculatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalculatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalculatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalculatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalculatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCalculatorListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCalculatorListener) ExitRoot(ctx *RootContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCalculatorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCalculatorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseCalculatorListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseCalculatorListener) ExitFactor(ctx *FactorContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseCalculatorListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseCalculatorListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterFn is called when production fn is entered.
func (s *BaseCalculatorListener) EnterFn(ctx *FnContext) {}

// ExitFn is called when production fn is exited.
func (s *BaseCalculatorListener) ExitFn(ctx *FnContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCalculatorListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCalculatorListener) ExitConstant(ctx *ConstantContext) {}

// EnterPercentage is called when production percentage is entered.
func (s *BaseCalculatorListener) EnterPercentage(ctx *PercentageContext) {}

// ExitPercentage is called when production percentage is exited.
func (s *BaseCalculatorListener) ExitPercentage(ctx *PercentageContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseCalculatorListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseCalculatorListener) ExitUnit(ctx *UnitContext) {}

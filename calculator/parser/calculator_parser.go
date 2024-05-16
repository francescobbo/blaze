// Code generated from calculator/parser/Calculator.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Calculator

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 109,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 3, 3, 3, 5,
	3, 36, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 50, 10, 3, 12, 3, 14, 3, 53, 11, 3, 3, 4, 3, 4, 5,
	4, 57, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 62, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 72, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 91, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 101, 10, 10, 3, 10, 7, 10, 104, 10, 10, 12, 10, 14, 10,
	107, 11, 10, 3, 10, 2, 4, 4, 18, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2,
	5, 3, 2, 7, 8, 3, 2, 11, 12, 4, 2, 9, 10, 13, 13, 2, 119, 2, 20, 3, 2,
	2, 2, 4, 35, 3, 2, 2, 2, 6, 54, 3, 2, 2, 2, 8, 71, 3, 2, 2, 2, 10, 73,
	3, 2, 2, 2, 12, 76, 3, 2, 2, 2, 14, 78, 3, 2, 2, 2, 16, 81, 3, 2, 2, 2,
	18, 90, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 3, 3, 2, 2, 2, 22, 23, 8, 3,
	1, 2, 23, 24, 5, 14, 8, 2, 24, 25, 7, 3, 2, 2, 25, 26, 5, 4, 3, 5, 26,
	36, 3, 2, 2, 2, 27, 29, 9, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2,
	2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30,
	3, 2, 2, 2, 33, 36, 5, 6, 4, 2, 34, 36, 7, 2, 2, 3, 35, 22, 3, 2, 2, 2,
	35, 30, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 51, 3, 2, 2, 2, 37, 38, 12,
	9, 2, 2, 38, 39, 9, 3, 2, 2, 39, 50, 5, 4, 3, 10, 40, 41, 12, 8, 2, 2,
	41, 42, 9, 4, 2, 2, 42, 50, 5, 4, 3, 9, 43, 44, 12, 6, 2, 2, 44, 45, 9,
	2, 2, 2, 45, 50, 5, 4, 3, 7, 46, 47, 12, 7, 2, 2, 47, 48, 7, 16, 2, 2,
	48, 50, 5, 6, 4, 2, 49, 37, 3, 2, 2, 2, 49, 40, 3, 2, 2, 2, 49, 43, 3,
	2, 2, 2, 49, 46, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 5, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 56, 5, 8, 5,
	2, 55, 57, 5, 18, 10, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 7,
	3, 2, 2, 2, 58, 61, 7, 17, 2, 2, 59, 62, 5, 12, 7, 2, 60, 62, 5, 10, 6,
	2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 72,
	3, 2, 2, 2, 63, 72, 5, 10, 6, 2, 64, 72, 5, 12, 7, 2, 65, 72, 5, 14, 8,
	2, 66, 72, 5, 16, 9, 2, 67, 68, 7, 14, 2, 2, 68, 69, 5, 4, 3, 2, 69, 70,
	7, 15, 2, 2, 70, 72, 3, 2, 2, 2, 71, 58, 3, 2, 2, 2, 71, 63, 3, 2, 2, 2,
	71, 64, 3, 2, 2, 2, 71, 65, 3, 2, 2, 2, 71, 66, 3, 2, 2, 2, 71, 67, 3,
	2, 2, 2, 72, 9, 3, 2, 2, 2, 73, 74, 7, 4, 2, 2, 74, 75, 5, 6, 4, 2, 75,
	11, 3, 2, 2, 2, 76, 77, 7, 5, 2, 2, 77, 13, 3, 2, 2, 2, 78, 79, 7, 17,
	2, 2, 79, 80, 7, 16, 2, 2, 80, 15, 3, 2, 2, 2, 81, 82, 7, 6, 2, 2, 82,
	83, 7, 17, 2, 2, 83, 17, 3, 2, 2, 2, 84, 85, 8, 10, 1, 2, 85, 86, 7, 14,
	2, 2, 86, 87, 5, 18, 10, 2, 87, 88, 7, 15, 2, 2, 88, 91, 3, 2, 2, 2, 89,
	91, 7, 18, 2, 2, 90, 84, 3, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 105, 3, 2,
	2, 2, 92, 93, 12, 6, 2, 2, 93, 94, 7, 9, 2, 2, 94, 104, 5, 18, 10, 7, 95,
	96, 12, 5, 2, 2, 96, 97, 7, 10, 2, 2, 97, 104, 5, 18, 10, 6, 98, 100, 12,
	7, 2, 2, 99, 101, 7, 11, 2, 2, 100, 99, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 102, 3, 2, 2, 2, 102, 104, 7, 17, 2, 2, 103, 92, 3, 2, 2, 2, 103,
	95, 3, 2, 2, 2, 103, 98, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3,
	2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 19, 3, 2, 2, 2, 107, 105, 3, 2, 2,
	2, 13, 30, 35, 49, 51, 56, 61, 71, 90, 100, 103, 105,
}
var literalNames = []string{
	"", "'of'", "", "", "", "'+'", "'-'", "'*'", "'/'", "'^'", "'**'", "'mod'",
	"'('", "')'", "'%'",
}
var symbolicNames = []string{
	"", "", "FUNC_NAME", "CONSTANT", "CURRENCY_SYMBOL", "ADD", "SUB", "MUL",
	"DIV", "POW", "PPOW", "MOD", "LPAREN", "RPAREN", "PCT", "NUMBER", "UNIT_NAME",
	"WS",
}

var ruleNames = []string{
	"root", "expression", "factor", "primary", "fn", "constant", "percentage",
	"currency", "unit",
}

type CalculatorParser struct {
	*antlr.BaseParser
}

// NewCalculatorParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CalculatorParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalculatorParser(input antlr.TokenStream) *CalculatorParser {
	this := new(CalculatorParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calculator.g4"

	return this
}

// CalculatorParser tokens.
const (
	CalculatorParserEOF             = antlr.TokenEOF
	CalculatorParserT__0            = 1
	CalculatorParserFUNC_NAME       = 2
	CalculatorParserCONSTANT        = 3
	CalculatorParserCURRENCY_SYMBOL = 4
	CalculatorParserADD             = 5
	CalculatorParserSUB             = 6
	CalculatorParserMUL             = 7
	CalculatorParserDIV             = 8
	CalculatorParserPOW             = 9
	CalculatorParserPPOW            = 10
	CalculatorParserMOD             = 11
	CalculatorParserLPAREN          = 12
	CalculatorParserRPAREN          = 13
	CalculatorParserPCT             = 14
	CalculatorParserNUMBER          = 15
	CalculatorParserUNIT_NAME       = 16
	CalculatorParserWS              = 17
)

// CalculatorParser rules.
const (
	CalculatorParserRULE_root       = 0
	CalculatorParserRULE_expression = 1
	CalculatorParserRULE_factor     = 2
	CalculatorParserRULE_primary    = 3
	CalculatorParserRULE_fn         = 4
	CalculatorParserRULE_constant   = 5
	CalculatorParserRULE_percentage = 6
	CalculatorParserRULE_currency   = 7
	CalculatorParserRULE_unit       = 8
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *CalculatorParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalculatorParserRULE_root)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUop returns the uop token.
	GetUop() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetUop sets the uop token.
	SetUop(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IExpressionContext

	// GetPct returns the pct rule contexts.
	GetPct() IPercentageContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// GetUnary returns the unary rule contexts.
	GetUnary() IFactorContext

	// GetRhss returns the rhss rule contexts.
	GetRhss() IFactorContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IExpressionContext)

	// SetPct sets the pct rule contexts.
	SetPct(IPercentageContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// SetUnary sets the unary rule contexts.
	SetUnary(IFactorContext)

	// SetRhss sets the rhss rule contexts.
	SetRhss(IFactorContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IExpressionContext
	pct    IPercentageContext
	rhs    IExpressionContext
	uop    antlr.Token
	unary  IFactorContext
	op     antlr.Token
	rhss   IFactorContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetUop() antlr.Token { return s.uop }

func (s *ExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionContext) SetUop(v antlr.Token) { s.uop = v }

func (s *ExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionContext) GetLhs() IExpressionContext { return s.lhs }

func (s *ExpressionContext) GetPct() IPercentageContext { return s.pct }

func (s *ExpressionContext) GetRhs() IExpressionContext { return s.rhs }

func (s *ExpressionContext) GetUnary() IFactorContext { return s.unary }

func (s *ExpressionContext) GetRhss() IFactorContext { return s.rhss }

func (s *ExpressionContext) SetLhs(v IExpressionContext) { s.lhs = v }

func (s *ExpressionContext) SetPct(v IPercentageContext) { s.pct = v }

func (s *ExpressionContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *ExpressionContext) SetUnary(v IFactorContext) { s.unary = v }

func (s *ExpressionContext) SetRhss(v IFactorContext) { s.rhss = v }

func (s *ExpressionContext) Percentage() IPercentageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPercentageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPercentageContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *ExpressionContext) AllADD() []antlr.TerminalNode {
	return s.GetTokens(CalculatorParserADD)
}

func (s *ExpressionContext) ADD(i int) antlr.TerminalNode {
	return s.GetToken(CalculatorParserADD, i)
}

func (s *ExpressionContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(CalculatorParserSUB)
}

func (s *ExpressionContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(CalculatorParserSUB, i)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalculatorParserEOF, 0)
}

func (s *ExpressionContext) POW() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPOW, 0)
}

func (s *ExpressionContext) PPOW() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPPOW, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalculatorParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMOD, 0)
}

func (s *ExpressionContext) PCT() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPCT, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CalculatorParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CalculatorParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalculatorParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(21)

			var _x = p.Percentage()

			localctx.(*ExpressionContext).pct = _x
		}
		{
			p.SetState(22)
			p.Match(CalculatorParserT__0)
		}
		{
			p.SetState(23)

			var _x = p.expression(3)

			localctx.(*ExpressionContext).rhs = _x
		}

	case 2:
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CalculatorParserADD || _la == CalculatorParserSUB {
			{
				p.SetState(25)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExpressionContext).uop = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == CalculatorParserADD || _la == CalculatorParserSUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExpressionContext).uop = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(31)

			var _x = p.Factor()

			localctx.(*ExpressionContext).unary = _x
		}

	case 3:
		{
			p.SetState(32)
			p.Match(CalculatorParserEOF)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(36)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalculatorParserPOW || _la == CalculatorParserPPOW) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(37)

					var _x = p.expression(8)

					localctx.(*ExpressionContext).rhs = _x
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(39)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CalculatorParserMUL)|(1<<CalculatorParserDIV)|(1<<CalculatorParserMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)

					var _x = p.expression(7)

					localctx.(*ExpressionContext).rhs = _x
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(42)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalculatorParserADD || _la == CalculatorParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(43)

					var _x = p.expression(5)

					localctx.(*ExpressionContext).rhs = _x
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(45)

					var _m = p.Match(CalculatorParserPCT)

					localctx.(*ExpressionContext).op = _m
				}
				{
					p.SetState(46)

					var _x = p.Factor()

					localctx.(*ExpressionContext).rhss = _x
				}

			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *FactorContext) Unit() IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *CalculatorParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalculatorParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Primary()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(53)
			p.unit(0)
		}

	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSub returns the sub rule contexts.
	GetSub() IExpressionContext

	// SetSub sets the sub rule contexts.
	SetSub(IExpressionContext)

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sub    IExpressionContext
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) GetSub() IExpressionContext { return s.sub }

func (s *PrimaryContext) SetSub(v IExpressionContext) { s.sub = v }

func (s *PrimaryContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNUMBER, 0)
}

func (s *PrimaryContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PrimaryContext) Fn() IFnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnContext)
}

func (s *PrimaryContext) Percentage() IPercentageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPercentageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPercentageContext)
}

func (s *PrimaryContext) Currency() ICurrencyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrencyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurrencyContext)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLPAREN, 0)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserRPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *CalculatorParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalculatorParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(CalculatorParserNUMBER)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(57)
				p.Constant()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(58)
				p.Fn()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Fn()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Constant()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.Percentage()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(64)
			p.Currency()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(65)
			p.Match(CalculatorParserLPAREN)
		}
		{
			p.SetState(66)

			var _x = p.expression(0)

			localctx.(*PrimaryContext).sub = _x
		}
		{
			p.SetState(67)
			p.Match(CalculatorParserRPAREN)
		}

	}

	return localctx
}

// IFnContext is an interface to support dynamic dispatch.
type IFnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnContext differentiates from other interfaces.
	IsFnContext()
}

type FnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnContext() *FnContext {
	var p = new(FnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_fn
	return p
}

func (*FnContext) IsFnContext() {}

func NewFnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnContext {
	var p = new(FnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_fn

	return p
}

func (s *FnContext) GetParser() antlr.Parser { return s.parser }

func (s *FnContext) FUNC_NAME() antlr.TerminalNode {
	return s.GetToken(CalculatorParserFUNC_NAME, 0)
}

func (s *FnContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterFn(s)
	}
}

func (s *FnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitFn(s)
	}
}

func (p *CalculatorParser) Fn() (localctx IFnContext) {
	localctx = NewFnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalculatorParserRULE_fn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(CalculatorParserFUNC_NAME)
	}
	{
		p.SetState(72)
		p.Factor()
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(CalculatorParserCONSTANT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *CalculatorParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalculatorParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(CalculatorParserCONSTANT)
	}

	return localctx
}

// IPercentageContext is an interface to support dynamic dispatch.
type IPercentageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPercentageContext differentiates from other interfaces.
	IsPercentageContext()
}

type PercentageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPercentageContext() *PercentageContext {
	var p = new(PercentageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_percentage
	return p
}

func (*PercentageContext) IsPercentageContext() {}

func NewPercentageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PercentageContext {
	var p = new(PercentageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_percentage

	return p
}

func (s *PercentageContext) GetParser() antlr.Parser { return s.parser }

func (s *PercentageContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNUMBER, 0)
}

func (s *PercentageContext) PCT() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPCT, 0)
}

func (s *PercentageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PercentageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterPercentage(s)
	}
}

func (s *PercentageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitPercentage(s)
	}
}

func (p *CalculatorParser) Percentage() (localctx IPercentageContext) {
	localctx = NewPercentageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalculatorParserRULE_percentage)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(CalculatorParserNUMBER)
	}
	{
		p.SetState(77)
		p.Match(CalculatorParserPCT)
	}

	return localctx
}

// ICurrencyContext is an interface to support dynamic dispatch.
type ICurrencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrencyContext differentiates from other interfaces.
	IsCurrencyContext()
}

type CurrencyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrencyContext() *CurrencyContext {
	var p = new(CurrencyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_currency
	return p
}

func (*CurrencyContext) IsCurrencyContext() {}

func NewCurrencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrencyContext {
	var p = new(CurrencyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_currency

	return p
}

func (s *CurrencyContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrencyContext) CURRENCY_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserCURRENCY_SYMBOL, 0)
}

func (s *CurrencyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNUMBER, 0)
}

func (s *CurrencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrencyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterCurrency(s)
	}
}

func (s *CurrencyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitCurrency(s)
	}
}

func (p *CalculatorParser) Currency() (localctx ICurrencyContext) {
	localctx = NewCurrencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalculatorParserRULE_currency)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(CalculatorParserCURRENCY_SYMBOL)
	}
	{
		p.SetState(80)
		p.Match(CalculatorParserNUMBER)
	}

	return localctx
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLPAREN, 0)
}

func (s *UnitContext) AllUnit() []IUnitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnitContext)(nil)).Elem())
	var tst = make([]IUnitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnitContext)
		}
	}

	return tst
}

func (s *UnitContext) Unit(i int) IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *UnitContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserRPAREN, 0)
}

func (s *UnitContext) UNIT_NAME() antlr.TerminalNode {
	return s.GetToken(CalculatorParserUNIT_NAME, 0)
}

func (s *UnitContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMUL, 0)
}

func (s *UnitContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalculatorParserDIV, 0)
}

func (s *UnitContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNUMBER, 0)
}

func (s *UnitContext) POW() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPOW, 0)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculatorListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *CalculatorParser) Unit() (localctx IUnitContext) {
	return p.unit(0)
}

func (p *CalculatorParser) unit(_p int) (localctx IUnitContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewUnitContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IUnitContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, CalculatorParserRULE_unit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserLPAREN:
		{
			p.SetState(83)
			p.Match(CalculatorParserLPAREN)
		}
		{
			p.SetState(84)
			p.unit(0)
		}
		{
			p.SetState(85)
			p.Match(CalculatorParserRPAREN)
		}

	case CalculatorParserUNIT_NAME:
		{
			p.SetState(87)
			p.Match(CalculatorParserUNIT_NAME)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_unit)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(91)
					p.Match(CalculatorParserMUL)
				}
				{
					p.SetState(92)
					p.unit(5)
				}

			case 2:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_unit)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(94)
					p.Match(CalculatorParserDIV)
				}
				{
					p.SetState(95)
					p.unit(4)
				}

			case 3:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_unit)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(98)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == CalculatorParserPOW {
					{
						p.SetState(97)
						p.Match(CalculatorParserPOW)
					}

				}
				{
					p.SetState(100)
					p.Match(CalculatorParserNUMBER)
				}

			}

		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

func (p *CalculatorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 8:
		var t *UnitContext = nil
		if localctx != nil {
			t = localctx.(*UnitContext)
		}
		return p.Unit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalculatorParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CalculatorParser) Unit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

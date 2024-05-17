// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calculator
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CalculatorParser struct {
	*antlr.BaseParser
}

var CalculatorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calculatorParserInit() {
	staticData := &CalculatorParserStaticData
	staticData.LiteralNames = []string{
		"", "'to'", "'in'", "'of'", "", "", "", "'+'", "'-'", "'*'", "'/'",
		"'^'", "'**'", "'mod'", "'('", "')'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "FUNC_NAME", "CONSTANT", "CURRENCY_SYMBOL", "ADD", "SUB",
		"MUL", "DIV", "POW", "PPOW", "MOD", "LPAREN", "RPAREN", "PCT", "NUMBER",
		"UNIT_NAME", "WS",
	}
	staticData.RuleNames = []string{
		"root", "expression", "factor", "primary", "fn", "constant", "percentage",
		"unit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 107, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 21,
		8, 1, 10, 1, 12, 1, 24, 9, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 45, 8, 1, 10, 1, 12, 1, 48, 9, 1, 1, 2, 1, 2, 3, 2, 52, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 59, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 64,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 73, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 89, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		99, 8, 7, 1, 7, 5, 7, 102, 8, 7, 10, 7, 12, 7, 105, 9, 7, 1, 7, 0, 2, 2,
		14, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 4, 1, 0, 7, 8, 1, 0, 11, 12, 3, 0,
		3, 3, 9, 10, 13, 13, 1, 0, 1, 2, 119, 0, 16, 1, 0, 0, 0, 2, 27, 1, 0, 0,
		0, 4, 58, 1, 0, 0, 0, 6, 72, 1, 0, 0, 0, 8, 74, 1, 0, 0, 0, 10, 77, 1,
		0, 0, 0, 12, 79, 1, 0, 0, 0, 14, 88, 1, 0, 0, 0, 16, 17, 3, 2, 1, 0, 17,
		1, 1, 0, 0, 0, 18, 22, 6, 1, -1, 0, 19, 21, 7, 0, 0, 0, 20, 19, 1, 0, 0,
		0, 21, 24, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 25,
		1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 25, 28, 3, 4, 2, 0, 26, 28, 5, 0, 0, 1,
		27, 18, 1, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 46, 1, 0, 0, 0, 29, 30, 10,
		6, 0, 0, 30, 31, 7, 1, 0, 0, 31, 45, 3, 2, 1, 7, 32, 33, 10, 5, 0, 0, 33,
		34, 7, 2, 0, 0, 34, 45, 3, 2, 1, 6, 35, 36, 10, 3, 0, 0, 36, 37, 7, 0,
		0, 0, 37, 45, 3, 2, 1, 4, 38, 39, 10, 7, 0, 0, 39, 40, 7, 3, 0, 0, 40,
		45, 3, 14, 7, 0, 41, 42, 10, 4, 0, 0, 42, 43, 5, 16, 0, 0, 43, 45, 3, 4,
		2, 0, 44, 29, 1, 0, 0, 0, 44, 32, 1, 0, 0, 0, 44, 35, 1, 0, 0, 0, 44, 38,
		1, 0, 0, 0, 44, 41, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 3, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 51, 3, 6,
		3, 0, 50, 52, 3, 14, 7, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		59, 1, 0, 0, 0, 53, 54, 5, 6, 0, 0, 54, 59, 3, 6, 3, 0, 55, 56, 3, 6, 3,
		0, 56, 57, 5, 6, 0, 0, 57, 59, 1, 0, 0, 0, 58, 49, 1, 0, 0, 0, 58, 53,
		1, 0, 0, 0, 58, 55, 1, 0, 0, 0, 59, 5, 1, 0, 0, 0, 60, 63, 5, 17, 0, 0,
		61, 64, 3, 10, 5, 0, 62, 64, 3, 8, 4, 0, 63, 61, 1, 0, 0, 0, 63, 62, 1,
		0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 73, 1, 0, 0, 0, 65, 73, 3, 8, 4, 0, 66,
		73, 3, 10, 5, 0, 67, 73, 3, 12, 6, 0, 68, 69, 5, 14, 0, 0, 69, 70, 3, 2,
		1, 0, 70, 71, 5, 15, 0, 0, 71, 73, 1, 0, 0, 0, 72, 60, 1, 0, 0, 0, 72,
		65, 1, 0, 0, 0, 72, 66, 1, 0, 0, 0, 72, 67, 1, 0, 0, 0, 72, 68, 1, 0, 0,
		0, 73, 7, 1, 0, 0, 0, 74, 75, 5, 4, 0, 0, 75, 76, 3, 4, 2, 0, 76, 9, 1,
		0, 0, 0, 77, 78, 5, 5, 0, 0, 78, 11, 1, 0, 0, 0, 79, 80, 5, 17, 0, 0, 80,
		81, 5, 16, 0, 0, 81, 13, 1, 0, 0, 0, 82, 83, 6, 7, -1, 0, 83, 84, 5, 14,
		0, 0, 84, 85, 3, 14, 7, 0, 85, 86, 5, 15, 0, 0, 86, 89, 1, 0, 0, 0, 87,
		89, 5, 18, 0, 0, 88, 82, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 103, 1, 0,
		0, 0, 90, 91, 10, 4, 0, 0, 91, 92, 5, 9, 0, 0, 92, 102, 3, 14, 7, 5, 93,
		94, 10, 3, 0, 0, 94, 95, 5, 10, 0, 0, 95, 102, 3, 14, 7, 4, 96, 98, 10,
		5, 0, 0, 97, 99, 5, 11, 0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99,
		100, 1, 0, 0, 0, 100, 102, 5, 17, 0, 0, 101, 90, 1, 0, 0, 0, 101, 93, 1,
		0, 0, 0, 101, 96, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0,
		0, 103, 104, 1, 0, 0, 0, 104, 15, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 12,
		22, 27, 44, 46, 51, 58, 63, 72, 88, 98, 101, 103,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalculatorParserInit initializes any static state used to implement CalculatorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalculatorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalculatorParserInit() {
	staticData := &CalculatorParserStaticData
	staticData.once.Do(calculatorParserInit)
}

// NewCalculatorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalculatorParser(input antlr.TokenStream) *CalculatorParser {
	CalculatorParserInit()
	this := new(CalculatorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalculatorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Calculator.g4"

	return this
}

// CalculatorParser tokens.
const (
	CalculatorParserEOF             = antlr.TokenEOF
	CalculatorParserT__0            = 1
	CalculatorParserT__1            = 2
	CalculatorParserT__2            = 3
	CalculatorParserFUNC_NAME       = 4
	CalculatorParserCONSTANT        = 5
	CalculatorParserCURRENCY_SYMBOL = 6
	CalculatorParserADD             = 7
	CalculatorParserSUB             = 8
	CalculatorParserMUL             = 9
	CalculatorParserDIV             = 10
	CalculatorParserPOW             = 11
	CalculatorParserPPOW            = 12
	CalculatorParserMOD             = 13
	CalculatorParserLPAREN          = 14
	CalculatorParserRPAREN          = 15
	CalculatorParserPCT             = 16
	CalculatorParserNUMBER          = 17
	CalculatorParserUNIT_NAME       = 18
	CalculatorParserWS              = 19
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
	CalculatorParserRULE_unit       = 7
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// GetUnary returns the unary rule contexts.
	GetUnary() IFactorContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// GetRhss returns the rhss rule contexts.
	GetRhss() IFactorContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IExpressionContext)

	// SetUnary sets the unary rule contexts.
	SetUnary(IFactorContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// SetRhss sets the rhss rule contexts.
	SetRhss(IFactorContext)

	// Getter signatures
	Factor() IFactorContext
	AllADD() []antlr.TerminalNode
	ADD(i int) antlr.TerminalNode
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	EOF() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	POW() antlr.TerminalNode
	PPOW() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	Unit() IUnitContext
	PCT() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IExpressionContext
	uop    antlr.Token
	unary  IFactorContext
	op     antlr.Token
	rhs    IExpressionContext
	rhss   IFactorContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *ExpressionContext) GetUnary() IFactorContext { return s.unary }

func (s *ExpressionContext) GetRhs() IExpressionContext { return s.rhs }

func (s *ExpressionContext) GetRhss() IFactorContext { return s.rhss }

func (s *ExpressionContext) SetLhs(v IExpressionContext) { s.lhs = v }

func (s *ExpressionContext) SetUnary(v IFactorContext) { s.unary = v }

func (s *ExpressionContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *ExpressionContext) SetRhss(v IFactorContext) { s.rhss = v }

func (s *ExpressionContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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

func (s *ExpressionContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserFUNC_NAME, CalculatorParserCONSTANT, CalculatorParserCURRENCY_SYMBOL, CalculatorParserADD, CalculatorParserSUB, CalculatorParserLPAREN, CalculatorParserNUMBER:
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CalculatorParserADD || _la == CalculatorParserSUB {
			{
				p.SetState(19)

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

			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(25)

			var _x = p.Factor()

			localctx.(*ExpressionContext).unary = _x
		}

	case CalculatorParserEOF:
		{
			p.SetState(26)
			p.Match(CalculatorParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(30)

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
					p.SetState(31)

					var _x = p.expression(7)

					localctx.(*ExpressionContext).rhs = _x
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(33)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9736) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(34)

					var _x = p.expression(6)

					localctx.(*ExpressionContext).rhs = _x
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(36)

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
					p.SetState(37)

					var _x = p.expression(4)

					localctx.(*ExpressionContext).rhs = _x
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(39)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CalculatorParserT__0 || _la == CalculatorParserT__1) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)
					p.unit(0)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).lhs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _m = p.Match(CalculatorParserPCT)

					localctx.(*ExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(43)

					var _x = p.Factor()

					localctx.(*ExpressionContext).rhss = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	Unit() IUnitContext
	CURRENCY_SYMBOL() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *FactorContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *FactorContext) CURRENCY_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserCURRENCY_SYMBOL, 0)
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
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Primary()
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(50)
				p.unit(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(CalculatorParserCURRENCY_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Primary()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Primary()
		}
		{
			p.SetState(56)
			p.Match(CalculatorParserCURRENCY_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	NUMBER() antlr.TerminalNode
	Constant() IConstantContext
	Fn() IFnContext
	Percentage() IPercentageContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	sub    IExpressionContext
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PrimaryContext) Fn() IFnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnContext)
}

func (s *PrimaryContext) Percentage() IPercentageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPercentageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPercentageContext)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLPAREN, 0)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserRPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(CalculatorParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(61)
				p.Constant()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(62)
				p.Fn()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Fn()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Constant()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.Percentage()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.Match(CalculatorParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)

			var _x = p.expression(0)

			localctx.(*PrimaryContext).sub = _x
		}
		{
			p.SetState(70)
			p.Match(CalculatorParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnContext is an interface to support dynamic dispatch.
type IFnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC_NAME() antlr.TerminalNode
	Factor() IFactorContext

	// IsFnContext differentiates from other interfaces.
	IsFnContext()
}

type FnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnContext() *FnContext {
	var p = new(FnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_fn
	return p
}

func InitEmptyFnContext(p *FnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_fn
}

func (*FnContext) IsFnContext() {}

func NewFnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnContext {
	var p = new(FnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_fn

	return p
}

func (s *FnContext) GetParser() antlr.Parser { return s.parser }

func (s *FnContext) FUNC_NAME() antlr.TerminalNode {
	return s.GetToken(CalculatorParserFUNC_NAME, 0)
}

func (s *FnContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(CalculatorParserFUNC_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Factor()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONSTANT() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(CalculatorParserCONSTANT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPercentageContext is an interface to support dynamic dispatch.
type IPercentageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	PCT() antlr.TerminalNode

	// IsPercentageContext differentiates from other interfaces.
	IsPercentageContext()
}

type PercentageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPercentageContext() *PercentageContext {
	var p = new(PercentageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_percentage
	return p
}

func InitEmptyPercentageContext(p *PercentageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_percentage
}

func (*PercentageContext) IsPercentageContext() {}

func NewPercentageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PercentageContext {
	var p = new(PercentageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(CalculatorParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(CalculatorParserPCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	RPAREN() antlr.TerminalNode
	UNIT_NAME() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	POW() antlr.TerminalNode

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalculatorParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLPAREN, 0)
}

func (s *UnitContext) AllUnit() []IUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitContext); ok {
			len++
		}
	}

	tst := make([]IUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitContext); ok {
			tst[i] = t.(IUnitContext)
			i++
		}
	}

	return tst
}

func (s *UnitContext) Unit(i int) IUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, CalculatorParserRULE_unit, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserLPAREN:
		{
			p.SetState(83)
			p.Match(CalculatorParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.unit(0)
		}
		{
			p.SetState(85)
			p.Match(CalculatorParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalculatorParserUNIT_NAME:
		{
			p.SetState(87)
			p.Match(CalculatorParserUNIT_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_unit)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(CalculatorParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
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
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(CalculatorParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
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
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				p.SetState(98)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == CalculatorParserPOW {
					{
						p.SetState(97)
						p.Match(CalculatorParserPOW)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(100)
					p.Match(CalculatorParserNUMBER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CalculatorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 7:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CalculatorParser) Unit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

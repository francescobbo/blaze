// Code generated from Calculator.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalculatorLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calculatorlexerLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "FUNC_NAME", "CONSTANT", "CURRENCY_SYMBOL",
		"ADD", "SUB", "MUL", "DIV", "POW", "PPOW", "MOD", "LPAREN", "RPAREN",
		"PCT", "NUMBER", "DIGIT", "UNIT_NAME", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 124, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 69, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		78, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 4, 16, 107, 8, 16, 11, 16, 12,
		16, 108, 1, 17, 1, 17, 1, 18, 4, 18, 114, 8, 18, 11, 18, 12, 18, 115, 1,
		19, 4, 19, 119, 8, 19, 11, 19, 12, 19, 120, 1, 19, 1, 19, 0, 0, 20, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 0, 37, 18, 39, 19,
		1, 0, 5, 2, 0, 101, 101, 966, 966, 6, 0, 36, 36, 163, 163, 165, 165, 8364,
		8364, 8377, 8377, 8381, 8381, 2, 0, 44, 44, 46, 46, 2, 0, 65, 90, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 134, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 44, 1, 0, 0, 0, 5, 47, 1,
		0, 0, 0, 7, 68, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 79, 1, 0, 0, 0, 13,
		81, 1, 0, 0, 0, 15, 83, 1, 0, 0, 0, 17, 85, 1, 0, 0, 0, 19, 87, 1, 0, 0,
		0, 21, 89, 1, 0, 0, 0, 23, 91, 1, 0, 0, 0, 25, 94, 1, 0, 0, 0, 27, 98,
		1, 0, 0, 0, 29, 100, 1, 0, 0, 0, 31, 102, 1, 0, 0, 0, 33, 106, 1, 0, 0,
		0, 35, 110, 1, 0, 0, 0, 37, 113, 1, 0, 0, 0, 39, 118, 1, 0, 0, 0, 41, 42,
		5, 116, 0, 0, 42, 43, 5, 111, 0, 0, 43, 2, 1, 0, 0, 0, 44, 45, 5, 105,
		0, 0, 45, 46, 5, 110, 0, 0, 46, 4, 1, 0, 0, 0, 47, 48, 5, 111, 0, 0, 48,
		49, 5, 102, 0, 0, 49, 6, 1, 0, 0, 0, 50, 51, 5, 115, 0, 0, 51, 52, 5, 113,
		0, 0, 52, 53, 5, 114, 0, 0, 53, 69, 5, 116, 0, 0, 54, 55, 5, 108, 0, 0,
		55, 56, 5, 111, 0, 0, 56, 69, 5, 103, 0, 0, 57, 58, 5, 108, 0, 0, 58, 69,
		5, 110, 0, 0, 59, 60, 5, 115, 0, 0, 60, 61, 5, 105, 0, 0, 61, 69, 5, 110,
		0, 0, 62, 63, 5, 99, 0, 0, 63, 64, 5, 111, 0, 0, 64, 69, 5, 115, 0, 0,
		65, 66, 5, 116, 0, 0, 66, 67, 5, 97, 0, 0, 67, 69, 5, 110, 0, 0, 68, 50,
		1, 0, 0, 0, 68, 54, 1, 0, 0, 0, 68, 57, 1, 0, 0, 0, 68, 59, 1, 0, 0, 0,
		68, 62, 1, 0, 0, 0, 68, 65, 1, 0, 0, 0, 69, 8, 1, 0, 0, 0, 70, 71, 5, 112,
		0, 0, 71, 78, 5, 105, 0, 0, 72, 78, 5, 960, 0, 0, 73, 74, 5, 112, 0, 0,
		74, 75, 5, 104, 0, 0, 75, 78, 5, 105, 0, 0, 76, 78, 7, 0, 0, 0, 77, 70,
		1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 77, 73, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0,
		78, 10, 1, 0, 0, 0, 79, 80, 7, 1, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5,
		43, 0, 0, 82, 14, 1, 0, 0, 0, 83, 84, 5, 45, 0, 0, 84, 16, 1, 0, 0, 0,
		85, 86, 5, 42, 0, 0, 86, 18, 1, 0, 0, 0, 87, 88, 5, 47, 0, 0, 88, 20, 1,
		0, 0, 0, 89, 90, 5, 94, 0, 0, 90, 22, 1, 0, 0, 0, 91, 92, 5, 42, 0, 0,
		92, 93, 5, 42, 0, 0, 93, 24, 1, 0, 0, 0, 94, 95, 5, 109, 0, 0, 95, 96,
		5, 111, 0, 0, 96, 97, 5, 100, 0, 0, 97, 26, 1, 0, 0, 0, 98, 99, 5, 40,
		0, 0, 99, 28, 1, 0, 0, 0, 100, 101, 5, 41, 0, 0, 101, 30, 1, 0, 0, 0, 102,
		103, 5, 37, 0, 0, 103, 32, 1, 0, 0, 0, 104, 107, 3, 35, 17, 0, 105, 107,
		7, 2, 0, 0, 106, 104, 1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 34, 1, 0, 0, 0,
		110, 111, 2, 48, 57, 0, 111, 36, 1, 0, 0, 0, 112, 114, 7, 3, 0, 0, 113,
		112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 38, 1, 0, 0, 0, 117, 119, 7, 4, 0, 0, 118, 117, 1, 0,
		0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 123, 6, 19, 0, 0, 123, 40, 1, 0, 0, 0, 7, 0,
		68, 77, 106, 108, 115, 120, 1, 6, 0, 0,
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

// CalculatorLexerInit initializes any static state used to implement CalculatorLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalculatorLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalculatorLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.once.Do(calculatorlexerLexerInit)
}

// NewCalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {
	CalculatorLexerInit()
	l := new(CalculatorLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalculatorLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerT__0            = 1
	CalculatorLexerT__1            = 2
	CalculatorLexerT__2            = 3
	CalculatorLexerFUNC_NAME       = 4
	CalculatorLexerCONSTANT        = 5
	CalculatorLexerCURRENCY_SYMBOL = 6
	CalculatorLexerADD             = 7
	CalculatorLexerSUB             = 8
	CalculatorLexerMUL             = 9
	CalculatorLexerDIV             = 10
	CalculatorLexerPOW             = 11
	CalculatorLexerPPOW            = 12
	CalculatorLexerMOD             = 13
	CalculatorLexerLPAREN          = 14
	CalculatorLexerRPAREN          = 15
	CalculatorLexerPCT             = 16
	CalculatorLexerNUMBER          = 17
	CalculatorLexerUNIT_NAME       = 18
	CalculatorLexerWS              = 19
)

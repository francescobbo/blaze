// Code generated from calculator/parser/Calculator.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 111,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	59, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 68, 10, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 6, 16, 96, 10, 16, 13, 16, 14, 16, 97, 3, 17,
	6, 17, 101, 10, 17, 13, 17, 14, 17, 102, 3, 18, 6, 18, 106, 10, 18, 13,
	18, 14, 18, 107, 3, 18, 3, 18, 2, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 3, 2, 7, 4, 2, 103, 103, 968, 968, 5, 2, 38, 38, 165,
	165, 8366, 8366, 5, 2, 46, 46, 48, 48, 50, 59, 3, 2, 99, 124, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 121, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2,
	2, 5, 58, 3, 2, 2, 2, 7, 67, 3, 2, 2, 2, 9, 69, 3, 2, 2, 2, 11, 71, 3,
	2, 2, 2, 13, 73, 3, 2, 2, 2, 15, 75, 3, 2, 2, 2, 17, 77, 3, 2, 2, 2, 19,
	79, 3, 2, 2, 2, 21, 81, 3, 2, 2, 2, 23, 84, 3, 2, 2, 2, 25, 88, 3, 2, 2,
	2, 27, 90, 3, 2, 2, 2, 29, 92, 3, 2, 2, 2, 31, 95, 3, 2, 2, 2, 33, 100,
	3, 2, 2, 2, 35, 105, 3, 2, 2, 2, 37, 38, 7, 113, 2, 2, 38, 39, 7, 104,
	2, 2, 39, 4, 3, 2, 2, 2, 40, 41, 7, 117, 2, 2, 41, 42, 7, 115, 2, 2, 42,
	43, 7, 116, 2, 2, 43, 59, 7, 118, 2, 2, 44, 45, 7, 110, 2, 2, 45, 46, 7,
	113, 2, 2, 46, 59, 7, 105, 2, 2, 47, 48, 7, 110, 2, 2, 48, 59, 7, 112,
	2, 2, 49, 50, 7, 117, 2, 2, 50, 51, 7, 107, 2, 2, 51, 59, 7, 112, 2, 2,
	52, 53, 7, 101, 2, 2, 53, 54, 7, 113, 2, 2, 54, 59, 7, 117, 2, 2, 55, 56,
	7, 118, 2, 2, 56, 57, 7, 99, 2, 2, 57, 59, 7, 112, 2, 2, 58, 40, 3, 2,
	2, 2, 58, 44, 3, 2, 2, 2, 58, 47, 3, 2, 2, 2, 58, 49, 3, 2, 2, 2, 58, 52,
	3, 2, 2, 2, 58, 55, 3, 2, 2, 2, 59, 6, 3, 2, 2, 2, 60, 61, 7, 114, 2, 2,
	61, 68, 7, 107, 2, 2, 62, 68, 7, 962, 2, 2, 63, 64, 7, 114, 2, 2, 64, 65,
	7, 106, 2, 2, 65, 68, 7, 107, 2, 2, 66, 68, 9, 2, 2, 2, 67, 60, 3, 2, 2,
	2, 67, 62, 3, 2, 2, 2, 67, 63, 3, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 8, 3,
	2, 2, 2, 69, 70, 9, 3, 2, 2, 70, 10, 3, 2, 2, 2, 71, 72, 7, 45, 2, 2, 72,
	12, 3, 2, 2, 2, 73, 74, 7, 47, 2, 2, 74, 14, 3, 2, 2, 2, 75, 76, 7, 44,
	2, 2, 76, 16, 3, 2, 2, 2, 77, 78, 7, 49, 2, 2, 78, 18, 3, 2, 2, 2, 79,
	80, 7, 96, 2, 2, 80, 20, 3, 2, 2, 2, 81, 82, 7, 44, 2, 2, 82, 83, 7, 44,
	2, 2, 83, 22, 3, 2, 2, 2, 84, 85, 7, 111, 2, 2, 85, 86, 7, 113, 2, 2, 86,
	87, 7, 102, 2, 2, 87, 24, 3, 2, 2, 2, 88, 89, 7, 42, 2, 2, 89, 26, 3, 2,
	2, 2, 90, 91, 7, 43, 2, 2, 91, 28, 3, 2, 2, 2, 92, 93, 7, 39, 2, 2, 93,
	30, 3, 2, 2, 2, 94, 96, 9, 4, 2, 2, 95, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2,
	2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 32, 3, 2, 2, 2, 99, 101,
	9, 5, 2, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2,
	2, 2, 102, 103, 3, 2, 2, 2, 103, 34, 3, 2, 2, 2, 104, 106, 9, 6, 2, 2,
	105, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107,
	108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 8, 18, 2, 2, 110, 36,
	3, 2, 2, 2, 8, 2, 58, 67, 97, 102, 107, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'of'", "", "", "", "'+'", "'-'", "'*'", "'/'", "'^'", "'**'", "'mod'",
	"'('", "')'", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "FUNC_NAME", "CONSTANT", "CURRENCY_SYMBOL", "ADD", "SUB", "MUL",
	"DIV", "POW", "PPOW", "MOD", "LPAREN", "RPAREN", "PCT", "NUMBER", "UNIT_NAME",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "FUNC_NAME", "CONSTANT", "CURRENCY_SYMBOL", "ADD", "SUB", "MUL",
	"DIV", "POW", "PPOW", "MOD", "LPAREN", "RPAREN", "PCT", "NUMBER", "UNIT_NAME",
	"WS",
}

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CalculatorLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {
	l := new(CalculatorLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerT__0            = 1
	CalculatorLexerFUNC_NAME       = 2
	CalculatorLexerCONSTANT        = 3
	CalculatorLexerCURRENCY_SYMBOL = 4
	CalculatorLexerADD             = 5
	CalculatorLexerSUB             = 6
	CalculatorLexerMUL             = 7
	CalculatorLexerDIV             = 8
	CalculatorLexerPOW             = 9
	CalculatorLexerPPOW            = 10
	CalculatorLexerMOD             = 11
	CalculatorLexerLPAREN          = 12
	CalculatorLexerRPAREN          = 13
	CalculatorLexerPCT             = 14
	CalculatorLexerNUMBER          = 15
	CalculatorLexerUNIT_NAME       = 16
	CalculatorLexerWS              = 17
)

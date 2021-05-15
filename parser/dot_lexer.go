// Code generated from Dot.g4 by ANTLR 4.9. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 158,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 11, 6, 11, 98, 10, 11, 13, 11, 14, 11, 99, 3,
	12, 3, 12, 3, 12, 3, 12, 7, 12, 106, 10, 12, 12, 12, 14, 12, 109, 11, 12,
	3, 12, 3, 12, 3, 13, 6, 13, 114, 10, 13, 13, 13, 14, 13, 115, 3, 14, 6,
	14, 119, 10, 14, 13, 14, 14, 14, 120, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 2,
	2, 29, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 3,
	2, 6, 3, 2, 50, 59, 5, 2, 12, 12, 15, 15, 36, 36, 5, 2, 50, 59, 67, 92,
	99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 162, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 3, 57, 3, 2, 2, 2, 5, 59,
	3, 2, 2, 2, 7, 61, 3, 2, 2, 2, 9, 63, 3, 2, 2, 2, 11, 65, 3, 2, 2, 2, 13,
	67, 3, 2, 2, 2, 15, 74, 3, 2, 2, 2, 17, 81, 3, 2, 2, 2, 19, 93, 3, 2, 2,
	2, 21, 97, 3, 2, 2, 2, 23, 101, 3, 2, 2, 2, 25, 113, 3, 2, 2, 2, 27, 118,
	3, 2, 2, 2, 29, 124, 3, 2, 2, 2, 31, 126, 3, 2, 2, 2, 33, 128, 3, 2, 2,
	2, 35, 130, 3, 2, 2, 2, 37, 132, 3, 2, 2, 2, 39, 135, 3, 2, 2, 2, 41, 138,
	3, 2, 2, 2, 43, 141, 3, 2, 2, 2, 45, 144, 3, 2, 2, 2, 47, 146, 3, 2, 2,
	2, 49, 148, 3, 2, 2, 2, 51, 150, 3, 2, 2, 2, 53, 153, 3, 2, 2, 2, 55, 156,
	3, 2, 2, 2, 57, 58, 7, 42, 2, 2, 58, 4, 3, 2, 2, 2, 59, 60, 7, 43, 2, 2,
	60, 6, 3, 2, 2, 2, 61, 62, 7, 125, 2, 2, 62, 8, 3, 2, 2, 2, 63, 64, 7,
	127, 2, 2, 64, 10, 3, 2, 2, 2, 65, 66, 7, 48, 2, 2, 66, 12, 3, 2, 2, 2,
	67, 68, 7, 80, 2, 2, 68, 69, 7, 119, 2, 2, 69, 70, 7, 111, 2, 2, 70, 71,
	7, 100, 2, 2, 71, 72, 7, 103, 2, 2, 72, 73, 7, 116, 2, 2, 73, 14, 3, 2,
	2, 2, 74, 75, 7, 85, 2, 2, 75, 76, 7, 118, 2, 2, 76, 77, 7, 116, 2, 2,
	77, 78, 7, 107, 2, 2, 78, 79, 7, 112, 2, 2, 79, 80, 7, 105, 2, 2, 80, 16,
	3, 2, 2, 2, 81, 82, 7, 111, 2, 2, 82, 83, 7, 117, 2, 2, 83, 84, 7, 105,
	2, 2, 84, 85, 7, 48, 2, 2, 85, 86, 7, 101, 2, 2, 86, 87, 7, 113, 2, 2,
	87, 88, 7, 112, 2, 2, 88, 89, 7, 118, 2, 2, 89, 90, 7, 103, 2, 2, 90, 91,
	7, 112, 2, 2, 91, 92, 7, 118, 2, 2, 92, 18, 3, 2, 2, 2, 93, 94, 7, 107,
	2, 2, 94, 95, 7, 104, 2, 2, 95, 20, 3, 2, 2, 2, 96, 98, 9, 2, 2, 2, 97,
	96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2,
	2, 2, 100, 22, 3, 2, 2, 2, 101, 107, 7, 36, 2, 2, 102, 106, 10, 3, 2, 2,
	103, 104, 7, 36, 2, 2, 104, 106, 7, 36, 2, 2, 105, 102, 3, 2, 2, 2, 105,
	103, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108,
	3, 2, 2, 2, 108, 110, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 111, 7, 36,
	2, 2, 111, 24, 3, 2, 2, 2, 112, 114, 9, 4, 2, 2, 113, 112, 3, 2, 2, 2,
	114, 115, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116,
	26, 3, 2, 2, 2, 117, 119, 9, 5, 2, 2, 118, 117, 3, 2, 2, 2, 119, 120, 3,
	2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2, 2,
	2, 122, 123, 8, 14, 2, 2, 123, 28, 3, 2, 2, 2, 124, 125, 7, 44, 2, 2, 125,
	30, 3, 2, 2, 2, 126, 127, 7, 49, 2, 2, 127, 32, 3, 2, 2, 2, 128, 129, 7,
	45, 2, 2, 129, 34, 3, 2, 2, 2, 130, 131, 7, 47, 2, 2, 131, 36, 3, 2, 2,
	2, 132, 133, 7, 63, 2, 2, 133, 134, 7, 63, 2, 2, 134, 38, 3, 2, 2, 2, 135,
	136, 7, 35, 2, 2, 136, 137, 7, 63, 2, 2, 137, 40, 3, 2, 2, 2, 138, 139,
	7, 64, 2, 2, 139, 140, 7, 63, 2, 2, 140, 42, 3, 2, 2, 2, 141, 142, 7, 62,
	2, 2, 142, 143, 7, 63, 2, 2, 143, 44, 3, 2, 2, 2, 144, 145, 7, 35, 2, 2,
	145, 46, 3, 2, 2, 2, 146, 147, 7, 64, 2, 2, 147, 48, 3, 2, 2, 2, 148, 149,
	7, 62, 2, 2, 149, 50, 3, 2, 2, 2, 150, 151, 7, 40, 2, 2, 151, 152, 7, 40,
	2, 2, 152, 52, 3, 2, 2, 2, 153, 154, 7, 126, 2, 2, 154, 155, 7, 126, 2,
	2, 155, 54, 3, 2, 2, 2, 156, 157, 7, 12, 2, 2, 157, 56, 3, 2, 2, 2, 8,
	2, 99, 105, 107, 115, 120, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'.'", "'Number'", "'String'", "'msg.content'",
	"'if'", "", "", "", "", "'*'", "'/'", "'+'", "'-'", "'=='", "'!='", "'>='",
	"'<='", "'!'", "'>'", "'<'", "'&&'", "'||'", "'\n'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "MSGCONTENT", "IF", "INT", "STRING", "WORDS",
	"WS", "MUL", "DIV", "ADD", "SUB", "Equals", "NEquals", "GTEquals", "LTEquals",
	"Excl", "GT", "LT", "And", "Or", "NL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "MSGCONTENT", "IF",
	"INT", "STRING", "WORDS", "WS", "MUL", "DIV", "ADD", "SUB", "Equals", "NEquals",
	"GTEquals", "LTEquals", "Excl", "GT", "LT", "And", "Or", "NL",
}

type DotLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewDotLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *DotLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDotLexer(input antlr.CharStream) *DotLexer {
	l := new(DotLexer)
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
	l.GrammarFileName = "Dot.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DotLexer tokens.
const (
	DotLexerT__0       = 1
	DotLexerT__1       = 2
	DotLexerT__2       = 3
	DotLexerT__3       = 4
	DotLexerT__4       = 5
	DotLexerT__5       = 6
	DotLexerT__6       = 7
	DotLexerMSGCONTENT = 8
	DotLexerIF         = 9
	DotLexerINT        = 10
	DotLexerSTRING     = 11
	DotLexerWORDS      = 12
	DotLexerWS         = 13
	DotLexerMUL        = 14
	DotLexerDIV        = 15
	DotLexerADD        = 16
	DotLexerSUB        = 17
	DotLexerEquals     = 18
	DotLexerNEquals    = 19
	DotLexerGTEquals   = 20
	DotLexerLTEquals   = 21
	DotLexerExcl       = 22
	DotLexerGT         = 23
	DotLexerLT         = 24
	DotLexerAnd        = 25
	DotLexerOr         = 26
	DotLexerNL         = 27
)

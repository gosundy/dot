// Code generated from Dot.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Dot

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 85, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 7, 2, 24, 10, 2,
	12, 2, 14, 2, 27, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 33, 10, 3, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	81, 10, 10, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 2, 7, 3, 2, 16, 17, 3, 2, 18, 19, 4, 2, 20, 23, 25, 26, 3, 2, 27,
	28, 3, 2, 8, 9, 2, 84, 2, 25, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 34, 3,
	2, 2, 2, 8, 37, 3, 2, 2, 2, 10, 62, 3, 2, 2, 2, 12, 64, 3, 2, 2, 2, 14,
	68, 3, 2, 2, 2, 16, 72, 3, 2, 2, 2, 18, 80, 3, 2, 2, 2, 20, 82, 3, 2, 2,
	2, 22, 24, 5, 4, 3, 2, 23, 22, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23,
	3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 3, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2,
	28, 29, 5, 6, 4, 2, 29, 30, 5, 12, 7, 2, 30, 33, 3, 2, 2, 2, 31, 33, 5,
	10, 6, 2, 32, 28, 3, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34,
	35, 7, 11, 2, 2, 35, 36, 5, 8, 5, 2, 36, 7, 3, 2, 2, 2, 37, 38, 7, 3, 2,
	2, 38, 39, 5, 10, 6, 2, 39, 40, 7, 4, 2, 2, 40, 9, 3, 2, 2, 2, 41, 42,
	5, 14, 8, 2, 42, 43, 9, 2, 2, 2, 43, 44, 5, 10, 6, 2, 44, 63, 3, 2, 2,
	2, 45, 46, 5, 14, 8, 2, 46, 47, 9, 3, 2, 2, 47, 48, 5, 10, 6, 2, 48, 63,
	3, 2, 2, 2, 49, 63, 5, 14, 8, 2, 50, 51, 7, 24, 2, 2, 51, 63, 5, 10, 6,
	2, 52, 53, 5, 14, 8, 2, 53, 54, 9, 4, 2, 2, 54, 55, 5, 10, 6, 2, 55, 63,
	3, 2, 2, 2, 56, 57, 5, 14, 8, 2, 57, 58, 9, 5, 2, 2, 58, 59, 5, 10, 6,
	2, 59, 63, 3, 2, 2, 2, 60, 63, 7, 12, 2, 2, 61, 63, 7, 13, 2, 2, 62, 41,
	3, 2, 2, 2, 62, 45, 3, 2, 2, 2, 62, 49, 3, 2, 2, 2, 62, 50, 3, 2, 2, 2,
	62, 52, 3, 2, 2, 2, 62, 56, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3,
	2, 2, 2, 63, 11, 3, 2, 2, 2, 64, 65, 7, 5, 2, 2, 65, 66, 5, 10, 6, 2, 66,
	67, 7, 6, 2, 2, 67, 13, 3, 2, 2, 2, 68, 69, 5, 16, 9, 2, 69, 70, 7, 7,
	2, 2, 70, 71, 5, 20, 11, 2, 71, 15, 3, 2, 2, 2, 72, 73, 7, 10, 2, 2, 73,
	74, 7, 7, 2, 2, 74, 75, 5, 18, 10, 2, 75, 17, 3, 2, 2, 2, 76, 77, 7, 14,
	2, 2, 77, 78, 7, 7, 2, 2, 78, 81, 5, 18, 10, 2, 79, 81, 7, 14, 2, 2, 80,
	76, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 19, 3, 2, 2, 2, 82, 83, 9, 6, 2,
	2, 83, 21, 3, 2, 2, 2, 6, 25, 32, 62, 80,
}
var literalNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'.'", "'Number'", "'String'", "'msg.content'",
	"'if'", "", "", "", "", "'*'", "'/'", "'+'", "'-'", "'=='", "'!='", "'>='",
	"'<='", "'!'", "'>'", "'<'", "'&&'", "'||'", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "MSGCONTENT", "IF", "INT", "STRING", "WORDS",
	"WS", "MUL", "DIV", "ADD", "SUB", "Equals", "NEquals", "GTEquals", "LTEquals",
	"Excl", "GT", "LT", "And", "Or", "NL",
}

var ruleNames = []string{
	"start", "statment", "ifStatment", "ifPartStatment", "exp", "block", "msgContentValue",
	"msgContentJsonPath", "jsonPath", "msgValueType",
}

type DotParser struct {
	*antlr.BaseParser
}

// NewDotParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DotParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDotParser(input antlr.TokenStream) *DotParser {
	this := new(DotParser)
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
	this.GrammarFileName = "Dot.g4"

	return this
}

// DotParser tokens.
const (
	DotParserEOF        = antlr.TokenEOF
	DotParserT__0       = 1
	DotParserT__1       = 2
	DotParserT__2       = 3
	DotParserT__3       = 4
	DotParserT__4       = 5
	DotParserT__5       = 6
	DotParserT__6       = 7
	DotParserMSGCONTENT = 8
	DotParserIF         = 9
	DotParserINT        = 10
	DotParserSTRING     = 11
	DotParserWORDS      = 12
	DotParserWS         = 13
	DotParserMUL        = 14
	DotParserDIV        = 15
	DotParserADD        = 16
	DotParserSUB        = 17
	DotParserEquals     = 18
	DotParserNEquals    = 19
	DotParserGTEquals   = 20
	DotParserLTEquals   = 21
	DotParserExcl       = 22
	DotParserGT         = 23
	DotParserLT         = 24
	DotParserAnd        = 25
	DotParserOr         = 26
	DotParserNL         = 27
)

// DotParser rules.
const (
	DotParserRULE_start              = 0
	DotParserRULE_statment           = 1
	DotParserRULE_ifStatment         = 2
	DotParserRULE_ifPartStatment     = 3
	DotParserRULE_exp                = 4
	DotParserRULE_block              = 5
	DotParserRULE_msgContentValue    = 6
	DotParserRULE_msgContentJsonPath = 7
	DotParserRULE_jsonPath           = 8
	DotParserRULE_msgValueType       = 9
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllStatment() []IStatmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatmentContext)(nil)).Elem())
	var tst = make([]IStatmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatmentContext)
		}
	}

	return tst
}

func (s *StartContext) Statment(i int) IStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatmentContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DotParserRULE_start)
	var _la int

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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserMSGCONTENT)|(1<<DotParserIF)|(1<<DotParserINT)|(1<<DotParserSTRING)|(1<<DotParserExcl))) != 0 {
		{
			p.SetState(20)
			p.Statment()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatmentContext is an interface to support dynamic dispatch.
type IStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatmentContext differentiates from other interfaces.
	IsStatmentContext()
}

type StatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatmentContext() *StatmentContext {
	var p = new(StatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_statment
	return p
}

func (*StatmentContext) IsStatmentContext() {}

func NewStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatmentContext {
	var p = new(StatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_statment

	return p
}

func (s *StatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *StatmentContext) CopyFrom(ctx *StatmentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatmentExpContext struct {
	*StatmentContext
}

func NewStatmentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatmentExpContext {
	var p = new(StatmentExpContext)

	p.StatmentContext = NewEmptyStatmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatmentContext))

	return p
}

func (s *StatmentExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatmentExpContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatmentExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitStatmentExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatmentIfContext struct {
	*StatmentContext
}

func NewStatmentIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatmentIfContext {
	var p = new(StatmentIfContext)

	p.StatmentContext = NewEmptyStatmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatmentContext))

	return p
}

func (s *StatmentIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatmentIfContext) IfStatment() IIfStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatmentContext)
}

func (s *StatmentIfContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatmentIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitStatmentIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Statment() (localctx IStatmentContext) {
	localctx = NewStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DotParserRULE_statment)

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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DotParserIF:
		localctx = NewStatmentIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.IfStatment()
		}
		{
			p.SetState(27)
			p.Block()
		}

	case DotParserMSGCONTENT, DotParserINT, DotParserSTRING, DotParserExcl:
		localctx = NewStatmentExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Exp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatmentContext is an interface to support dynamic dispatch.
type IIfStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatmentContext differentiates from other interfaces.
	IsIfStatmentContext()
}

type IfStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatmentContext() *IfStatmentContext {
	var p = new(IfStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_ifStatment
	return p
}

func (*IfStatmentContext) IsIfStatmentContext() {}

func NewIfStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatmentContext {
	var p = new(IfStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_ifStatment

	return p
}

func (s *IfStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatmentContext) IF() antlr.TerminalNode {
	return s.GetToken(DotParserIF, 0)
}

func (s *IfStatmentContext) IfPartStatment() IIfPartStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfPartStatmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfPartStatmentContext)
}

func (s *IfStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitIfStatment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) IfStatment() (localctx IIfStatmentContext) {
	localctx = NewIfStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DotParserRULE_ifStatment)

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
		p.SetState(32)
		p.Match(DotParserIF)
	}
	{
		p.SetState(33)
		p.IfPartStatment()
	}

	return localctx
}

// IIfPartStatmentContext is an interface to support dynamic dispatch.
type IIfPartStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfPartStatmentContext differentiates from other interfaces.
	IsIfPartStatmentContext()
}

type IfPartStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfPartStatmentContext() *IfPartStatmentContext {
	var p = new(IfPartStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_ifPartStatment
	return p
}

func (*IfPartStatmentContext) IsIfPartStatmentContext() {}

func NewIfPartStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfPartStatmentContext {
	var p = new(IfPartStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_ifPartStatment

	return p
}

func (s *IfPartStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *IfPartStatmentContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IfPartStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfPartStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfPartStatmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitIfPartStatment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) IfPartStatment() (localctx IIfPartStatmentContext) {
	localctx = NewIfPartStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DotParserRULE_ifPartStatment)

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
		p.SetState(35)
		p.Match(DotParserT__0)
	}
	{
		p.SetState(36)
		p.Exp()
	}
	{
		p.SetState(37)
		p.Match(DotParserT__1)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) CopyFrom(ctx *ExpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotContext struct {
	*ExpContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) Excl() antlr.TerminalNode {
	return s.GetToken(DotParserExcl, 0)
}

func (s *NotContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContentAddSubContext struct {
	*ExpContext
	op antlr.Token
}

func NewContentAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContentAddSubContext {
	var p = new(ContentAddSubContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *ContentAddSubContext) GetOp() antlr.Token { return s.op }

func (s *ContentAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ContentAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentAddSubContext) MsgContentValue() IMsgContentValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContentValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContentValueContext)
}

func (s *ContentAddSubContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ContentAddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(DotParserSUB, 0)
}

func (s *ContentAddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(DotParserADD, 0)
}

func (s *ContentAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitContentAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetIntContext struct {
	*ExpContext
}

func NewGetIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetIntContext {
	var p = new(GetIntContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *GetIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetIntContext) INT() antlr.TerminalNode {
	return s.GetToken(DotParserINT, 0)
}

func (s *GetIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitGetInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetStringContext struct {
	*ExpContext
}

func NewGetStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetStringContext {
	var p = new(GetStringContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *GetStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(DotParserSTRING, 0)
}

func (s *GetStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitGetString(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareContext struct {
	*ExpContext
	op antlr.Token
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *CompareContext) GetOp() antlr.Token { return s.op }

func (s *CompareContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) MsgContentValue() IMsgContentValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContentValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContentValueContext)
}

func (s *CompareContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *CompareContext) Equals() antlr.TerminalNode {
	return s.GetToken(DotParserEquals, 0)
}

func (s *CompareContext) NEquals() antlr.TerminalNode {
	return s.GetToken(DotParserNEquals, 0)
}

func (s *CompareContext) GTEquals() antlr.TerminalNode {
	return s.GetToken(DotParserGTEquals, 0)
}

func (s *CompareContext) LTEquals() antlr.TerminalNode {
	return s.GetToken(DotParserLTEquals, 0)
}

func (s *CompareContext) GT() antlr.TerminalNode {
	return s.GetToken(DotParserGT, 0)
}

func (s *CompareContext) LT() antlr.TerminalNode {
	return s.GetToken(DotParserLT, 0)
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetContentValueContext struct {
	*ExpContext
}

func NewGetContentValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetContentValueContext {
	var p = new(GetContentValueContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *GetContentValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetContentValueContext) MsgContentValue() IMsgContentValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContentValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContentValueContext)
}

func (s *GetContentValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitGetContentValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndOrContext struct {
	*ExpContext
	op antlr.Token
}

func NewAndOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndOrContext {
	var p = new(AndOrContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *AndOrContext) GetOp() antlr.Token { return s.op }

func (s *AndOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOrContext) MsgContentValue() IMsgContentValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContentValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContentValueContext)
}

func (s *AndOrContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AndOrContext) And() antlr.TerminalNode {
	return s.GetToken(DotParserAnd, 0)
}

func (s *AndOrContext) Or() antlr.TerminalNode {
	return s.GetToken(DotParserOr, 0)
}

func (s *AndOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitAndOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContentMulDivContext struct {
	*ExpContext
	op antlr.Token
}

func NewContentMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContentMulDivContext {
	var p = new(ContentMulDivContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *ContentMulDivContext) GetOp() antlr.Token { return s.op }

func (s *ContentMulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *ContentMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentMulDivContext) MsgContentValue() IMsgContentValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContentValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContentValueContext)
}

func (s *ContentMulDivContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ContentMulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(DotParserDIV, 0)
}

func (s *ContentMulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(DotParserMUL, 0)
}

func (s *ContentMulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitContentMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DotParserRULE_exp)
	var _la int

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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewContentMulDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.MsgContentValue()
		}
		{
			p.SetState(40)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ContentMulDivContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DotParserMUL || _la == DotParserDIV) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ContentMulDivContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(41)
			p.Exp()
		}

	case 2:
		localctx = NewContentAddSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.MsgContentValue()
		}
		{
			p.SetState(44)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ContentAddSubContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DotParserADD || _la == DotParserSUB) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ContentAddSubContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(45)
			p.Exp()
		}

	case 3:
		localctx = NewGetContentValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)
			p.MsgContentValue()
		}

	case 4:
		localctx = NewNotContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(48)
			p.Match(DotParserExcl)
		}
		{
			p.SetState(49)
			p.Exp()
		}

	case 5:
		localctx = NewCompareContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.MsgContentValue()
		}
		{
			p.SetState(51)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserEquals)|(1<<DotParserNEquals)|(1<<DotParserGTEquals)|(1<<DotParserLTEquals)|(1<<DotParserGT)|(1<<DotParserLT))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(52)
			p.Exp()
		}

	case 6:
		localctx = NewAndOrContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(54)
			p.MsgContentValue()
		}
		{
			p.SetState(55)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AndOrContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DotParserAnd || _la == DotParserOr) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AndOrContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(56)
			p.Exp()
		}

	case 7:
		localctx = NewGetIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(58)
			p.Match(DotParserINT)
		}

	case 8:
		localctx = NewGetStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(59)
			p.Match(DotParserSTRING)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DotParserRULE_block)

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
		p.SetState(62)
		p.Match(DotParserT__2)
	}
	{
		p.SetState(63)
		p.Exp()
	}
	{
		p.SetState(64)
		p.Match(DotParserT__3)
	}

	return localctx
}

// IMsgContentValueContext is an interface to support dynamic dispatch.
type IMsgContentValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgContentValueContext differentiates from other interfaces.
	IsMsgContentValueContext()
}

type MsgContentValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgContentValueContext() *MsgContentValueContext {
	var p = new(MsgContentValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_msgContentValue
	return p
}

func (*MsgContentValueContext) IsMsgContentValueContext() {}

func NewMsgContentValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgContentValueContext {
	var p = new(MsgContentValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_msgContentValue

	return p
}

func (s *MsgContentValueContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgContentValueContext) MsgContentJsonPath() IMsgContentJsonPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContentJsonPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContentJsonPathContext)
}

func (s *MsgContentValueContext) MsgValueType() IMsgValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgValueTypeContext)
}

func (s *MsgContentValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgContentValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgContentValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitMsgContentValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) MsgContentValue() (localctx IMsgContentValueContext) {
	localctx = NewMsgContentValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DotParserRULE_msgContentValue)

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
		p.SetState(66)
		p.MsgContentJsonPath()
	}
	{
		p.SetState(67)
		p.Match(DotParserT__4)
	}
	{
		p.SetState(68)
		p.MsgValueType()
	}

	return localctx
}

// IMsgContentJsonPathContext is an interface to support dynamic dispatch.
type IMsgContentJsonPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgContentJsonPathContext differentiates from other interfaces.
	IsMsgContentJsonPathContext()
}

type MsgContentJsonPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgContentJsonPathContext() *MsgContentJsonPathContext {
	var p = new(MsgContentJsonPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_msgContentJsonPath
	return p
}

func (*MsgContentJsonPathContext) IsMsgContentJsonPathContext() {}

func NewMsgContentJsonPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgContentJsonPathContext {
	var p = new(MsgContentJsonPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_msgContentJsonPath

	return p
}

func (s *MsgContentJsonPathContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgContentJsonPathContext) MSGCONTENT() antlr.TerminalNode {
	return s.GetToken(DotParserMSGCONTENT, 0)
}

func (s *MsgContentJsonPathContext) JsonPath() IJsonPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonPathContext)
}

func (s *MsgContentJsonPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgContentJsonPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgContentJsonPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitMsgContentJsonPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) MsgContentJsonPath() (localctx IMsgContentJsonPathContext) {
	localctx = NewMsgContentJsonPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DotParserRULE_msgContentJsonPath)

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
		p.SetState(70)
		p.Match(DotParserMSGCONTENT)
	}
	{
		p.SetState(71)
		p.Match(DotParserT__4)
	}
	{
		p.SetState(72)
		p.JsonPath()
	}

	return localctx
}

// IJsonPathContext is an interface to support dynamic dispatch.
type IJsonPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonPathContext differentiates from other interfaces.
	IsJsonPathContext()
}

type JsonPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonPathContext() *JsonPathContext {
	var p = new(JsonPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_jsonPath
	return p
}

func (*JsonPathContext) IsJsonPathContext() {}

func NewJsonPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonPathContext {
	var p = new(JsonPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_jsonPath

	return p
}

func (s *JsonPathContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonPathContext) WORDS() antlr.TerminalNode {
	return s.GetToken(DotParserWORDS, 0)
}

func (s *JsonPathContext) JsonPath() IJsonPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonPathContext)
}

func (s *JsonPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitJsonPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) JsonPath() (localctx IJsonPathContext) {
	localctx = NewJsonPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DotParserRULE_jsonPath)

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(DotParserWORDS)
		}
		{
			p.SetState(75)
			p.Match(DotParserT__4)
		}
		{
			p.SetState(76)
			p.JsonPath()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(DotParserWORDS)
		}

	}

	return localctx
}

// IMsgValueTypeContext is an interface to support dynamic dispatch.
type IMsgValueTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgValueTypeContext differentiates from other interfaces.
	IsMsgValueTypeContext()
}

type MsgValueTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgValueTypeContext() *MsgValueTypeContext {
	var p = new(MsgValueTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_msgValueType
	return p
}

func (*MsgValueTypeContext) IsMsgValueTypeContext() {}

func NewMsgValueTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgValueTypeContext {
	var p = new(MsgValueTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_msgValueType

	return p
}

func (s *MsgValueTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *MsgValueTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgValueTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgValueTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitMsgValueType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) MsgValueType() (localctx IMsgValueTypeContext) {
	localctx = NewMsgValueTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DotParserRULE_msgValueType)
	var _la int

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
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DotParserT__5 || _la == DotParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

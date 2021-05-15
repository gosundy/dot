package dot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gosundy/dot/parser"

	"github.com/Jeffail/gabs/v2"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var ErrNotFoundKey = func(key string) error {
	return fmt.Errorf("key:%s not found", key)
}
var ErrDivZero = errors.New("div zero")
var ErrNotSupportType = errors.New("not support type")
var ErrTypeNotMatch = func(type1 string, type2 string) error {
	return fmt.Errorf("type not match, type1:%s,type2:%s", type1, type2)
}
var ErrIfNotMatch = errors.New("if not match")
var ErrNotStatement = errors.New("not match any statement")
var ErrValueTypeUnknown = func(unknownType string) error { return fmt.Errorf("type:%s is not define", unknownType) }
var ErrCompare = func(typ1 string, typ2 string) error { return fmt.Errorf("type1:%s can't compared type2:%s", typ1, typ2) }
var ErrTypeCantCompare = func(typ string) error { return fmt.Errorf("type:%s can't compared", typ) }

type ValueType int
type ValueFunc func() *Value

const (
	UnKnown ValueType = iota
	Number
	String
	Bool
)

var ValueTypeMap = map[string]ValueType{"Number": Number, "String": String}
var ValueTypeMapStr = map[ValueType]string{Number: "Number", String: "String", UnKnown: "unKnown"}

type Value struct {
	Value interface{}
	typ   ValueType
	Err   error
}
type MsgInterface interface {
	GetContentValueNumber(path string) (value int64, err error)
	SetContentValueNumber(path string, value int64) error
	GetContentValueString(path string) (value string, err error)
	SetContentValueString(path string, value string) error
}
type Msg struct {
	Content *gabs.Container
	Raw     []byte
	Meta    MsgMeta
}
type MsgMeta map[interface{}]interface{}

func NewMsg(rawJson []byte) (*Msg, error) {
	container, err := gabs.ParseJSON(rawJson)
	if err != nil {
		return nil, err
	}
	return &Msg{container, rawJson, make(MsgMeta)}, nil
}

func (msg *Msg) GetContentValueNumber(path string) (value int64, err error) {
	v, ok := msg.Content.Path(path).Data().(float64)
	if !ok {
		return 0, ErrNotFoundKey(path)
	}
	return int64(v), nil
}

func (msg *Msg) SetContentValueNumber(path string, value int64) error {
	_, err := msg.Content.SetP(value, path)
	return err
}

func (msg *Msg) GetContentValueString(path string) (value string, err error) {
	v, ok := msg.Content.Path(path).Data().(string)
	if !ok {
		return "", ErrNotFoundKey(path)
	}
	return v, nil
}

func (msg *Msg) SetContentValueString(path string, value string) error {
	_, err := msg.Content.SetP(value, path)
	return err
}

func NewDotVisitor() *DotVisitor {
	return &DotVisitor{}
}

type DotVisitor struct {
	*parser.BaseDotVisitor
	Msg MsgInterface
}

func (v *DotVisitor) visitRule(node antlr.RuleNode) interface{} {
	return node.Accept(v)
}

func (v *DotVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	exps := make([]func() *Value, len(ctx.AllStatment()))
	for idx, statementExp := range ctx.AllStatment() {
		exps[idx] = v.visitRule(statementExp).(func() *Value)
	}
	return func() *Value {
		for _, exp := range exps {
			value := exp()
			if value.Err != nil {
				if errors.Is(value.Err, ErrIfNotMatch) {
					continue
				}
				return value
			}
			return value
		}
		return &Value{Err: ErrNotStatement}
	}
}

func (v *DotVisitor) VisitContentMulDiv(ctx *parser.ContentMulDivContext) interface{} {
	valueFunc1 := v.visitRule(ctx.MsgContentValue()).(func() *Value)
	valueFunc2 := v.visitRule(ctx.Exp()).(func() *Value)
	return func() *Value {
		value1 := valueFunc1()
		if value1.Err != nil {
			return &Value{Err: value1.Err}
		}
		value2 := valueFunc2()
		if value2.Err != nil {
			return &Value{Err: value2.Err}
		}
		v1 := value1.Value.(int64)
		v2 := value2.Value.(int64)
		switch ctx.GetOp().GetTokenType() {
		case parser.DotParserMUL:
			return &Value{Value: v1 * v2, typ: Number}
		case parser.DotParserDIV:
			if value2.Value == 0 {
				return &Value{Err: ErrDivZero}
			}
			return &Value{Value: v1 / v2, typ: Number}
		default:
			return &Value{Err: ErrNotFoundKey(ctx.GetOp().GetText())}
		}
	}
}

func (v *DotVisitor) VisitContentAddSub(ctx *parser.ContentAddSubContext) interface{} {
	valueFunc1 := v.visitRule(ctx.MsgContentValue()).(func() *Value)
	valueFunc2 := v.visitRule(ctx.Exp()).(func() *Value)
	return func() *Value {
		value1 := valueFunc1()
		if value1.Err != nil {
			return &Value{Err: value1.Err}
		}
		value2 := valueFunc2()
		if value2.Err != nil {
			return &Value{Err: value2.Err}
		}
		v1 := value1.Value.(int64)
		v2 := value2.Value.(int64)
		switch ctx.GetOp().GetTokenType() {
		case parser.DotParserADD:
			return &Value{Value: v1 + v2, typ: Number}
		case parser.DotParserSUB:
			return &Value{Value: v1 - v2, typ: Number}
		default:
			return &Value{Err: ErrNotFoundKey(ctx.GetOp().GetText())}

		}
	}
}

func (v *DotVisitor) VisitGetContentValue(ctx *parser.GetContentValueContext) interface{} {
	return v.visitRule(ctx.MsgContentValue())
}

func (v *DotVisitor) VisitGetInt(ctx *parser.GetIntContext) interface{} {
	return func() *Value {
		v, err := strconv.ParseInt(ctx.GetText(), 10, 64)
		if err != nil {
			return &Value{Err: err}
		}
		return &Value{Value: v, typ: Number}
	}
}

func (v *DotVisitor) VisitGetString(ctx *parser.GetStringContext) interface{} {
	return func() *Value {
		return &Value{Value: ctx.GetText(), typ: String}
	}
}

func (v *DotVisitor) VisitMsgContentValue(ctx *parser.MsgContentValueContext) interface{} {
	jsonPath := v.visitRule(ctx.MsgContentJsonPath()).(string)
	valueType := v.visitRule(ctx.MsgValueType()).(string)
	return func() *Value {
		switch ValueTypeMap[valueType] {
		case Number:
			value, err := v.Msg.GetContentValueNumber(jsonPath)
			if err != nil {
				return &Value{Err: err}
			}
			return &Value{Value: value, typ: Number}
		case String:
			value, err := v.Msg.GetContentValueString(jsonPath)
			if err != nil {
				return &Value{Err: err}
			}
			return &Value{Value: value, typ: String}
		default:
			return &Value{Err: ErrValueTypeUnknown(valueType)}
		}
	}
}

func (v *DotVisitor) VisitMsgContentJsonPath(ctx *parser.MsgContentJsonPathContext) interface{} {
	return v.visitRule(ctx.JsonPath())
}

func (v *DotVisitor) VisitJsonPath(ctx *parser.JsonPathContext) interface{} {
	return ctx.GetText()
}

func (v *DotVisitor) VisitNot(ctx *parser.NotContext) interface{} {
	valueFunc := v.visitRule(ctx.Exp()).(func() *Value)
	return func() (bool, error) {
		cond, err := getBoolFunc(valueFunc)()
		if err != nil {
			return false, err
		}
		return !cond, nil
	}
}

func (v *DotVisitor) VisitStatmentIf(ctx *parser.StatmentIfContext) interface{} {
	ifStatement := v.visitRule(ctx.IfStatment()).(func() *Value)
	blockExp := v.visitRule(ctx.Block()).(func() *Value)
	return func() *Value {
		cond, err := getBoolFunc(ifStatement)()
		if err != nil {
			return &Value{Err: err}
		}
		if cond {
			return &Value{Value: blockExp().Value, typ: Number}
		}
		return &Value{Err: ErrIfNotMatch}
	}
}

func (v *DotVisitor) VisitStatmentExp(ctx *parser.StatmentExpContext) interface{} {
	return v.visitRule(ctx.Exp())
}

func (v *DotVisitor) VisitIfStatment(ctx *parser.IfStatmentContext) interface{} {
	return v.visitRule(ctx.IfPartStatment())
}

func (v *DotVisitor) VisitIfPartStatment(ctx *parser.IfPartStatmentContext) interface{} {
	return v.visitRule(ctx.Exp())
}

func (v *DotVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.visitRule(ctx.Exp())
}

func (v *DotVisitor) VisitCompare(ctx *parser.CompareContext) interface{} {
	leftValueFunc := v.visitRule(ctx.MsgContentValue()).(func() *Value)
	rightValueFunc := v.visitRule(ctx.Exp()).(func() *Value)
	return func() *Value {
		value1 := leftValueFunc()
		if value1.Err != nil {
			return value1
		}
		value2 := rightValueFunc()
		if value2.Err != nil {
			return value2
		}
		if value1.typ != value2.typ {
			return &Value{Err: ErrCompare(ValueTypeMapStr[value1.typ], ValueTypeMapStr[value2.typ])}
		}

		switch ctx.GetOp().GetTokenType() {
		case parser.DotLexerEquals:
			if value1.Value == value2.Value {
				return &Value{Value: true, typ: Bool}
			}
			return &Value{Value: false, typ: Bool}
		case parser.DotLexerNEquals:
			if value1.Value != value2.Value {
				return &Value{Value: true, typ: Bool}
			}
			return &Value{Value: false, typ: Bool}
		case parser.DotLexerGTEquals:
			if value1.typ == Number {
				if value1.Value.(int64) >= value2.Value.(int64) {
					return &Value{Value: true, typ: Bool}
				}
				return &Value{Value: false, typ: Bool}
			}
			return &Value{Err: ErrTypeCantCompare(ValueTypeMapStr[value1.typ])}
		case parser.DotLexerLTEquals:
			if value1.typ == Number {
				if value1.Value.(int64) <= value2.Value.(int64) {
					return &Value{Value: true, typ: Bool}
				}
				return &Value{Value: false, typ: Bool}
			}
			return &Value{Err: ErrTypeCantCompare(ValueTypeMapStr[value1.typ])}
		case parser.DotLexerGT:
			if value1.typ == Number {
				if value1.Value.(int64) > value2.Value.(int64) {
					return &Value{Value: true, typ: Bool}
				}
				return &Value{Value: false, typ: Bool}
			}
			return &Value{Err: ErrTypeCantCompare(ValueTypeMapStr[value1.typ])}
		case parser.DotLexerLT:
			if value1.typ == Number {
				if value1.Value.(int64) < value2.Value.(int64) {
					return &Value{Value: true, typ: Bool}
				}
				return &Value{Value: false, typ: Bool}
			}
			return &Value{Err: ErrTypeCantCompare(ValueTypeMapStr[value1.typ])}
		}
		return &Value{Err: ErrNotSupportType}
	}
}

func (v *DotVisitor) VisitAndOr(ctx *parser.AndOrContext) interface{} {
	leftValueFunc := v.visitRule(ctx.MsgContentValue()).(func() *Value)
	rightValueFunc := v.visitRule(ctx.Exp()).(func() *Value)
	return func() *Value {
		cond1, err := getBoolFunc(leftValueFunc)()
		if err != nil {
			return &Value{Err: err}
		}
		cond2, err := getBoolFunc(rightValueFunc)()
		if err != nil {
			return &Value{Err: err}
		}
		switch ctx.GetOp().GetTokenType() {
		case parser.DotParserAnd:
			if cond1 && cond2 {
				return &Value{Value: true, typ: Bool}
			}
			return &Value{Value: false, typ: Bool}
		case parser.DotParserOr:
			if cond1 || cond2 {
				return &Value{Value: true, typ: Bool}
			}
			return &Value{Value: false, typ: Bool}
		}
		return &Value{Err: ErrNotSupportType}
	}
}

func getBoolFunc(valueFunc func() *Value) func() (bool, error) {
	return func() (bool, error) {
		value := valueFunc()
		if value.Err != nil {
			return false, value.Err
		}
		switch value.typ {
		case Number:
			if value.Value.(int64) == 0 {
				return false, nil
			}
			return true, nil
		case String:
			if value.Value.(string) == "" {
				return false, nil
			}
		case Bool:
			if value.Value.(bool) == true {
				return true, nil
			}
			return false, nil
		default:
			return false, ErrNotSupportType
		}
		return false, ErrNotSupportType
	}
}

func (v *DotVisitor) VisitMsgValueType(ctx *parser.MsgValueTypeContext) interface{} {
	return ctx.GetText()
}

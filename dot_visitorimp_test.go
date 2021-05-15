package dot

import (
	"errors"
	"fmt"
	"testing"
)

type Calc struct {
	calcFunc func() *Value
	visitor  *DotVisitor
}

func NewCalc(visitor *DotVisitor, calcFunc func() *Value) *Calc {
	return &Calc{visitor: visitor, calcFunc: calcFunc}
}

func (calc *Calc) Calc(msg *Msg) *Value {
	calc.visitor.Msg = msg
	return calc.calcFunc()
}


func TestCalc_CalcAdd(t *testing.T) {
	calc := NewCalc(Input([]byte("msg.content.field.Number+1")))
	for i := int64(0); i < 100; i++ {
		msg, err := NewMsg([]byte(fmt.Sprintf(`{"field":%d}`, i)))
		if err != nil {
			t.Fatal(err)
		}
		value := calc.Calc(msg)
		if value.Err != nil {
			t.Fatal(value.Err)
		}
		if value.Value != i+1 {
			t.Fatalf("calc error, expect value:%d, actual:%d", i+1, value.Value)
		}
	}
}

func TestCalc_CalcSub(t *testing.T) {
	calc := NewCalc(Input([]byte("msg.content.field.Number-1")))
	for i := int64(0); i < 100; i++ {
		msg, err := NewMsg([]byte(fmt.Sprintf(`{"field":%d}`, i)))
		if err != nil {
			t.Fatal(err)
		}
		value := calc.Calc(msg)
		if value.Err != nil {
			t.Fatal(value.Err)
		}
		if value.Value != i-1 {
			t.Fatalf("calc error, expect value:%d, actual:%d", i-1, value.Value)
		}
	}
}

func TestCalc_CalcMul(t *testing.T) {
	calc := NewCalc(Input([]byte("msg.content.field.Number*2")))
	for i := int64(0); i < 100; i++ {
		msg, err := NewMsg([]byte(fmt.Sprintf(`{"field":%d}`, i)))
		if err != nil {
			t.Fatal(err)
		}
		value := calc.Calc(msg)
		if value.Err != nil {
			t.Fatal(value.Err)
		}
		if value.Value != i*2 {
			t.Fatalf("calc error, expect value:%d, actual:%d", i*2, value.Value)
		}
	}
}

func TestCalc_CalcDiv(t *testing.T) {
	calc := NewCalc(Input([]byte("msg.content.field.Number/2")))
	for i := int64(0); i < 100; i++ {
		msg, err := NewMsg([]byte(fmt.Sprintf(`{"field":%d}`, i)))
		if err != nil {
			t.Fatal(err)
		}
		value := calc.Calc(msg)
		if value.Err != nil {
			t.Fatal(err)
		}
		if value.Value != i/2 {
			t.Fatalf("calc error, expect value:%d, actual:%d", i/2, value.Value)
		}
	}
}

func TestDotVisitor_VisitAndOr01(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number && msg.content.b.Number && msg.content.c.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":1,"b":1,"c":1}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if value.Value.(int64) != 1 {
		t.Fatal(value.Err)
	}
}

func TestDotVisitor_VisitAndOr02(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number && msg.content.b.Number && msg.content.c.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":1,"b":1,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if !errors.Is(value.Err, ErrNotStatement) {
		t.Fatal("compute failure")
	}
}

func TestDotVisitor_VisitAndOr03(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number || msg.content.b.Number || msg.content.c.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":0,"b":0,"c":1}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if value.Value.(int64) != 1 {
		t.Fatal(err)
	}
}

func TestDotVisitor_VisitAndOr04(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number && msg.content.b.Number && msg.content.c.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":0,"b":0,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if !errors.Is(value.Err, ErrNotStatement) {
		t.Fatal("compute failure")
	}
}

func TestDotVisitor_VisitCompare01Gt(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number > msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":1,"b":0,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if value.Err != nil {
		t.Fatal(value.Err)
	}
	if value.Value.(int64) != 1 {
		t.Fatal(value.Err)
	}
}

func TestDotVisitor_VisitCompare01GtFailure(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number > msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":0,"b":0,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if !errors.Is(value.Err, ErrNotStatement) {
		t.Fatal("compute failure")
	}
}

func TestDotVisitor_VisitCompare02Lt(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number < msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":0,"b":1,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if value.Err != nil {
		t.Fatal(err)
	}
	if value.Value.(int64) != 1 {
		t.Fatal(err)
	}
}

func TestDotVisitor_VisitCompare02LtFailure(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number < msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":0,"b":0,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if !errors.Is(value.Err, ErrNotStatement) {
		t.Fatal("compute failure")
	}
}

func TestDotVisitor_VisitCompare03LtEq(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number <= msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":1,"b":1,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if value.Err != nil {
		t.Fatal(err)
	}
	if value.Value.(int64) != 1 {
		t.Fatal(err)
	}
}

func TestDotVisitor_VisitCompare03LtEqFailure(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number <= msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":1,"b":0,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if !errors.Is(value.Err, ErrNotStatement) {
		t.Fatal("compute failure")
	}
}

func TestDotVisitor_VisitCompare04Eq(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number == msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":1,"b":1,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if value.Err != nil {
		t.Fatal(err)
	}
	if value.Value.(int64) != 1 {
		t.Fatal(err)
	}
}

func TestDotVisitor_VisitCompare04EqFailure(t *testing.T) {
	statement := []byte(`
if(msg.content.a.Number == msg.content.b.Number){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":1,"b":0,"c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if !errors.Is(value.Err, ErrNotStatement) {
		t.Fatal("compute failure")
	}
}

func TestDotVisitor_VisitCompare05EqStr(t *testing.T) {
	statement := []byte(`
if(msg.content.a.String == msg.content.b.String){
	1
}
`)
	calc := NewCalc(Input(statement))
	msg, err := NewMsg([]byte(`{"a":"1","b":"1","c":0}`))
	if err != nil {
		t.Fatal(err)
	}
	value := calc.Calc(msg)
	if value.Err != nil {
		t.Fatal(value.Err)
	}
	if value.Value.(int64) != 1 {
		t.Fatal("compute failure")
	}
}

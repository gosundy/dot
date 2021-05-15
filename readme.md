## 动机
使用配置文件可以描述json字段的运算，包括'+','-','*','\','or','and','not' 和 'if',以后可以内置一些函数可以实现流计算中的算子功能
### 例子
IF 返回命中的值
```golang
package main

import (
	"fmt"

	"github.com/gosundy/dot"
)

type Calc struct {
	calcFunc func() *dot.Value
	visitor  *dot.DotVisitor
}

func main() {
	statement := []byte(`
if(msg.content.c.Number && msg.content.a.Number){
	1
}

if (msg.content.b.Number){
	2
}

if (msg.content.a.Number){
    3
}
4
`)
	calc := NewCalc(dot.Input(statement))
	msg, err := dot.NewMsg([]byte(`{"c":0,"a":0,"b":0}`))
	if err != nil {
		panic(err)
	}
	value := calc.Calc(msg)
	if value.Err != nil {
		panic(err)
	}
	fmt.Println(value.Value)
}

func NewCalc(visitor *dot.DotVisitor, calcFunc func() *dot.Value) *Calc {
	return &Calc{visitor: visitor, calcFunc: calcFunc}
}

func (calc *Calc) Calc(msg *dot.Msg) *dot.Value {
	calc.visitor.Msg = msg
	return calc.calcFunc()
}
```
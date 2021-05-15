package dot

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkCalc_Calc(b *testing.B) {
	b.ReportAllocs()
	calc := NewCalc(Input([]byte("msg.content.field.Number+1\n")))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		msg, err := NewMsg([]byte(fmt.Sprintf(`{"field":%d}`, i)))
		if err != nil {
			b.Fatal(err)
		}
		value := calc.Calc(msg)
		if value.Err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCalc_Calc02(b *testing.B) {
	a := 0
	b.Run("test", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			time.Sleep(time.Millisecond)
			a = a + 1
		}

	})
	fmt.Println(a)
}

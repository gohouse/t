package t

import (
	"fmt"
	"testing"
)

func TestNewT(t *testing.T) {
	var b T

	//b = New("abc")
	//b = New(1.3)
	b = New("2.3")
	//b = New(23)
	//b = New(true)
	//b = New(false)
	//b = New(nil)

	fmt.Println(b.String())
	fmt.Println(b.Float64())
	fmt.Println(b.Float32())
	fmt.Println(b.Int64())
	fmt.Println(b.Int())
	fmt.Println(b.Int32())
	fmt.Println(b.Int16())
	fmt.Println(b.Int8())
	fmt.Println(b.Uint64())
	fmt.Println(b.Uint())
	fmt.Println(b.Uint32())
	fmt.Println(b.Uint16())
	fmt.Println(b.Uint8())
	fmt.Println(b.Bool())
}

func BenchmarkType_Int64(b *testing.B) {
	var a T = New("2.3")
	for i:=0;i<b.N;i++ {
		a.Int64()
	}
}

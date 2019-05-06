package t

import (
	"fmt"
	"testing"
)

func TestNewT(t *testing.T) {
	var b T

	b = New("abc")
	//b = New(1.3)
	//b = New("2.3")
	//b = New(23)
	//b = New(true)
	//b = New(false)
	//b = New(nil)
	
	fmt.Println(b.String())
	fmt.Println(b.Float64())
	fmt.Println(b.Float32())
	fmt.Println(b.Int64())
	fmt.Println(b.Int())
	fmt.Println(b.Uint64())
	fmt.Println(b.Uint())
	fmt.Println(b.Bool())
}
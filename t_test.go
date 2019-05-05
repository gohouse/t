package t

import (
	"fmt"
	"testing"
)

func TestNewT(t *testing.T) {
	var b T = New(4.9)
	fmt.Println(b==nil)
	fmt.Println(b.Int64())
	fmt.Println(b.Int())
	fmt.Println(b.Uint64())
	fmt.Println(b.Uint())
	fmt.Println(b.Float64())
	fmt.Println(b.Float32())
	fmt.Println(b.Bool())
}
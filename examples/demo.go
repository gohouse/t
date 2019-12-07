package main

import (
	"fmt"
	"github.com/gohouse/t"
)

func main() {
	var mobile float64 = 12341234

	res:=t.New(mobile).String()
	fmt.Println(res)
}

package main

import (
	"fmt"
	"github.com/gohouse/t"
)

func main() {
	convertComplex()
}

func convertComplex() {
	var cc interface{}
	cc = map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("%#v \n", t.New(cc).Map())
	fmt.Printf("%#v \n", t.New(cc).MapInterface())
	fmt.Printf("%#v \n", t.New(cc).MapString())
	fmt.Printf("%#v \n", t.New(cc).MapStringInterface())
}

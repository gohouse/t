package main

import (
	"fmt"
	"github.com/gohouse/t"
)

func main() {
	convertComplex()
	jsonExtract()
}

func convertComplex() {
	var cc interface{}
	cc = map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("%#v \n", t.New(cc).Map())
	fmt.Printf("%#v \n", t.New(cc).MapInterfaceT())
	fmt.Printf("%#v \n", t.New(cc).MapStringT())
	fmt.Printf("%#v \n", t.New(cc).MapStringInterface())
}

func jsonExtract() {
	var js = `{"a":1,"b":["c",{"d":2}]}`
	//var js2 = map[string]string{
	//	"a":"aa",
	//	"b":"bb",
	//}
	j := t.New(js)
	fmt.Println("extract test:", j.Extract("a"), j.Extract("b.0"), j.Extract("b.1.d"))
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gohouse/t"
)

func main() {

	fmt.Println(t.New("2006-01-02 15:04:05").Length())

	var data = bytes.NewBufferString("abcdefg")
	var aa = make([]byte, 1024)
	n,err := bufio.NewReader(data).Read(aa)
	fmt.Println(string(aa),n,err)
	var a t.Type = t.New(222)
	fmt.Printf("%#v \n",a.Slice())
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

func jsonExtract()  {
	var js = `{"a":1,"b":["c",{"d":2}]}`
	//var js2 = map[string]string{
	//	"a":"aa",
	//	"b":"bb",
	//}
	j := t.New(js)
	fmt.Println("extract test:",j.Extract("a"), j.Extract("b.0"), j.Extract("b.1.d"))
}
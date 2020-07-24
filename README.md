## t
golang type extension. Conversion of Basic Types.  
(golang基本数据类型的相互转换)

## catalog index
- [1. type bind convertion (类型定义转换)](#1-type-bind-convertion-类型定义转换)
- [2. standard type convertion (标准类型转换)](#2-standard-type-convertion-标准类型转换)  

## usage

### 1. type bind convertion (类型定义转换)
```go
package main

import (
	"fmt"
	"github.com/gohouse/golib/t"
)

func main()  {
    var b t.Type
    b = t.New("2.3")

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
```
result
```sh
2.3
2.3
2.3
2
2
2
2
2
2
2
2
2
2
true
```

### 2. standard type convertion (标准类型转换)
```go
package main

import (
	"fmt"
	"github.com/gohouse/golib/t"
)

func main()  {
	var a = "8.8"
	fmt.Println(t.ParseString(a))
	fmt.Println(t.ParseFloat64(a))
	fmt.Println(t.ParseInt64(a))
	fmt.Println(t.ParseInt(a))
	fmt.Println(t.ParseUint64(a))
	fmt.Println(t.ParseUint(a))
	fmt.Println(t.ParseBool(a))
}
```
result
```bash
8.8
8.8
8
8
8
8
true
```

### 3. complex type converts (复杂类型转换)
```go
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
```
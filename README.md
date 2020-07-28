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

### 4. 验证规则判断
```go
t.New("http://www.gorose.com").IsUrl()
t.New("qq@gorose.com").IsEmail()
```
所有判断方法
```go
// iDetermine 判断
type iDetermine interface {
	IsNumeric() bool                        // 是否数字
	IsInteger() bool                        // 是否整数
	IsFloat() bool                          // 是否浮点数
	IsZero() bool                           // 是否零值
	IsChineseCharacters() bool              // 是否汉字
	IsChineseName() bool                    // 是否中文名字
	IsHost() bool                           // 是否域名
	IsUrl() bool                            // 是否互联网url地址
	IsEmail() bool                          // 是否邮箱地址
	IsChineseMobile() bool                  // 是否中国手机号
	IsDate() bool                           // 是否常用的日期格式
	IsDatetime() bool                       // 是否常用的日期时间格式
	IsIpV4() bool                           // 是否ipv4地址
	IsIpV6() bool                           // 是否ipv6地址
	IsIp() bool                             // 是否ip地址
	IsChineseID() bool                      // 是否中国大陆身份证号码
	IsXml() bool                            // 是否xml
	IsJson() bool                           // 是否json
	IsJsonMap() bool                        // 是否是json对象
	IsJsonSlice() bool                      // 是否是json数组
	IsBetween(min, max interface{}) bool    // 是否在两数之间
	IsBetweenFloat64(min, max float64) bool // 是否在两个浮点数之间
	IsBetweenAlpha(min, max string) bool    // 是否在两个字符之间
}
```
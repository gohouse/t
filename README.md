# t
golang type extention

## usage
```go
package main

import (
	"fmt"
	"github.com/gohouse/t"
)

func main()  {
	var b t.T 
	fmt.Println(b==nil)
	
	b = t.New(4.9)
	
	fmt.Println(b.Int64())
	fmt.Println(b.Int())
	fmt.Println(b.Uint64())
	fmt.Println(b.Uint())
	fmt.Println(b.Float64())
	fmt.Println(b.Float32())
	fmt.Println(b.Bool())
}
```
result
```sh
false
1
1
1
1
1
1
true
```
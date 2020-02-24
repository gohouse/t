package t

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewT(t *testing.T) {
	var b T

	//b = New("abc")
	//b = New(1.3)
	//b = New("2.3")
	//b = New(23)
	//b = New(true)
	//b = New(false)
	//b = New(nil)
	//b = New(struct {A string}{"bcd"})
	//b = New(New(3))
	b = New(map[T]T{New("a"): New(3)})
	t.Log(b.MapInt64()[0])

	t.Log(b.String())
	t.Log(b.Float64())
	t.Log(b.Float32())
	t.Log(b.Int64())
	t.Log(b.Int())
	t.Log(b.Int32())
	t.Log(b.Int16())
	t.Log(b.Int8())
	t.Log(b.Uint64())
	t.Log(b.Uint())
	t.Log(b.Uint32())
	t.Log(b.Uint16())
	t.Log(b.Uint8())
	t.Log(b.Bool())
}

func TestType_MapStr(t *testing.T) {
	var m = New(map[string]interface{}{"a": 2, "b": "c", "d": "d"})
	a := m.MapString()
	t.Log(a)
	t.Log(a["a"])
}

func TestType_Map(t *testing.T) {
	var m = New(map[interface{}]interface{}{"a": 2, 2: "c", 3.3: "d", true: true, false: 3})
	a := m.MapInterface()
	t.Log(a)
	t.Log(a["a"])
	t.Log(a[2].String())
	t.Log(a[3.3])
	t.Log(a[true].Bool())
	t.Log(a[false])
}

func TestType_Map2(t *testing.T) {
	var m = New(`{"a": 2, "b":3,"33":{"331":"d"}}`)
	a := m.Map()
	t.Log(a[New("a")])
	t.Log(m.Extract("33.331"))
}

func TestType_Slice(t *testing.T) {
	var a = New([]string{"a", "b"})

	for _, v := range a.Slice() {
		t.Log(v.String(), v.Bool())
	}
}

func TestType_Bind(t *testing.T) {
	type json struct {
		A interface{} `json:"a"`
		B string `json:"b"`
	}
	var a = map[string]interface{}{"a":1,"b":"bbb"}
	var js json
	New(a).Bind(&js)
	t.Logf("%+v",js)
}

func BenchmarkType_Int64(b *testing.B) {
	var a T = New("2.3")
	for i := 0; i < b.N; i++ {
		a.Int64()
	}
}

func TestType_Read(t *testing.T) {
	res := New("abc")
	//res := New(strings.NewReader("abc"))
	b,_:=ioutil.ReadAll(res)
	fmt.Printf("%s\n",b)
}
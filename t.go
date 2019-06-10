package t

import (
	"reflect"
)

type T interface {
	String() string
	Float64() float64
	Float32() float32
	Int64() int64
	Int() int
	Int32() int32
	Int16() int16
	Int8() int8
	Uint64() uint64
	Uint() uint
	Uint32() uint32
	Uint16() uint16
	Uint8() uint8
	Bool() bool
	Slice() []T
	Map() map[interface{}]T
	MapStr() map[string]T
	MapInt() map[int64]T
}

type Type struct {
	val interface{}
}

var _ T = &Type{}

func New(o interface{}) T {
	return &Type{val: o}
}

func (t Type) String() string {
	return ParseString(t.val)
}

func (t Type) Float64() float64 {
	return ParseFloat64(t.val)
}
func (t Type) Float32() float32 {
	return ParseFloat32(t.val)
}

func (t Type) Int64() int64 {
	return ParseInt64(t.val)
}
func (t Type) Int() int {
	return ParseInt(t.val)
}
func (t Type) Int32() int32 {
	return int32(ParseInt64(t.val))
}
func (t Type) Int16() int16 {
	return int16(ParseInt64(t.val))
}
func (t Type) Int8() int8 {
	return int8(ParseInt64(t.val))
}

func (t Type) Uint64() uint64 {
	return ParseUint64(t.val)
}
func (t Type) Uint() uint {
	return ParseUint(t.val)
}
func (t Type) Uint32() uint32 {
	return uint32(ParseUint(t.val))
}
func (t Type) Uint16() uint16 {
	return uint16(ParseUint(t.val))
}
func (t Type) Uint8() uint8 {
	return uint8(ParseUint(t.val))
}

func (t Type) Bool() bool {
	return ParseBool(t.val)
}

func (t Type) Slice() []T {
	ref := reflect.ValueOf(t.val)
	l:=ref.Len()
	v:=ref.Slice(0,l)
	var res = []T{}
	for i:=0;i<l;i++{
		res = append(res, New(v.Index(i).Interface()))
	}
	return res
}

func (t Type) Map() map[interface{}]T {
	ref := reflect.ValueOf(t.val)
	var res = make(map[interface{}]T)
	keys := ref.MapKeys()
	for _,item := range keys {
		res[item.Interface()] = New(ref.MapIndex(item).Interface())
	}
	return res
}

func (t Type) MapStr() map[string]T {
	ref := reflect.ValueOf(t.val)
	var res = make(map[string]T)
	keys := ref.MapKeys()
	for _,item := range keys {
		res[item.String()] = New(ref.MapIndex(item).Interface())
	}
	return res
}

func (t Type) MapInt() map[int64]T {
	ref := reflect.ValueOf(t.val)
	var res = make(map[int64]T)
	keys := ref.MapKeys()
	for _,item := range keys {
		res[item.Int()] = New(ref.MapIndex(item).Interface())
	}
	return res
}

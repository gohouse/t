package t

import (
	"reflect"
)

type Map map[T]T
type MapInterface map[interface{}]T
type MapString map[string]T
type MapInt64 map[int64]T
type Slice []T
type SliceInterface []interface{}
type SliceString []string
type SliceInt64 []int64

type T interface {
	Interface() interface{}
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
	Slice() Slice
	SliceInterface() []interface{}
	SliceString() []string
	SliceInt64() []int64
	Map() Map
	MapInterface() MapInterface
	MapString() MapString
	MapInt64() MapInt64
}

type Type struct {
	val interface{}
}

var _ T = &Type{}

func New(o interface{}) Type {
	switch o.(type) {
	case T:
		return Type{o.(T).Interface()}
	default:
		return Type{val: o}
	}
}

func (t Type) Interface() interface{} {
	return t.val
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

func (t Type) Slice() Slice {
	ref := reflect.ValueOf(t.val)
	l := ref.Len()
	v := ref.Slice(0, l)
	var res = Slice{}
	for i := 0; i < l; i++ {
		res = append(res, New(v.Index(i).Interface()))
	}
	return res
}

func (t Type) SliceInterface() []interface{} {
	s := t.Slice()
	var res = []interface{}{}
	for _, item := range s {
		res = append(res, item.Interface())
	}
	return res
}

func (t Type) SliceString() []string {
	s := t.Slice()
	var res = []string{}
	for _, item := range s {
		res = append(res, item.String())
	}
	return res
}

func (t Type) SliceInt64() []int64 {
	s := t.Slice()
	var res = []int64{}
	for _, item := range s {
		res = append(res, item.Int64())
	}
	return res
}

func (t Type) Map() Map {
	ref := reflect.ValueOf(t.val)
	var res = make(Map)
	keys := ref.MapKeys()
	for _, item := range keys {
		res[New(item.Interface())] = New(ref.MapIndex(item).Interface())
	}
	return res
}

func (t Type) MapInterface() MapInterface {
	m := t.Map()
	var res = make(MapInterface)
	for k, v := range m {
		res[k.Interface()] = v
	}
	return res
}

func (t Type) MapString() MapString {
	m := t.Map()
	var res = make(MapString)
	for k, v := range m {
		res[k.String()] = v
	}
	return res
}

func (t Type) MapInt64() MapInt64 {
	m := t.Map()
	var res = make(MapInt64)
	for k, v := range m {
		res[k.Int64()] = v
	}
	return res
}

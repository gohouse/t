package t

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Map ...
type Map map[T]T

// MapInterface ...
type MapInterface map[interface{}]T

// MapString ...
type MapString map[string]T

// MapInt64 ...
type MapInt64 map[int64]T

//// MapStringInterface ...
//type MapStringInterface map[string]interface{}

// Slice ...
type Slice []T

//// SliceInterface ...
//type SliceInterface []interface{}

// SliceString ...
type SliceString []string

// SliceInt64 ...
type SliceInt64 []int64

// T ...
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
	Byte() byte
	Bytes() []byte
	Slice() Slice
	SliceInterface() []interface{}
	SliceMapStringInterface() []map[string]interface{}
	SliceString() []string
	SliceInt64() []int64
	Map() Map
	MapInterface() MapInterface
	MapString() MapString
	MapInt64() MapInt64
	MapStringInterface() map[string]interface{}
	SliceMapString() []MapString
	BindJson(o interface{}) error
	Extract(key string, defaultVal ...interface{}) interface{}
}

// Type ...
type Type struct {
	val interface{}
}

var _ T = &Type{}

// New ...
func New(o interface{}) Type {
	switch o.(type) {
	case T:
		return Type{o.(T).Interface()}
	default:
		return Type{val: o}
	}
}

// Interface ...
func (t Type) Interface() interface{} {
	return t.val
}

// String ...
func (t Type) String() string {
	return ParseString(t.val)
}

// Float64 ...
func (t Type) Float64() float64 {
	return ParseFloat64(t.val)
}

// Float32 ...
func (t Type) Float32() float32 {
	return ParseFloat32(t.val)
}

// Int64 ...
func (t Type) Int64() int64 {
	return ParseInt64(t.val)
}

// Int ...
func (t Type) Int() int {
	return ParseInt(t.val)
}

// Int32 ...
func (t Type) Int32() int32 {
	return int32(ParseInt64(t.val))
}

// Int16 ...
func (t Type) Int16() int16 {
	return int16(ParseInt64(t.val))
}

// Int8 ...
func (t Type) Int8() int8 {
	return int8(ParseInt64(t.val))
}

// Uint64 ...
func (t Type) Uint64() uint64 {
	return ParseUint64(t.val)
}

// Uint ...
func (t Type) Uint() uint {
	return ParseUint(t.val)
}

// Uint32 ...
func (t Type) Uint32() uint32 {
	return uint32(ParseUint(t.val))
}

// Uint16 ...
func (t Type) Uint16() uint16 {
	return uint16(ParseUint(t.val))
}

// Uint8 ...
func (t Type) Uint8() uint8 {
	return uint8(ParseUint(t.val))
}

// Bool ...
func (t Type) Bool() bool {
	return ParseBool(t.val)
}

// Byte ...
func (t Type) Byte() byte {
	return ParseByte(t.val)
}

// Bytes ...
func (t Type) Bytes() []byte {
	return ParseBytes(t.val)
}

// Slice ...
func (t Type) Slice() Slice {
	ref := reflect.Indirect(reflect.ValueOf(t.val))
	l := ref.Len()
	v := ref.Slice(0, l)
	var res = Slice{}
	for i := 0; i < l; i++ {
		res = append(res, New(v.Index(i).Interface()))
	}
	return res
}

// SliceInterface ...
func (t Type) SliceInterface() []interface{} {
	s := t.Slice()
	var res = []interface{}{}
	for _, item := range s {
		res = append(res, item.Interface())
	}
	return res
}

// SliceMapStringInterface ...
func (t Type) SliceMapStringInterface() []map[string]interface{} {
	s := t.Slice()
	var res = []map[string]interface{}{}
	for _, item := range s {
		res = append(res, item.MapStringInterface())
	}
	return res
}

// SliceString ...
func (t Type) SliceString() []string {
	s := t.Slice()
	var res = []string{}
	for _, item := range s {
		res = append(res, item.String())
	}
	return res
}

// SliceInt64 ...
func (t Type) SliceInt64() []int64 {
	s := t.Slice()
	var res = []int64{}
	for _, item := range s {
		res = append(res, item.Int64())
	}
	return res
}

// Map ...
func (t Type) Map() Map {
	ref := reflect.Indirect(reflect.ValueOf(t.val))
	var res = Map{}
	switch ref.Kind() {
	case reflect.Map:
		//var res = make(Map)
		keys := ref.MapKeys()
		for _, item := range keys {
			res[New(item.Interface())] = New(ref.MapIndex(item).Interface())
		}
	default:
		var res2 = map[string]interface{}{}
		err := t.BindJson(&res2)
		if err != nil {
			return Map{}
		}
		for k, v := range res2 {
			res[New(k)] = New(v)
		}
	}
	return res
}

// MapInterface ...
func (t Type) MapInterface() MapInterface {
	m := t.Map()
	var res = make(MapInterface)
	for k, v := range m {
		res[k.Interface()] = v
	}
	return res
}

// MapString ...
func (t Type) MapString() MapString {
	m := t.Map()
	var res = make(MapString)
	for k, v := range m {
		res[k.String()] = v
	}
	return res
}

// MapStringInterface ...
func (t Type) MapStringInterface() map[string]interface{} {
	m := t.Map()
	var res = make(map[string]interface{})
	for k, v := range m {
		res[k.String()] = v.Interface()
	}
	return res
}

// MapInt64 ...
func (t Type) MapInt64() MapInt64 {
	m := t.Map()
	var res = make(MapInt64)
	for k, v := range m {
		res[k.Int64()] = v
	}
	return res
}

// SliceMapString ...
func (t Type) SliceMapString() []MapString {
	m := t.Slice()
	var result []MapString
	for _, item := range m {
		result = append(result, item.MapString())
	}
	return result
}

// BindJson 将绑定结果当做json,来绑定到对象上
func (t Type) BindJson(o interface{}) error {
	return json.Unmarshal(t.Bytes(), o)
}
// Extract 多层次抽取值
// 例: var m = New(`{"a": 2, "b":3,"33":{"cc":"d"}}`); println(m.Extract("33.cc")) // d
func (t Type) Extract(key string, defaultVal ...interface{}) interface{} {
	if key == "" {
		return nil
	}
	var split []string
	// 如果key包含了点,则为多级调用
	if strings.Contains(key, ".") {
		split = strings.Split(key, ".")
	} else { // 如果key不包含点, 则就是直接一级调用
		split = []string{key}
	}

	// 取指定语言的配置
	var currentMap = t.Map()
	var currentVal interface{}
	for _, item := range split {
		if v, ok := currentMap[New(item)]; ok {
			currentMap = New(v).Map()
			currentVal = v
		} else {
			break
		}
	}

	// 如果没有取到且传入了默认值, 则返回默认值
	if currentVal == nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return currentVal
}

package t

import "reflect"

type iSlice interface {
	Slice() Slice
	SliceString() []string
	SliceFloat64() []float64
	SliceInt64() []int64
	SliceInt() []int
	SliceInterface() []interface{}
	SliceMapString() []MapString
	SliceMapStringInterface() []map[string]interface{}
}

// Slice ...
type Slice []T

// Slice ...
func (t Type) Slice() Slice {
	if t.val == nil {
		return Slice{}
	}
	ref := reflect.Indirect(reflect.ValueOf(t.val))
	var res = Slice{}
	switch ref.Kind() {
	case reflect.Slice:
		l := ref.Len()
		v := ref.Slice(0, l)
		for i := 0; i < l; i++ {
			res = append(res, New(v.Index(i).Interface()))
		}
	default:
		res = append(res, New(ref.Interface()))
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

// SliceFloat64 ...
func (t Type) SliceFloat64() []float64 {
	s := t.Slice()
	var res = []float64{}
	for _, item := range s {
		res = append(res, item.Float64())
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

// SliceInt ...
func (t Type) SliceInt() []int {
	s := t.Slice()
	var res = []int{}
	for _, item := range s {
		res = append(res, item.Int())
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

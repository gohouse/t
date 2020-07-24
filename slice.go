package t

import (
	"reflect"
)

type iSlice interface {
	Slice() Slice
	SliceString() []string
	SliceFloat64() []float64
	SliceInt64() []int64
	SliceInt() []int
	SliceInterface() []interface{}
	SliceMapStringT() []MapStringT
	SliceMapStringInterface() []map[string]interface{}
}

// Slice ...
type Slice []Type

// Slice ...
func (tc TypeContext) Slice() Slice {
	if tc.val == nil {
		return Slice{}
	}
	ref := reflect.Indirect(reflect.ValueOf(tc.val))
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
func (tc TypeContext) SliceInterface() []interface{} {
	s := tc.Slice()
	var res = []interface{}{}
	for _, item := range s {
		res = append(res, item.Interface())
	}
	return res
}

// SliceMapStringInterface ...
func (tc TypeContext) SliceMapStringInterface() []map[string]interface{} {
	s := tc.Slice()
	var res = []map[string]interface{}{}
	for _, item := range s {
		res = append(res, item.MapStringInterface())
	}
	return res
}

// SliceString ...
func (tc TypeContext) SliceString() []string {
	s := tc.Slice()
	var res = []string{}
	for _, item := range s {
		res = append(res, item.String())
	}
	return res
}

// SliceFloat64 ...
func (tc TypeContext) SliceFloat64() []float64 {
	s := tc.Slice()
	var res = []float64{}
	for _, item := range s {
		res = append(res, item.Float64())
	}
	return res
}

// SliceInt64 ...
func (tc TypeContext) SliceInt64() []int64 {
	s := tc.Slice()
	var res = []int64{}
	for _, item := range s {
		res = append(res, item.Int64())
	}
	return res
}

// SliceInt ...
func (tc TypeContext) SliceInt() []int {
	s := tc.Slice()
	var res = []int{}
	for _, item := range s {
		res = append(res, item.Int())
	}
	return res
}

// SliceMapStringT ...
func (tc TypeContext) SliceMapStringT() []MapStringT {
	m := tc.Slice()
	var result []MapStringT
	for _, item := range m {
		result = append(result, item.MapStringT())
	}
	return result
}

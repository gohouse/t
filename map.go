package t

import (
	"reflect"
)

type iMap interface {
	Map() Map
	MapInterfaceT() MapInterfaceT
	MapStringT() MapStringT
	MapIntT64T() MapIntT64T
	MapIntT() MapIntT
	MapStringInterface() map[string]interface{}
}
// Map ...
type Map map[Type]Type

// MapInterfaceT ...
type MapInterfaceT map[interface{}]Type

// MapStringT ...
type MapStringT map[string]Type

// MapIntT64T ...
type MapIntT64T map[int64]Type
type MapIntT map[int]Type

//// MapStringInterface ...
//type MapStringInterface map[string]interface{}


// Map ...
func (tc TypeContext) Map() Map {
	if tc.val==nil {
		return Map{}
	}
	ref := reflect.Indirect(reflect.ValueOf(tc.val))
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
		err := tc.Bind(&res2)
		if err != nil {
			return Map{}
		}
		for k, v := range res2 {
			res[New(k)] = New(v)
		}
	}
	return res
}

// MapInterfaceT ...
func (tc TypeContext) MapInterfaceT() MapInterfaceT {
	m := tc.Map()
	var res = make(MapInterfaceT)
	for k, v := range m {
		res[k.Interface()] = v
	}
	return res
}

// MapStringT ...
func (tc TypeContext) MapStringT() MapStringT {
	m := tc.Map()
	var res = make(MapStringT)
	for k, v := range m {
		res[k.String()] = v
	}
	return res
}

// MapStringInterface ...
func (tc TypeContext) MapStringInterface() map[string]interface{} {
	m := tc.Map()
	var res = make(map[string]interface{})
	for k, v := range m {
		res[k.String()] = v.Interface()
	}
	return res
}

// MapIntT64T ...
func (tc TypeContext) MapIntT64T() MapIntT64T {
	m := tc.Map()
	var res = make(MapIntT64T)
	for k, v := range m {
		res[k.Int64()] = v
	}
	return res
}

// MapIntT64T ...
func (tc TypeContext) MapIntT() MapIntT {
	m := tc.Map()
	var res = make(MapIntT)
	for k, v := range m {
		res[k.Int()] = v
	}
	return res
}


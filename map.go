package t

import "reflect"

type iMap interface {
	Map() Map
	MapInterface() MapInterface
	MapString() MapString
	MapInt64() MapInt64
	MapInt() MapInt
	MapStringInterface() map[string]interface{}
}
// Map ...
type Map map[T]T

// MapInterface ...
type MapInterface map[interface{}]T

// MapString ...
type MapString map[string]T

// MapInt64 ...
type MapInt64 map[int64]T
type MapInt map[int]T

//// MapStringInterface ...
//type MapStringInterface map[string]interface{}


// Map ...
func (t Type) Map() Map {
	if t.val==nil {
		return Map{}
	}
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
		err := t.Bind(&res2)
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

// MapInt64 ...
func (t Type) MapInt() MapInt {
	m := t.Map()
	var res = make(MapInt)
	for k, v := range m {
		res[k.Int()] = v
	}
	return res
}


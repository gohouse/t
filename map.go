package t

import "reflect"

type iMap interface {
	Map() Map
	MapInterfaceT() MapInterfaceT
	MapStringT() MapStringT
	MapIntT64T() MapIntT64T
	MapIntT() MapIntT
	MapStringInterface() map[string]interface{}
}
// Map ...
type Map map[T]T

// MapInterfaceT ...
type MapInterfaceT map[interface{}]T

// MapStringT ...
type MapStringT map[string]T

// MapIntT64T ...
type MapIntT64T map[int64]T
type MapIntT map[int]T

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

// MapInterfaceT ...
func (t Type) MapInterfaceT() MapInterfaceT {
	m := t.Map()
	var res = make(MapInterfaceT)
	for k, v := range m {
		res[k.Interface()] = v
	}
	return res
}

// MapStringT ...
func (t Type) MapStringT() MapStringT {
	m := t.Map()
	var res = make(MapStringT)
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

// MapIntT64T ...
func (t Type) MapIntT64T() MapIntT64T {
	m := t.Map()
	var res = make(MapIntT64T)
	for k, v := range m {
		res[k.Int64()] = v
	}
	return res
}

// MapIntT64T ...
func (t Type) MapIntT() MapIntT {
	m := t.Map()
	var res = make(MapIntT)
	for k, v := range m {
		res[k.Int()] = v
	}
	return res
}


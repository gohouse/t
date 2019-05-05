package t

import (
	"reflect"
	"strconv"
	"strings"
)

//type inType int
//const(
//	integer inType = iota
//	uinteger
//	floater
//	stringer
//	booler
//)
//func typeParser(o interface{}) interface{} {
//	switch o.(type) {
//	case int,int8,int16,int32,int64:
//		return integer
//	case uint,uint8,uint16,uint32,uint64,uintptr:
//		return uinteger
//	case string:
//		return stringer
//	case float32,float64:
//		return floater
//	case bool:
//		return booler
//	default:
//		return nil
//	}
//}

func ParseString(o interface{}) string {
	switch o.(type) {
	case int,int8,int16,int32,int64:
		return strconv.Itoa(int(reflect.ValueOf(o).Int()))
	case uint,uint8,uint16,uint32,uint64,uintptr:
		return strconv.Itoa(int(reflect.ValueOf(o).Uint()))
	case float32,float64:
		return strconv.Itoa(int(reflect.ValueOf(o).Float()))
	case string:
		return o.(string)
	case bool:
		return strconv.FormatBool(o.(bool))
	}
	return ""
}

func ParseInt64(o interface{}) (res int64) {
	res, _ = strconv.ParseInt(ParseString(o), 10, 64)
	return
}

func ParseInt(o interface{}) int {
	return int(ParseInt64(o))
}

func ParseUint64(o interface{}) uint64 {
	return uint64(ParseInt64(o))
}

func ParseUint(o interface{}) uint {
	return uint(ParseInt64(o))
}

func ParseFloat64(o interface{}) (res float64) {
	res, _ = strconv.ParseFloat(ParseString(o), 64)
	return
}

func ParseFloat32(o interface{}) float32 {
	return float32(ParseFloat64(o))
}

func ParseBool(o interface{}) bool {
	var str = ParseString(o)
	//res,_:=strconv.ParseBool(str)
	str = strings.ToUpper(str)
	switch str {
	//case "", "0", "f", "F", "false", "FALSE", "False":
	case "", "0", "F", "FALSE":
		return false
	default:
		return true
	}
}

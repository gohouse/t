package t

import (
	"fmt"
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
	return fmt.Sprint(o)
	//switch o.(type) {
	//case int,int8,int16,int32,int64:
	//	//return strconv.Itoa(int(reflect.ValueOf(o).Int()))
	//	return strconv.FormatInt(reflect.ValueOf(o).Int(), 10)
	//case uint,uint8,uint16,uint32,uint64,uintptr:
	//	//return strconv.Itoa(int(reflect.ValueOf(o).Uint()))
	//	return strconv.FormatUint(reflect.ValueOf(o).Uint(), 10)
	//case float32,float64:
	//	//return strconv.FormatFloat(reflect.ValueOf(o).Float(),'f',6 , 64)
	//	return fmt.Sprint(o)
	//case string:
	//	return o.(string)
	//case bool:
	//	return strconv.FormatBool(o.(bool))
	//}
	//return ""
}

func ParseFloat64(o interface{}) (res float64) {
	var str = ParseString(o)
	var strUpper = strings.ToUpper(str)
	if strUpper=="TRUE" || strUpper=="T" {
		str = "1"
	} else if strUpper=="FALSE" || strUpper=="F" {
		str = "0"
	}
	res, _ = strconv.ParseFloat(str, 64)
	return
}

func ParseFloat32(o interface{}) float32 {
	return float32(ParseFloat64(o))
}

func ParseInt64(o interface{}) (res int64) {
	res = int64(ParseFloat64(o))
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

func ParseBool(o interface{}) bool {
	var str = ParseString(o)
	//res,_:=strconv.ParseBool(str)
	str = strings.ToUpper(str)
	switch str {
	//case "", "0", "f", "F", "false", "FALSE", "False":
	case "", "0", "F", "FALSE", "<NIL>":
		return false
	default:
		return true
	}
}

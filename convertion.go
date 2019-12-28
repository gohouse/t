package t

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseString(o interface{}) string {
	switch o.(type) {
	case float64:
		return fmt.Sprintf("%v",int64(o.(float64)))
	case float32:
		return fmt.Sprintf("%v",int64(o.(float32)))
	case []byte:
		return string(o.([]byte))
	}
	return fmt.Sprint(o)
}

func ParseFloat64(o interface{}) (res float64) {
	var str = ParseString(o)
	var strUpper = strings.ToUpper(str)
	if strUpper == "TRUE" || strUpper == "T" {
		str = "1"
	} else if strUpper == "FALSE" || strUpper == "F" {
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

func ParseInt32(o interface{}) (res int32) {
	res = int32(ParseFloat64(o))
	return
}

func ParseInt16(o interface{}) (res int16) {
	res = int16(ParseFloat64(o))
	return
}

func ParseInt8(o interface{}) (res int8) {
	res = int8(ParseFloat64(o))
	return
}

func ParseInt(o interface{}) int {
	return int(ParseInt64(o))
}

func ParseUint64(o interface{}) uint64 {
	return uint64(ParseInt64(o))
}

func ParseUint32(o interface{}) uint32 {
	return uint32(ParseInt64(o))
}

func ParseUint16(o interface{}) uint16 {
	return uint16(ParseInt64(o))
}

func ParseUint8(o interface{}) uint8 {
	return uint8(ParseInt64(o))
}

func ParseUint(o interface{}) uint {
	return uint(ParseInt64(o))
}

func ParseBool(o interface{}) bool {
	var str = ParseString(o)
	str = strings.ToUpper(str)
	switch str {
	case "", "0", "F", "FALSE", "<NIL>":
		return false
	default:
		return true
	}
}

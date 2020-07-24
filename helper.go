package t

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ParseFloat64 ...
func ParseFloat64(o interface{}) (res float64) {
	switch value := o.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	default:
		var str = ParseString(o)
		var strUpper = strings.ToUpper(str)
		if strUpper == "TRUE" {
			str = "1"
		} else if strUpper == "FALSE" {
			str = "0"
		}
		res, _ = strconv.ParseFloat(str, 64)
	}
	return
}

// ParseFloat32 ...
func ParseFloat32(o interface{}) float32 {
	return float32(ParseFloat64(o))
}

// ParseInt64 ...
func ParseInt64(o interface{}) (res int64) {
	switch val := o.(type) {
	case int64:
		return val
	case int:
		return int64(val)
	case int32:
		return int64(val)
	case int16:
		return int64(val)
	case int8:
		return int64(val)
	case uint64:
		return int64(val)
	case uint:
		return int64(val)
	case uint32:
		return int64(val)
	case uint16:
		return int64(val)
	case uint8:
		return int64(val)
	default:
		return int64(ParseFloat64(o))
	}
}

// ParseInt32 ...
func ParseInt32(o interface{}) (res int32) {
	res = int32(ParseFloat64(o))
	return
}

// ParseInt16 ...
func ParseInt16(o interface{}) (res int16) {
	res = int16(ParseFloat64(o))
	return
}

// ParseInt8 ...
func ParseInt8(o interface{}) (res int8) {
	res = int8(ParseFloat64(o))
	return
}

// ParseInt ...
func ParseInt(o interface{}) int {
	return int(ParseInt64(o))
}

// ParseUint64 ...
func ParseUint64(o interface{}) uint64 {
	return uint64(ParseInt64(o))
}

// ParseUint32 ...
func ParseUint32(o interface{}) uint32 {
	return uint32(ParseInt64(o))
}

// ParseUint16 ...
func ParseUint16(o interface{}) uint16 {
	return uint16(ParseInt64(o))
}

// ParseUint8 ...
func ParseUint8(o interface{}) uint8 {
	return uint8(ParseInt64(o))
}

// ParseUint ...
func ParseUint(o interface{}) uint {
	return uint(ParseInt64(o))
}

// ParseBool ...
func ParseBool(o interface{}) bool {
	var str = ParseString(o)
	str = strings.ToUpper(str)
	switch str {
	case "", "0", "FALSE", "<NIL>":
		return false
	default:
		return true
	}
}

// ParseByte ...
func ParseByte(i interface{}) byte {
	if v, ok := i.(byte); ok {
		return v
	}
	return ParseUint8(i)
}

// ParseBytes ...
func ParseBytes(o interface{}) []byte {
	if o == nil {
		return nil
	}
	switch value := o.(type) {
	case []byte:
		return value
	default:
		tmp := ParseString(o)
		return []byte(tmp)
	}
}

// ParseRune ...
func ParseRune(o interface{}) rune {
	if v, ok := o.(rune); ok {
		return v
	}
	return rune(ParseInt32(o))
}

// ParseRunes ...
func ParseRunes(o interface{}) []rune {
	if v, ok := o.([]rune); ok {
		return v
	}
	return []rune(ParseString(o))
}

// ParseString 任意对象转换为字符串
func ParseString(o interface{}) string {
	if o == nil {
		return ""
	}
	switch value := o.(type) {
	case string:
		return value
	case []byte:
		return string(value)
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		if value == nil {
			return ""
		}
		// 如果是复杂类型,则转换为json
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		}
		if kind == reflect.Ptr {
			return ParseString(rv.Elem().Interface())
		}
		// 如果不能转换为json的, 则原样输出为字符串
		if js, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(js)
		}
	}
}

// If 三元判断
func If(b bool, x, y interface{}) interface{} {
	if b {
		return x
	}
	return y
}

// SpliteAndTrimSpace 字符串分割并且去掉两边的空格
func SpliteAndTrimSpace(s, sep string) (res []string) {
	split := strings.Split(s, sep)
	for _, v := range split {
		res = append(res, strings.TrimSpace(v))
	}
	return
}

// InArray 是否存在给定的数组中
func InArray(arg interface{}, args []interface{}) bool {
	for _, v := range args {
		if New(v).String() == New(arg).String() {
			return true
		}
	}
	return false
}

// InArrayString 是否存在给定的字符串数组中
func InArrayString(arg string, args []string) bool {
	for _, v := range args {
		if v == arg {
			return true
		}
	}
	return false
}

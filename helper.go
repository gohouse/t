package t

import (
	"strings"
)

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

type iHelper interface {
	InArray(args []interface{}) bool
	InArrayString(args []string) bool
	SpliteAndTrimSpace(sep string) (res []string)
}

// SpliteAndTrimSpace 字符串分割并且去掉两边的空格
func (t Type) SpliteAndTrimSpace(sep string) (res []string) {
	return SpliteAndTrimSpace(t.String(), sep)
}

// InArray 是否存在给定的数组中
func (t Type) InArray(args []interface{}) bool {
	return InArray(t.val, args)
}

// InArrayString 是否存在给定的字符串数组中
func (t Type) InArrayString(args []string) bool {
	return InArrayString(t.String(), args)
}

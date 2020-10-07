package t

import (
	"strings"
)

// If 三元判断
func If(b bool, x, y interface{}) Type {
	if b {
		return New(x)
	}
	return New(y)
}

// SpliteAndTrimSpace 字符串分割并且去掉两边的空格
func SpliteAndTrimSpace(s, sep string) (res []string) {
	split := strings.Split(s, sep)
	for _, v := range split {
		if v != "" {
			res = append(res, strings.TrimSpace(v))
		}
	}
	return
}

// InArray 是否存在给定的数组中
func InArray(argv interface{}, haystack interface{}) bool {
	for _, v := range New(haystack).Slice() {
		if v.String() == New(argv).String() {
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

// Max 获取最大的数
func Max(args ...interface{}) Type {
	if len(args) == 0 {
		return New(0)
	}
	var max = New(args[0]).Float64()
	for k, arg := range args {
		if k == 0 {
			continue
		}
		tmp := New(arg).Float64()
		if tmp > max {
			max = tmp
		}
	}
	return New(max)
}

// Min 获取最小的数
func Min(args ...interface{}) Type {
	if len(args) == 0 {
		return New(0)
	}
	var min = New(args[0]).Float64()
	for k, arg := range args {
		if k == 0 {
			continue
		}
		tmp := New(arg).Float64()
		if tmp < min {
			min = tmp
		}
	}
	return New(min)
}

// ArrayUnique 数组去重
func ArrayUniqueString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// ArrayUnique 数组去重
func ArrayUnique(array interface{}) (newArr Slice) {
	arr := New(array).Slice()
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}
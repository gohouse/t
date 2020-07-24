package t

import (
	"encoding/json"
	"strings"
)

// iHandle json或者map,struct对象绑定和字段抽取
type iHandle interface {
	Bind(o interface{}) error
	Extract(keys ...interface{}) Type
	ExtractWithDefault(keys []interface{}, defaultVal interface{}) Type
}

// Bind 将绑定结果当做json,来绑定到对象上
func (tc TypeContext) Bind(o interface{}) error {
	return json.Unmarshal(tc.Bytes(), o)
}

// ExtractWithDefault 多层次抽取值, 如果是零值, 则返回给定的默认值
func (tc TypeContext) ExtractWithDefault(keys []interface{}, defaultVal interface{}) Type {
	v := tc.Extract(keys...)
	return If(v.IsZero(), New(defaultVal), v).(Type)
}

// Extract 多层次抽取值
// 例: New(`{"a":1,"b":["c",{"d":2}]}`).Extract("b.1.d") // 2
// 或者 New(`{"a":1,"b":["c",{"d":2}]}`).Extract("b",1,"d") // 2
// 以上两种方式, 只能使用其中一种, 不能两种混用, 如果key中有点., 则可以使用第二种,不会按照点去分割的
func (tc TypeContext) Extract(keys ...interface{}) Type {
	if len(keys) == 0 {
		return nil
	}
	var split []string
	if len(keys) == 1 {
		keyTmp := New(keys[0]).String()
		// 如果key包含了点,则为多级调用
		if strings.Contains(keyTmp, ".") {
			split = strings.Split(keyTmp, ".")
		} else { // 如果key不包含点, 则就是直接一级调用
			split = New(keys).SliceString()
		}
	} else {
		split = New(keys).SliceString()
	}

	// 取指定语言的配置
	//var currentMap = tc.Map()
	var currentVal interface{} = tc
	for _, item := range split {
		currentT := New(currentVal)
		if currentT.IsJsonMap() {
			var currentMap = currentT.Map()
			if v, ok := currentMap[New(item)]; ok {
				//currentMap = New(v).Map()
				currentVal = v
			} else {
				currentVal = nil
				break
			}
		} else if currentT.IsJsonSlice() {
			var currentSlice = currentT.Slice()
			var itemT = New(item)
			var itemIdx = int(itemT.Uint())
			if itemT.IsInteger() && len(currentSlice) > itemIdx {
				currentVal = currentSlice[itemIdx]
			} else {
				currentVal = nil
				break
			}
		} else {
			currentVal = nil
			break
		}
	}

	return New(currentVal)
}

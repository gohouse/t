package t

import (
	"encoding/json"
	"strings"
)

type iHandle interface {
	Bind(o interface{}) error
	Extract(key string, defaultVal ...interface{}) T
}

// Bind 将绑定结果当做json,来绑定到对象上
func (t Type) Bind(o interface{}) error {
	return json.Unmarshal(t.Bytes(), o)
}
// Extract 多层次抽取值
// 例: var m = New(`{"a": 2, "b":3,"33":{"cc":"d"}}`); println(m.Extract("33.cc")) // d
func (t Type) Extract(key string, defaultVal ...interface{}) T {
	if key == "" {
		return nil
	}
	var split []string
	// 如果key包含了点,则为多级调用
	if strings.Contains(key, ".") {
		split = strings.Split(key, ".")
	} else { // 如果key不包含点, 则就是直接一级调用
		split = []string{key}
	}

	// 取指定语言的配置
	var currentMap = t.Map()
	var currentVal interface{}
	for _, item := range split {
		if v, ok := currentMap[New(item)]; ok {
			currentMap = New(v).Map()
			currentVal = v
		} else {
			break
		}
	}

	// 如果没有取到且传入了默认值, 则返回默认值
	if currentVal == nil && len(defaultVal) > 0 {
		currentVal = defaultVal[0]
	}
	return New(currentVal)
}

package t

import "testing"

func TestTypeContext_IsFloat(t *testing.T) {
	//var a = "2342.2"
	var a float64 = 2342
	t.Log(New(a).IsFloat())
}


func TestTypeContext_IsInteger(t *testing.T) {
	var a = "2342"
	t.Log(New(a).IsInteger())
	t.Log(New(a).IsNumeric())
}

func TestTypeContext_IsChineseName(t *testing.T) {
	var a = "你好·啊"
	t.Log(New(a).IsChineseName())
	a = "你好•啊"
	t.Log(New(a).IsChineseName())
	a = "你好●啊"
	t.Log(New(a).IsChineseName())
}
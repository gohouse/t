package t

import "testing"

func TestType_IsFloat(t *testing.T) {
	//var a = "2342.2"
	var a float64 = 2342
	t.Log(New(a).IsFloat())
}


func TestType_IsInteger(t *testing.T) {
	var a = "2342"
	t.Log(New(a).IsInteger())
	t.Log(New(a).IsNumeric())
}
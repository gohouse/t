package t

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	var tmp = []interface{}{1,2,3}

	res := Max(tmp...)

	if !reflect.DeepEqual(res, 3) {
		t.Error("max 错误")
		return
	}
	t.Log("max 测试通过")
}

func TestMin(t *testing.T) {
	var tmp = []interface{}{"1",2,3,"4"}

	res := Min(tmp...)

	if !reflect.DeepEqual(res, 1) {
		t.Error("min 错误")
		return
	}
	t.Log("min 测试通过")
}

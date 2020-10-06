package t

import "testing"

func TestTypeContext_InArray(t *testing.T) {
	var a = []interface{}{1, 2, "a"}
	t.Log(New("a").InArray(a))
}

func TestArrayUnique(t *testing.T) {
	var a = []string{
		"a","b","b","d","c","a",
	}
	t.Log(ArrayUniqueString(a))
	t.Log(ArrayUnique(a).String())
}
package t

import (
	"testing"
)

func TestParseReader(t *testing.T) {
	var a = make([]byte, 3)
	b := New("你好12abc哈哈")
	buf := b.Reader()
	buf.Read(a)
	t.Log(string(a))
	buf.Read(a)
	t.Log(string(a))
}

package t

import (
	"testing"
)

func TestParser(t *testing.T) {
	var a = "12.34"
	//var a = map[string]interface{}{"a":1}
	t.Log(ParseString(a))
	t.Log(ParseFloat64(a))
	t.Log(ParseInt64(a))
	t.Log(ParseInt32(a))
	t.Log(ParseInt16(a))
	t.Log(ParseInt8(a))
	t.Log(ParseInt(a))
	t.Log(ParseUint64(a))
	t.Log(ParseUint32(a))
	t.Log(ParseUint16(a))
	t.Log(ParseUint8(a))
	t.Log(ParseUint(a))
	t.Log(ParseBool(a))
}

func TestParseString(t *testing.T) {
	var a = []byte("10")

	t.Log(ParseString(a))
}

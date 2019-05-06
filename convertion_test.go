package t

import (
	"testing"
)

func TestParser(t *testing.T) {
	var a = "8.8"
	t.Log(ParseString(a))
	t.Log(ParseFloat64(a))
	t.Log(ParseInt64(a))
	t.Log(ParseInt(a))
	t.Log(ParseUint64(a))
	t.Log(ParseUint(a))
	t.Log(ParseBool(a))
}

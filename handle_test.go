package t

import "testing"

func TestTypeContext_ExtractUrl(t *testing.T) {
	str := "abehttp://www.xx.com/a asdf http://a.com"
	res := New(str).ExtractUrl()

	t.Log(res)
}

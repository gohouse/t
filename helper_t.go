package t
type iHelper interface {
	InArray(args []interface{}) bool
	InArrayString(args []string) bool
	SpliteAndTrimSpace(sep string) (res []string)
	Length() int
}

// SpliteAndTrimSpace 字符串分割并且去掉两边的空格
func (tc TypeContext) SpliteAndTrimSpace(sep string) (res []string) {
	return SpliteAndTrimSpace(tc.String(), sep)
}

// InArray 是否存在给定的数组中
func (tc TypeContext) InArray(args []interface{}) bool {
	return InArray(tc.val, args)
}

// InArrayString 是否存在给定的字符串数组中
func (tc TypeContext) InArrayString(args []string) bool {
	return InArrayString(tc.String(), args)
}

// Length 数据长度, 以 []byte 统计
func (tc TypeContext) Length() int {
	return len(tc.Bytes())
}

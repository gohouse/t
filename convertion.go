package t

type iConvertion interface {

	Interface() interface{}
	// String 转换对象为 string 类型
	String() string
	Float64() float64
	Float32() float32
	Int64() int64
	Int() int
	Int32() int32
	Int16() int16
	Int8() int8
	Uint64() uint64
	Uint() uint
	Uint32() uint32
	Uint16() uint16
	Uint8() uint8
	Bool() bool
	Byte() byte
	Bytes() []byte
	Rune() rune
	Runes() []rune
}

// Interface ...
func (tc TypeContext) Interface() interface{} {
	return tc.val
}

// String ...
func (tc TypeContext) String() string {
	return ParseString(tc.val)
}

// Float64 ...
func (tc TypeContext) Float64() float64 {
	return ParseFloat64(tc.val)
}

// Float32 ...
func (tc TypeContext) Float32() float32 {
	return ParseFloat32(tc.val)
}

// Int64 ...
func (tc TypeContext) Int64() int64 {
	return ParseInt64(tc.val)
}

// Int ...
func (tc TypeContext) Int() int {
	return ParseInt(tc.val)
}

// Int32 ...
func (tc TypeContext) Int32() int32 {
	return int32(ParseInt64(tc.val))
}

// Int16 ...
func (tc TypeContext) Int16() int16 {
	return int16(ParseInt64(tc.val))
}

// Int8 ...
func (tc TypeContext) Int8() int8 {
	return int8(ParseInt64(tc.val))
}

// Uint64 ...
func (tc TypeContext) Uint64() uint64 {
	return ParseUint64(tc.val)
}

// Uint ...
func (tc TypeContext) Uint() uint {
	return ParseUint(tc.val)
}

// Uint32 ...
func (tc TypeContext) Uint32() uint32 {
	return uint32(ParseUint(tc.val))
}

// Uint16 ...
func (tc TypeContext) Uint16() uint16 {
	return uint16(ParseUint(tc.val))
}

// Uint8 ...
func (tc TypeContext) Uint8() uint8 {
	return uint8(ParseUint(tc.val))
}

// Bool ...
func (tc TypeContext) Bool() bool {
	return ParseBool(tc.val)
}

// Byte ...
func (tc TypeContext) Byte() byte {
	return ParseByte(tc.val)
}

// Bytes ...
func (tc TypeContext) Bytes() []byte {
	return ParseBytes(tc.val)
}

// Rune ...
func (tc TypeContext) Rune() rune {
	return ParseRune(tc.val)
}

// ParseRunes ...
func (tc TypeContext) Runes() []rune {
	return ParseRunes(tc.val)
}
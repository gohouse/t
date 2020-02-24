package t

// T ...
type T interface {
	iSlice
	iMap
	iHandle
	Interface() interface{}
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
}

// Type ...
type Type struct {
	val interface{}
}

var _ T = &Type{}

// New ...
func New(o interface{}) Type {
	switch o.(type) {
	case T:
		return Type{o.(T).Interface()}
	default:
		return Type{val: o}
	}
}

// Interface ...
func (t Type) Interface() interface{} {
	return t.val
}

// String ...
func (t Type) String() string {
	return ParseString(t.val)
}

// Float64 ...
func (t Type) Float64() float64 {
	return ParseFloat64(t.val)
}

// Float32 ...
func (t Type) Float32() float32 {
	return ParseFloat32(t.val)
}

// Int64 ...
func (t Type) Int64() int64 {
	return ParseInt64(t.val)
}

// Int ...
func (t Type) Int() int {
	return ParseInt(t.val)
}

// Int32 ...
func (t Type) Int32() int32 {
	return int32(ParseInt64(t.val))
}

// Int16 ...
func (t Type) Int16() int16 {
	return int16(ParseInt64(t.val))
}

// Int8 ...
func (t Type) Int8() int8 {
	return int8(ParseInt64(t.val))
}

// Uint64 ...
func (t Type) Uint64() uint64 {
	return ParseUint64(t.val)
}

// Uint ...
func (t Type) Uint() uint {
	return ParseUint(t.val)
}

// Uint32 ...
func (t Type) Uint32() uint32 {
	return uint32(ParseUint(t.val))
}

// Uint16 ...
func (t Type) Uint16() uint16 {
	return uint16(ParseUint(t.val))
}

// Uint8 ...
func (t Type) Uint8() uint8 {
	return uint8(ParseUint(t.val))
}

// Bool ...
func (t Type) Bool() bool {
	return ParseBool(t.val)
}

// Byte ...
func (t Type) Byte() byte {
	return ParseByte(t.val)
}

// Bytes ...
func (t Type) Bytes() []byte {
	return ParseBytes(t.val)
}

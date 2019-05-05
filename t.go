package t

type T interface {
	String() string
	Int64() int64
	Int() int
	Uint64() uint64
	Uint() uint
	Float64() float64
	Float32() float32
	Bool() bool
}
type Type struct {
	val interface{}
}

func New(o interface{}) T {
	return &Type{val:o}
}

func (t Type) String() string {
	return ParseString(t.val)
}

func (t Type) Int64() int64 {
	return ParseInt64(t.val)
}

func (t Type) Int() int {
	return ParseInt(t.val)
}

func (t Type) Uint64() uint64 {
	return ParseUint64(t.val)
}

func (t Type) Uint() uint {
	return ParseUint(t.val)
}

func (t Type) Float64() float64 {
	return ParseFloat64(t.val)
}

func (t Type) Float32() float32 {
	return ParseFloat32(t.val)
}

func (t Type) Bool() bool {
	return ParseBool(t.val)
}
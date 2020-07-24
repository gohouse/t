package t

// Type 超级type类型
type Type interface {
	iHelper     // 工具函数
	iDetermine  // 判断函数
	iSlice      // slice转换
	iMap        // map转换
	iHandle     // 复杂数据处理,如:json
	iConvertion // 基本类型转换
}

// TypeContext ...
type TypeContext struct {
	val interface{}
}

// 执行一次, 确保实现了 Type 接口, 在编译期间即可暴露出问题
var _ Type = &TypeContext{}

// New 初始化 Type 对象
func New(o interface{}) TypeContext {
	switch o.(type) {
	case Type:
		return TypeContext{val: o.(Type).Interface()}
	default:
		return TypeContext{val: o}
	}
}

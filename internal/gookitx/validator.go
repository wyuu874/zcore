package gookitx

import "github.com/gookit/validate"

type Errors = validate.Errors

// CustomValidator 自定义验证器
type CustomValidator struct{}

// ValidateStruct 验证结构体
func (c *CustomValidator) ValidateStruct(ptr any) error {
	v := validate.Struct(ptr)
	v.Validate() // 调用验证

	return v.Errors
}

// Engine 返回验证器引擎
func (c *CustomValidator) Engine() any {
	return nil
}

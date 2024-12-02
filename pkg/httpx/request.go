package httpx

import (
	"fmt"
	"github.com/wyuu874/zcore/internal/ginx"
	"github.com/wyuu874/zcore/internal/gookitx"
)

// BindingError 定义绑定错误的结构
type BindingError struct {
	// 操作类型（如：JSON绑定、查询参数绑定等）
	Operation string
	// 原始错误
	Err error
}

// Error 实现error接口
func (e *BindingError) Error() string {
	return fmt.Sprintf("绑定失败 - %s: %v", e.Operation, e.Err)
}

func init() {
	// 初始化验证器
	ginx.SetValidator(&gookitx.CustomValidator{})
}

// bind 通用的绑定函数，用于处理不同类型的绑定操作
// 参数说明：
// - v: 需要绑定的目标结构体
// - operation: 操作类型描述
// - binder: 具体的绑定函数
func bind(v interface{}, operation string, binder func(interface{}) error) error {
	if err := binder(v); err != nil {
		return handleBindError(err, operation)
	}
	return nil
}

// BindJSON 绑定JSON请求体到指定的结构体
// 参数说明：
// - c: 上下文
// - v: 目标结构体指针
func BindJSON(c *Context, v interface{}) error {
	return bind(v, "JSON数据绑定", c.ShouldBindJSON)
}

// BindQuery 绑定URL查询参数到指定的结构体
// 参数说明：
// - c: 上下文
// - v: 目标结构体指针
func BindQuery(c *Context, v interface{}) error {
	return bind(v, "URL查询参数绑定", c.ShouldBindQuery)
}

// BindForm 绑定表单数据到指定的结构体
// 参数说明：
// - c: 上下文
// - v: 目标结构体指针
func BindForm(c *Context, v interface{}) error {
	return bind(v, "表单数据绑定", c.ShouldBind)
}

// handleBindError 处理绑定过程中的错误
// 参数说明：
// - err: 原始错误
// - operation: 操作类型描述
func handleBindError(err error, operation string) error {
	// 处理验证器错误
	if validationErrors, ok := err.(gookitx.Errors); ok {
		if !validationErrors.Empty() {
			return &BindingError{
				Operation: operation,
				Err:       validationErrors.OneError(),
			}
		}
		return nil
	}

	// 处理其他类型错误
	return &BindingError{
		Operation: operation,
		Err:       err,
	}
}

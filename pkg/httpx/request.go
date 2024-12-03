package httpx

import (
	"github.com/wyuu874/zcore/internal/ginx"
	"github.com/wyuu874/zcore/internal/gookitx"
)

func init() {
	// 初始化验证器
	ginx.SetValidator(&gookitx.CustomValidator{})
}

// bind 通用的绑定函数，用于处理不同类型的绑定操作
// 参数说明：
// - v: 需要绑定的目标结构体
// - binder: 具体的绑定函数
func bind(v interface{}, binder func(interface{}) error) error {
	if err := binder(v); err != nil {
		return handleBindError(err)
	}
	return nil
}

// BindJSON 绑定JSON请求体到指定的结构体
// 参数说明：
// - c: 上下文
// - v: 目标结构体指针
func BindJSON(c *Context, v interface{}) error {
	return bind(v, c.ShouldBindJSON)
}

// BindQuery 绑定URL查询参数到指定的结构体
// 参数说明：
// - c: 上下文
// - v: 目标结构体指针
func BindQuery(c *Context, v interface{}) error {
	return bind(v, c.ShouldBindQuery)
}

// BindForm 绑定表单数据到指定的结构体
// 参数说明：
// - c: 上下文
// - v: 目标结构体指针
func BindForm(c *Context, v interface{}) error {
	return bind(v, c.ShouldBind)
}

// handleBindError 处理绑定过程中的错误
// 参数说明：
// - err: 原始错误
func handleBindError(err error) error {
	// 处理验证器错误
	if validationErrors, ok := err.(gookitx.Errors); ok {
		if !validationErrors.Empty() {
			return validationErrors.OneError()
		}
		return nil
	}

	// 处理其他类型错误
	return err
}

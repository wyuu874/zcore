package httpx

import (
	"github.com/wyuu874/zcore/pkg/locale"
	"net/http"
)

// ResponseCode 定义API响应状态码
type ResponseCode int

// API 响应状态码常量
const (
	CodeSuccess      ResponseCode = 200  // 请求成功
	CodeError        ResponseCode = 300  // 请求错误
	CodeUnauthorized ResponseCode = 401  // 未经授权
	CodeLogout       ResponseCode = 4011 // 已登出
	CodeRefreshToken ResponseCode = 4012 // 需要刷新Token
	CodeForbidden    ResponseCode = 403  // 访问被禁止
	CodeInternal     ResponseCode = 500  // 服务器内部错误
)

// ApiResponse 定义API响应结构
type ApiResponse struct {
	Code ResponseCode `json:"code"` // 响应状态码
	Msg  string       `json:"msg"`  // 响应消息
	Data interface{}  `json:"data"` // 响应数据
}

// responseParams 定义响应参数
type responseParams struct {
	context *Context
	code    ResponseCode
	msg     string
	data    interface{}
	msgArgs []map[string]interface{}
}

// newResponse 创建新的响应
func newResponse(params responseParams) *ApiResponse {
	return &ApiResponse{
		Code: params.code,
		Msg:  locale.Translate(params.context.GetString("lang"), params.msg, ApiMsgArgs(params.msgArgs...)),
		Data: params.data,
	}
}

// Json 发送JSON响应
// 参数说明：
// - c: 上下文
// - resp: API响应对象
func Json(c *Context, resp *ApiResponse) {
	c.JSON(http.StatusOK, resp)
}

// ApiSuccess 发送成功响应
// 参数说明：
// - c: 上下文
// - data: 响应数据
// - msg: 响应消息
// - msgArgs: 消息参数（可选）
func ApiSuccess(c *Context, data interface{}, msg string, msgArgs ...map[string]interface{}) {
	Json(c, newResponse(responseParams{
		context: c,
		code:    CodeSuccess,
		msg:     msg,
		data:    data,
		msgArgs: msgArgs,
	}))
}

// ApiError 发送错误响应
func ApiError(c *Context, msg string, data interface{}, msgArgs ...map[string]interface{}) {
	Json(c, newResponse(responseParams{
		context: c,
		code:    CodeError,
		msg:     msg,
		data:    data,
		msgArgs: msgArgs,
	}))
}

// ApiUnauthorized 发送未授权响应
func ApiUnauthorized(c *Context, msg string, data interface{}, msgArgs ...map[string]interface{}) {
	Json(c, newResponse(responseParams{
		context: c,
		code:    CodeUnauthorized,
		msg:     msg,
		data:    data,
		msgArgs: msgArgs,
	}))
}

// ApiLogout 发送登出响应
func ApiLogout(c *Context, msg string, data interface{}, msgArgs ...map[string]interface{}) {
	Json(c, newResponse(responseParams{
		context: c,
		code:    CodeLogout,
		msg:     msg,
		data:    data,
		msgArgs: msgArgs,
	}))
}

// ApiRefreshToken 发送刷新Token响应
func ApiRefreshToken(c *Context, msg string, data interface{}, msgArgs ...map[string]interface{}) {
	Json(c, newResponse(responseParams{
		context: c,
		code:    CodeRefreshToken,
		msg:     msg,
		data:    data,
		msgArgs: msgArgs,
	}))
}

// ApiForbidden 发送禁止访问响应
func ApiForbidden(c *Context, msg string, data interface{}, msgArgs ...map[string]interface{}) {
	Json(c, newResponse(responseParams{
		context: c,
		code:    CodeForbidden,
		msg:     msg,
		data:    data,
		msgArgs: msgArgs,
	}))
}

// ApiInternal 发送内部错误响应
func ApiInternal(c *Context, msg string, data interface{}, msgArgs ...map[string]interface{}) {
	Json(c, newResponse(responseParams{
		context: c,
		code:    CodeInternal,
		msg:     msg,
		data:    data,
		msgArgs: msgArgs,
	}))
}

// ApiMsgArgs 处理消息参数
// 参数说明：
// - msgArgs: 可变参数，消息参数列表
// 返回值：
// - map[string]interface{}: 处理后的消息参数映射
func ApiMsgArgs(msgArgs ...map[string]interface{}) map[string]interface{} {
	if len(msgArgs) > 0 {
		return msgArgs[0]
	}
	return make(map[string]interface{})
}

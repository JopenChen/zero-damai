package errx

import (
	"fmt"
)

/*
*
- 服务编号：10=API、20=RPC、30=Websocket。
- 接口类型（增删查改）: 01=POST，02=DELETE，03=GET，04=PUT。
- 错误来源: 表示错误属于那种类型，两位数字递增。01=内部调用错误、02=数据库报错、03=缓存服务报错、04=第三方服务报错、10=参数错误、11=具体业务错误。
示例：API模块 用户登录接口 手机号格式错误

	10  +    01    +     10    = 100110

*
*/
const (
	// OK 成功
	OK uint32 = 0
	// ServerCommonError 默认系统错误
	ServerCommonError uint32 = 999901
	// RequestParamError 参数错误
	RequestParamError uint32 = 999910
	// PageNotFound 404
	PageNotFound uint32 = 999911
	// DbError 数据库操作失败
	DbError uint32 = 999902

	// UserNotRegisterError 用户未注册
	UserNotRegisterError = 200112
	// UserNotFoundError 用户不存在
	UserNotFoundError = 200113
	// UserAlreadyExistError 用户已存在
	UserAlreadyExistError = 200114
	// UserForbiddenError 用户被禁用
	UserForbiddenError = 200115
	// UserLoginFailError 用户登录失败
	UserLoginFailError = 200116

	// LoginTypeNotSupportError 登录方式未支持
	LoginTypeNotSupportError = 200117
	// LoginPasswordIncorrectError 登录密码错误
	LoginPasswordIncorrectError = 200118
)

type CodeError struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var message map[uint32]string

const (
	UnknownError = "未知错误，请联系客服人员！"
)

func init() {
	message = make(map[uint32]string)
	message[OK] = "成功"
	message[ServerCommonError] = "服务繁忙，稍后再试"
	message[RequestParamError] = "参数错误"
	message[PageNotFound] = "页面不存在"
	message[DbError] = "数据库操作错误"

	message[UserNotRegisterError] = "用户未注册"
	message[UserNotFoundError] = "用户不存在"
	message[UserAlreadyExistError] = "用户已存在"
	message[UserForbiddenError] = "用户被禁用"
	message[UserLoginFailError] = "用户登录失败"
	message[LoginTypeNotSupportError] = "登录方式未支持"
	message[LoginPasswordIncorrectError] = "登录密码错误"
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.Code
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Msg
}

// GetErrData 开发环境返回给前端所有错误堆栈信息
func (e *CodeError) GetErrData() interface{} {
	return e.Data
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("Code:%d，Msg:%s", e.Code, e.Msg)
}

func (e *CodeError) ErrorDetail() string {
	return fmt.Sprintf("Code:%d，Msg:%s，Data:%+v", e.Code, e.Msg, e.Data)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{
		Code: errCode,
		Msg:  errMsg,
	}
}

func NewErrCode(errCode uint32) *CodeError {
	if msg, ok := message[errCode]; ok {
		return &CodeError{
			Code: errCode,
			Msg:  msg,
			Data: nil,
		}
	} else {
		return &CodeError{
			Code: errCode,
			Msg:  UnknownError,
			Data: nil,
		}
	}
}

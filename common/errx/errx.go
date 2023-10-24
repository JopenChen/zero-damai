package errx

import (
	"fmt"
)

/**
常用通用固定错误
*/

// 状态码返回
const (
	HTTP_STATUS_OK             uint32 = 200 //正常返回消息，什么问题没有
	HTTP_STATUS_BAD_REQUEST    uint32 = 400 //请求错误
	HTTP_STATUS_UNAUTHORIZED   uint32 = 401 //无权限访问（已登录）
	HTTP_STATUS_FORBIDDEN      uint32 = 403 //拒绝访问（未登录）
	HTTP_STATUS_INTERNAL_ERROR uint32 = 500 //程序奔溃
)

/*
*
- 应用编号，表示错误属于哪个应用，1-3位为服务模块编号，方便定位。
- 功能域标识，表示错误属于应用中的哪个功能模块，两位数字递增。
- 接口类型（增删查改），1是添加接口，2是删除接口，3是读取接口，4是更新接口，5是导入，6是导出 其他的累加
- 错误来源，表示错误属于那种类型，两位数字递增。01是内部调用错误，02是数据库报错，03是缓存服务报错，04是其他第三方服务报错，10是参数错误，11及以上是具体的业务报错
示例：小程序端 用户登录 手机号格式错误
000+01+3+01
*
*/
const (
	//系统全局
	OK                  uint32 = 0        //执行成功
	SERVER_COMMON_ERROR uint32 = 99950001 //默认系统错误
	REUQEST_PARAM_ERROR uint32 = 99940110 //参数错误
	PAGE_NOT_FOUND      uint32 = 99940401 //404
	DB_ERROR            uint32 = 99900102 //数据库操作失败

)

type CodeError struct {
	errCode uint32
	errMsg  string
	errData interface{}
}

var message map[uint32]string

const (
	UNKNOW_ERROR = "网络开小差啦,请稍后再来！"
)

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[PAGE_NOT_FOUND] = "页面不存在"
	message[DB_ERROR] = "数据库操作失败"
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

// 开发环境返回给前端所有错误堆栈信息
func (e *CodeError) GetErrData() interface{} {
	return e.errData
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func (e *CodeError) ErrorDetail() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s，ErrData:%+v", e.errCode, e.errMsg, e.errData)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{
		errCode: errCode,
		errMsg:  errMsg,
	}
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{
		errCode: errCode,
		errMsg:  MapErrMsg(errCode),
	}
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return UNKNOW_ERROR
	}
}

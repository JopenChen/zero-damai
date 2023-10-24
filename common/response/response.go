package response

import (
	"github.com/JopenChen/zero-damai/common/errx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// EmptyObj 空对象
type emptyObj struct{}

// Response 公共响应体
type Response struct {
	Code     uint32      `json:"code"`
	Msg      string      `json:"msg"`
	Response interface{} `json:"response"`
}

func success(resp interface{}) *Response {
	return &Response{0, "OK", resp}
}

func fail(errCode uint32, errMsg string) *Response {
	return &Response{errCode, errMsg, &emptyObj{}}
}

// Handler 处理所有响应
func Handler(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		var errCode uint32
		var errMsg string
		// 获取最内层错误
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*errx.CodeError); ok {
			// 自定义错误类型
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if grpcStatus, ok := status.FromError(causeErr); ok {
				// GRPC 错误类型
				errCode = uint32(grpcStatus.Code())
				errMsg = grpcStatus.Message()
			}
		}

		httpx.WriteJson(w, http.StatusBadRequest, fail(errCode, errMsg))
	}
}

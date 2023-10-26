package performance

import (
	"net/http"

	"github.com/JopenChen/zero-damai/common/response"
	"github.com/JopenChen/zero-damai/gateway/internal/logic/v0/performance"
	"github.com/JopenChen/zero-damai/gateway/internal/svc"
	"github.com/JopenChen/zero-damai/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PerformanceRetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PerformanceRetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := performance.NewPerformanceRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.PerformanceRetrieve(&req)
		response.Handler(r, w, resp, err)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}

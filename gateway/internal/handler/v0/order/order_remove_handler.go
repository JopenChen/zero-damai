package order

import (
	"net/http"

	"gateway/internal/logic/v0/order"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/JopenChen/zero-damai/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderRemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GeneralPathId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderRemoveLogic(r.Context(), svcCtx)
		resp, err := l.OrderRemove(&req)
		response.Handler(r, w, resp, err)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}

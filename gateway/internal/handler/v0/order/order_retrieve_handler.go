package order

import (
	"net/http"

	"gateway/internal/logic/v0/order"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/JopenChen/zero-damai/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderRetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderRetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.OrderRetrieve(&req)
		response.Handler(r, w, resp, err)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}

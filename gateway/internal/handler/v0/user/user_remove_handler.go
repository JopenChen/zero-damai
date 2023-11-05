package user

import (
	"net/http"

	"gateway/internal/logic/v0/user"
	"gateway/internal/svc"
	"gateway/internal/types"
	"github.com/JopenChen/zero-damai/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GeneralPathId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserRemoveLogic(r.Context(), svcCtx)
		resp, err := l.UserRemove(&req)
		response.Handler(r, w, resp, err)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}

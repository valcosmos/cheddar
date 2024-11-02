package user

import (
	"net/http"

	"cheddar/internal/logic/v1/user"
	"cheddar/internal/svc"
	"cheddar/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

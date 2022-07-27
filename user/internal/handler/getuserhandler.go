package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/leave8080/zero_demo/user/internal/logic"
	"github/leave8080/zero_demo/user/internal/svc"
	"github/leave8080/zero_demo/user/internal/types"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

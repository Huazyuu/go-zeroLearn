package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero_learn/api_learn/user/apiv2/internal/logic"
	"gozero_learn/api_learn/user/apiv2/internal/svc"
	"gozero_learn/api_learn/user/apiv2/internal/types"
	"net/http"

	"gozero_learn/api_learn/user/common/response"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(r, w, resp, err)

	}
}

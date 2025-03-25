package handler

import (
	"gozero_learn/api_learn/user/common/response"
	"net/http"

	"gozero_learn/api_learn/user/apiv3_prefix/internal/logic"
	"gozero_learn/api_learn/user/apiv3_prefix/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)

	}
}

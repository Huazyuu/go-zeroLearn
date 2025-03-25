package handler

import (
	"gozero_learn/api_learn/user/apiv2/internal/logic"
	"gozero_learn/api_learn/user/apiv2/internal/svc"
	"gozero_learn/api_learn/user/common/response"
	"net/http"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)

	}
}

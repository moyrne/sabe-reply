package handler

import (
	"net/http"

	"github.com/moyrne/sabe-reply/api-service/internal/logic"
	"github.com/moyrne/sabe-reply/api-service/internal/svc"
	"github.com/moyrne/sabe-reply/api-service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SabeReplyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SabeReplyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSabeReplyLogic(r.Context(), svcCtx)
		resp, err := l.SabeReply(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

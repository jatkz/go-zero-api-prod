package handler

import (
	"net/http"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := types.Response{Message: "success"}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

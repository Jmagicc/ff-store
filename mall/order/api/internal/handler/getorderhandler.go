package handler

import (
	"net/http"

	"ff-store/mall/order/api/internal/logic"
	"ff-store/mall/order/api/internal/svc"
	"ff-store/mall/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

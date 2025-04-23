package health

import (
	"net/http"

	"go-mq/application/dto"
	"go-mq/common/response"
	"go-mq/infrastructure/provider"
	"go-mq/infrastructure/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// SaveHandler handles requests to create a new health record
func SaveHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.HealthRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		healthService := provider.InitializeHealthService(svc)
		resp, err := healthService.Save(r.Context(), &req)
		response.Response(r, w, resp, err)
	}
}

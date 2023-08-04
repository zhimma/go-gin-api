package article

import (
	"github.com/zhimma/go-gin-api/internal/pkg/core"
)

type destroyRequest struct{}

type destroyResponse struct{}

// Destroy 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body destroyRequest true "请求信息"
// @Success 200 {object} destroyResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/articles/{id} [DELETE]
func (h *handler) Destroy() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}

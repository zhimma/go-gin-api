package article

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 更新文章
// @Summary 更新文章
// @Description 更新文章
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/articles/{id} [PUT]
func (h *handler) Update() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}

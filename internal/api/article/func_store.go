package article

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type storeRequest struct{}

type storeResponse struct{}

// Store 新增文章
// @Summary 新增文章
// @Description 新增文章
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body storeRequest true "请求信息"
// @Success 200 {object} storeResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/articles [POST]
func (h *handler) Store() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}

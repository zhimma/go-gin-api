package article

import (
	"github.com/zhimma/go-gin-api/internal/code"
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/pkg/validation"
	"github.com/zhimma/go-gin-api/internal/services/article"
	"net/http"
)

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
		storeData := new(article.StoreData)
		if err := ctx.ShouldBind(storeData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err))
			return
		}

		id, err := h.articleService.Store(ctx, storeData)
		if id == 0 || err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.StoreError,
				code.Text(code.StoreError)).WithError(err),
			)
			return
		}

		ctx.Payload(struct {
			Id int32 `json:"id"`
		}{
			id,
		})
	}
}

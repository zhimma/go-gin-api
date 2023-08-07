package article

import (
	"github.com/zhimma/go-gin-api/internal/code"
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/pkg/validation"
	"github.com/zhimma/go-gin-api/internal/services/article"
	"net/http"
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
		updateData := new(article.UpdateData)
		if err := ctx.ShouldBind(&updateData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
		err := h.articleService.Update(ctx, updateData.Id, updateData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UpdateError,
				code.Text(code.UpdateError)).WithError(err),
			)
			return
		}
		ctx.Payload(struct {
			Id int32 `json:"id"`
		}{
			updateData.Id,
		})
		return
	}
}

package article

import (
	"github.com/zhimma/go-gin-api/internal/code"
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/pkg/validation"
	"github.com/zhimma/go-gin-api/internal/services/article"
	"net/http"
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
		req := new(article.IdData)
		if err := ctx.ShouldBindURI(&req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err))
			return
		}

		if err := h.articleService.Destroy(ctx, req.Id); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.DestroyError,
				code.Text(code.DestroyError)).WithError(err))
			return
		}
		ctx.Payload(struct {
			Id int32 `json:"id"`
		}{req.Id})
		return
	}
}

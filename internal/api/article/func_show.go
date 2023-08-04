package article

import (
	"github.com/zhimma/go-gin-api/internal/code"
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/services/article"
	"net/http"
	"time"
)

type showResponse struct {
	Id         int32     `json:"id"`          // 主键
	CategoryId int32     `json:"category_id"` // 父类ID
	Title      string    `json:"title"`       // 菜单名称
	ShortTitle string    `json:"short_title"` // 短标题
	MainImage  string    `json:"main_image"`  // 主图
	Content    string    `json:"content"`     // 正文
	Sort       int32     `json:"sort"`        // 排序
	Status     int32     `json:"status"`      // 是否启用 1:是 -1:否
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`  // 更新时间
}

// Show 文章详情
// @Summary 文章详情
// @Description 文章详情
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body showRequest true "请求信息"
// @Success 200 {object} showResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/articles/{id} [GET]
func (h *handler) Show() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(article.SearchData)
		res := new(showResponse)
		if err := ctx.ShouldBindURI(&req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithError(err))
			return
		}
		detail, err := h.articleService.Show(ctx, req)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithError(err))
			return
		}
		res.Id = detail.Id
		res.CategoryId = detail.CategoryId
		res.Title = detail.Title
		res.ShortTitle = detail.ShortTitle
		res.MainImage = detail.MainImage
		res.Content = detail.Content
		res.Sort = detail.Sort
		res.Status = detail.Status
		res.CreatedAt = detail.CreatedAt
		res.UpdatedAt = detail.UpdatedAt
		ctx.Payload(res)
	}
}

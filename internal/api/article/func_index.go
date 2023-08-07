package article

import (
	"github.com/zhimma/go-gin-api/internal/code"
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/pkg/validation"
	"github.com/zhimma/go-gin-api/internal/services/article"
	"net/http"
	"time"
)

type listData struct {
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
type indexResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int   `json:"current_page"`
		PerPageCount int64 `json:"per_page_count"`
	} `json:"pagination"`
}

// Index 文章列表
// @Summary 文章列表
// @Description 文章列表
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body indexRequest true "请求信息"
// @Success 200 {object} indexResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/articles [GET]
func (h *handler) Index() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(article.ListData)
		res := new(indexResponse)
		if err := ctx.ShouldBind(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err))
			return
		}
		pageData, err := h.articleService.Paginate(ctx, req)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithError(err),
			)
			return
		}

		res.List = make([]listData, len(pageData.List))
		res.Pagination.PerPageCount = pageData.Pagination.PerPageCount
		res.Pagination.Total = pageData.Pagination.Total
		res.Pagination.CurrentPage = pageData.Pagination.CurrentPage

		for k, record := range pageData.List {
			tmpData := listData{
				Id:         record.Id,
				CategoryId: record.CategoryId,
				Title:      record.Title,
				ShortTitle: record.ShortTitle,
				MainImage:  record.MainImage,
				Content:    record.Content,
				Sort:       record.Sort,
				Status:     record.Status,
				UpdatedAt:  record.UpdatedAt,
			}
			res.List[k] = tmpData
		}

		ctx.Payload(res)
	}
}

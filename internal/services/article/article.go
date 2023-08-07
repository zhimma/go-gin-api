package article

import (
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/repository/mysql"
	"github.com/zhimma/go-gin-api/internal/repository/mysql/article"
	"github.com/zhimma/go-gin-api/internal/repository/redis"
)

type ListData struct {
	Id         int32  `form:"id" json:"id"`
	Page       int    `form:"page" json:"page"`
	PageSize   int    `form:"page_size" json:"page_size"`
	CategoryId int32  `form:"category_id" json:"category_id"`
	Title      string `form:"title" json:"title"`
	ShortTitle string `form:"short_title" json:"short_title"`
	Status     int32  `form:"status" json:"status"`
}

type IdData struct {
	Id int32 `uri:"id" form:"id" json:"id" binding:"required"`
}

type StoreData struct {
	CategoryId int32  `form:"category_id" json:"category_id" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
	ShortTitle string `form:"short_title" json:"short_title" binding:"required"`
	MainImage  string `form:"main_image" json:"main_image" binding:"required"` // 主图
	Content    string `form:"content" json:"content" binding:"required"`       // 正文
	Sort       int32  `form:"sort" json:"sort" binding:"required"`             // 排序
	Status     int32  `form:"status" json:"status" binding:"required"`         // 是否启用 1:是 -1:否
}

type UpdateData struct {
	IdData
	StoreData
}

// Pagination 分页meta参数
type Pagination struct {
	Total        int64 `json:"total"`
	CurrentPage  int   `json:"current_page"`
	PerPageCount int64 `json:"per_page_count"`
}

// PageResult 分页结果
type PageResult struct {
	List       []*article.Article `json:"list"`
	Pagination Pagination         `json:"pagination"`
}

type Service interface {
	i()
	Paginate(ctx core.Context, listData *ListData) (result PageResult, err error)
	Show(ctx core.Context, id int32) (result *article.Article, err error)
	Store(ctx core.Context, storeData *StoreData) (id int32, err error)
	Update(ctx core.Context, id int32, updateData *UpdateData) (err error)
	Destroy(ctx core.Context, id int32) (err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func (s *service) Paginate(ctx core.Context, listData *ListData) (result PageResult, err error) {
	page := listData.Page
	if page == 0 {
		page = 1
	}

	pageSize := listData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	result.Pagination.CurrentPage = page
	result.Pagination.PerPageCount = int64(pageSize)

	offset := (page - 1) * pageSize

	qb := article.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	if listData.Title != "" {
		qb.WhereCategoryId(mysql.EqualPredicate, listData.CategoryId)
	}

	if listData.Title != "" {
		qb.WhereTitle(mysql.LikePredicate, listData.Title)
	}

	if listData.ShortTitle != "" {
		qb.WhereShortTitle(mysql.LikePredicate, listData.ShortTitle)
	}

	if listData.Status != 0 {
		qb.WhereStatus(mysql.LikePredicate, listData.Status)
	}
	result.Pagination.Total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return result, err
	}
	result.List, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return result, err
	}
	return
}

func (s *service) Show(ctx core.Context, id int32) (result *article.Article, err error) {
	qb := article.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if id != 0 {
		qb.WhereId(mysql.EqualPredicate, id)
	} else {
		return
	}

	result, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}

func (s *service) Store(ctx core.Context, data *StoreData) (id int32, err error) {
	model := article.NewModel()
	model.CategoryId = data.CategoryId
	model.Title = data.Title
	model.ShortTitle = data.ShortTitle
	model.MainImage = data.MainImage
	model.Content = data.Content
	model.Sort = data.Sort
	model.Status = data.Status
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.UpdatedUser = ctx.SessionUserInfo().UserName
	model.IsDeleted = -1
	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}

func (s *service) Update(ctx core.Context, id int32, updateData *UpdateData) (err error) {
	qb := article.NewQueryBuilder()
	data := map[string]interface{}{
		"category_id":  updateData.CategoryId,
		"title":        updateData.Title,
		"short_title":  updateData.ShortTitle,
		"main_image":   updateData.MainImage,
		"content":      updateData.Content,
		"sort":         updateData.Sort,
		"status":       updateData.Status,
		"updated_user": ctx.SessionUserInfo().UserName,
	}
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}
	return
}
func (s *service) Destroy(ctx core.Context, id int32) (err error) {
	qb := article.NewQueryBuilder()
	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) i() {}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{db: db, cache: cache}
}

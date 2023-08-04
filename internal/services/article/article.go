package article

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/article"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
)

type SearchData struct {
	Id         int32  `form:"id" json:"id"`
	Page       int    `form:"page" json:"page"`
	PageSize   int    `form:"page_size" json:"page_size"`
	CategoryId int32  `form:"category_id" json:"category_id"`
	Title      string `form:"title" json:"title"`
	ShortTitle string `form:"short_title" json:"short_title"`
	Status     int32  `form:"status" json:"status"`
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
	Paginate(ctx core.Context, searchData *SearchData) (result PageResult, err error)
	Show(ctx core.Context, searchData *SearchData) (result *article.Article, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func (s *service) Paginate(ctx core.Context, searchData *SearchData) (result PageResult, err error) {
	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	result.Pagination.CurrentPage = page
	result.Pagination.PerPageCount = int64(pageSize)

	offset := (page - 1) * pageSize

	qb := article.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	if searchData.Title != "" {
		qb.WhereCategoryId(mysql.EqualPredicate, searchData.CategoryId)
	}

	if searchData.Title != "" {
		qb.WhereTitle(mysql.LikePredicate, searchData.Title)
	}

	if searchData.ShortTitle != "" {
		qb.WhereShortTitle(mysql.LikePredicate, searchData.ShortTitle)
	}

	if searchData.Status != 0 {
		qb.WhereStatus(mysql.LikePredicate, searchData.Status)
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

func (s *service) Show(ctx core.Context, searchData *SearchData) (result *article.Article, err error) {
	qb := article.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, searchData.Id)
	}

	result, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
func (s *service) i() {}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{db: db, cache: cache}
}

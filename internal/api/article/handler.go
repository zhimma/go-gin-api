package article

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/article"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	// Index 文章列表
	// @Tags API.admin
	// @Router /api/admin/articles [GET]
	Index() core.HandlerFunc
	// Store 新增文章
	// @Tags API.admin
	// @Router /api/admin/articles [POST]
	Store() core.HandlerFunc
	// Update 更新文章
	// @Tags API.admin
	// @Router /api/admin/articles/{id} [PUT]
	Update() core.HandlerFunc
	// Show 文章详情
	// @Tags API.admin
	// @Router /api/admin/articles/{id} [GET]
	Show() core.HandlerFunc
	// Destroy 删除文章
	// @Tags API.admin
	// @Router /api/admin/articles/{id} [DELETE]
	Destroy() core.HandlerFunc
}

type handler struct {
	logger         *zap.Logger
	cache          redis.Repo
	hashids        hash.Hash
	articleService article.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:         logger,
		cache:          cache,
		hashids:        hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		articleService: article.New(db, cache),
	}
}

func (h *handler) i() {}

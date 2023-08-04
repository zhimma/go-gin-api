package article

import (
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/repository/mysql"
	"github.com/zhimma/go-gin-api/internal/repository/redis"

	"go.uber.org/zap"
)

const minBusinessCode = 20000

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("article_list", nil)
	}
}

func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("article_create", nil)
	}
}

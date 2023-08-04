package admin

import (
	"github.com/zhimma/go-gin-api/configs"
	"github.com/zhimma/go-gin-api/internal/pkg/core"
	"github.com/zhimma/go-gin-api/internal/pkg/password"
	"github.com/zhimma/go-gin-api/internal/repository/mysql"
	"github.com/zhimma/go-gin-api/internal/repository/mysql/admin"
	"github.com/zhimma/go-gin-api/internal/repository/redis"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}

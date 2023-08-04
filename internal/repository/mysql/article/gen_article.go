///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package article

import (
	"fmt"
	"time"

	"github.com/zhimma/go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Article {
	return new(Article)
}

func NewQueryBuilder() *articleQueryBuilder {
	return new(articleQueryBuilder)
}

func (t *Article) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type articleQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *articleQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *articleQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Article{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *articleQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Article{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *articleQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Article{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *articleQueryBuilder) First(db *gorm.DB) (*Article, error) {
	ret := &Article{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *articleQueryBuilder) QueryOne(db *gorm.DB) (*Article, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *articleQueryBuilder) QueryAll(db *gorm.DB) ([]*Article, error) {
	var ret []*Article
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *articleQueryBuilder) Limit(limit int) *articleQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *articleQueryBuilder) Offset(offset int) *articleQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *articleQueryBuilder) WhereId(p mysql.Predicate, value int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereIdIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereIdNotIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderById(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereCategoryId(p mysql.Predicate, value int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereCategoryIdIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereCategoryIdNotIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByCategoryId(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_id "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereTitle(p mysql.Predicate, value string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereTitleIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereTitleNotIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByTitle(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "title "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereShortTitle(p mysql.Predicate, value string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "short_title", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereShortTitleIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "short_title", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereShortTitleNotIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "short_title", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByShortTitle(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "short_title "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereMainImage(p mysql.Predicate, value string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "main_image", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereMainImageIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "main_image", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereMainImageNotIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "main_image", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByMainImage(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "main_image "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereContent(p mysql.Predicate, value string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereContentIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereContentNotIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByContent(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereSort(p mysql.Predicate, value int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereSortIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereSortNotIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderBySort(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereStatus(p mysql.Predicate, value int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereStatusIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereStatusNotIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByStatus(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereIsDeletedIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereIsDeletedNotIn(value []int32) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByIsDeleted(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereCreatedAtIn(value []time.Time) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByCreatedAt(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereCreatedUserIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereCreatedUserNotIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByCreatedUser(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereUpdatedAtIn(value []time.Time) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByUpdatedAt(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *articleQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereUpdatedUserIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) WhereUpdatedUserNotIn(value []string) *articleQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *articleQueryBuilder) OrderByUpdatedUser(asc bool) *articleQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}

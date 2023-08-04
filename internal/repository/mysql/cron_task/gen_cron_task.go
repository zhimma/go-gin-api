///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package cron_task

import (
	"fmt"
	"time"

	"github.com/zhimma/go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *CronTask {
	return new(CronTask)
}

func NewQueryBuilder() *cronTaskQueryBuilder {
	return new(cronTaskQueryBuilder)
}

func (t *CronTask) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type cronTaskQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *cronTaskQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *cronTaskQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&CronTask{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *cronTaskQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&CronTask{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *cronTaskQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CronTask{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *cronTaskQueryBuilder) First(db *gorm.DB) (*CronTask, error) {
	ret := &CronTask{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *cronTaskQueryBuilder) QueryOne(db *gorm.DB) (*CronTask, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *cronTaskQueryBuilder) QueryAll(db *gorm.DB) ([]*CronTask, error) {
	var ret []*CronTask
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *cronTaskQueryBuilder) Limit(limit int) *cronTaskQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *cronTaskQueryBuilder) Offset(offset int) *cronTaskQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *cronTaskQueryBuilder) WhereId(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereIdIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereIdNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderById(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereName(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNameIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNameNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByName(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereSpec(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "spec", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereSpecIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "spec", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereSpecNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "spec", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderBySpec(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "spec "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCommand(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "command", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCommandIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "command", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCommandNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "command", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByCommand(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "command "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereProtocol(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "protocol", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereProtocolIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "protocol", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereProtocolNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "protocol", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByProtocol(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "protocol "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereHttpMethod(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "http_method", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereHttpMethodIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "http_method", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereHttpMethodNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "http_method", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByHttpMethod(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "http_method "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereTimeout(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "timeout", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereTimeoutIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "timeout", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereTimeoutNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "timeout", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByTimeout(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "timeout "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRetryTimes(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retry_times", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRetryTimesIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retry_times", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRetryTimesNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retry_times", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByRetryTimes(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "retry_times "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRetryInterval(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retry_interval", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRetryIntervalIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retry_interval", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRetryIntervalNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retry_interval", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByRetryInterval(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "retry_interval "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyStatus(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_status", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyStatusIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_status", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyStatusNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByNotifyStatus(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "notify_status "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyType(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_type", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyTypeIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_type", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyTypeNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByNotifyType(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "notify_type "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyReceiverEmail(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_receiver_email", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyReceiverEmailIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_receiver_email", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyReceiverEmailNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_receiver_email", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByNotifyReceiverEmail(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "notify_receiver_email "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyKeyword(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_keyword", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyKeywordIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_keyword", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereNotifyKeywordNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "notify_keyword", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByNotifyKeyword(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "notify_keyword "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRemark(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRemarkIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereRemarkNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByRemark(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "remark "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereIsUsed(p mysql.Predicate, value int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereIsUsedIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereIsUsedNotIn(value []int32) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByIsUsed(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_used "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCreatedAtIn(value []time.Time) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByCreatedAt(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCreatedUserIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereCreatedUserNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByCreatedUser(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereUpdatedAtIn(value []time.Time) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByUpdatedAt(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *cronTaskQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereUpdatedUserIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) WhereUpdatedUserNotIn(value []string) *cronTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cronTaskQueryBuilder) OrderByUpdatedUser(asc bool) *cronTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}

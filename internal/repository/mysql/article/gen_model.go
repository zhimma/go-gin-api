package article

import "time"

// Article 文章表
//
//go:generate gormgen -structs Article -input .
type Article struct {
	Id          int32     // 主键
	CategoryId  int32     // 父类ID
	Title       string    // 菜单名称
	ShortTitle  string    // 短标题
	MainImage   string    // 主图
	Content     string    // 正文
	Sort        int32     // 排序
	Status      int32     // 是否启用 1:是 -1:否
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}

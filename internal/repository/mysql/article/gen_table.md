#### go_gin_api.article 
文章表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | category_id | 父类ID | int(11) unsigned |  | NO |  | 0 |
| 3 | title | 菜单名称 | varchar(32) |  | NO |  |  |
| 4 | short_title | 短标题 | varchar(100) |  | NO |  |  |
| 5 | main_image | 主图 | varchar(60) |  | NO |  |  |
| 6 | content | 正文 | text |  | YES |  |  |
| 7 | sort | 排序 | int(11) unsigned |  | NO |  | 0 |
| 8 | status | 是否启用 1:是 -1:否 | tinyint(1) |  | NO |  | 1 |
| 9 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 10 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 11 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 12 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 13 | updated_user | 更新人 | varchar(60) |  | NO |  |  |

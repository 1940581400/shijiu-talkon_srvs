package utils

import "gorm.io/gorm"

type PageFunc func(db *gorm.DB) *gorm.DB

// Paginate 分页公用组件
func Paginate(page, pageSize int32) PageFunc {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

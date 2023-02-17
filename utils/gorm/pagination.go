package gorm

import (
	"gorm.io/gorm"
)

type Pagination struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageIndex <= 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}

func Paginate(pagination Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pagination.PageIndex - 1) * pagination.GetPageSize()
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pagination.GetPageSize())
	}
}

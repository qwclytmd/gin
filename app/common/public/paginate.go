package public

import (
	"gorm.io/gorm"
)

type PaginationResult struct {
	Data       func(db *gorm.DB) *gorm.DB
	TotalCount int64
}

func Paginate(db *gorm.DB, page int, size int) *PaginationResult {
	var totalCount int64
	db.Count(&totalCount)

	data := func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		return db.Offset((page - 1) * size).Limit(size)
	}

	return &PaginationResult{
		Data:       data,
		TotalCount: totalCount,
	}
}

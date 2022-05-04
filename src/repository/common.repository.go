package repository

import (
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
	"gorm.io/gorm"
	"math"
)

func Pagination(value interface{}, meta *proto.PaginationMetadata) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var totalItems int64
		db.Model(value).Count(&totalItems)

		meta.TotalItem = totalItems
		totalPages := math.Ceil(float64(totalItems) / float64(meta.ItemsPerPage))
		meta.TotalPage = int64(totalPages)

		if meta.CurrentPage < 1 {
			meta.CurrentPage = 1
		}

		switch {
		case meta.ItemsPerPage > 100:
			meta.ItemsPerPage = 100
		case meta.ItemsPerPage < 10:
			meta.ItemsPerPage = 10
		}

		offset := (meta.CurrentPage - 1) * meta.ItemsPerPage
		return db.Offset(int(offset)).Limit(int(meta.ItemsPerPage))
	}
}

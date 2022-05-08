package repository

import (
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"gorm.io/gorm"
	"math"
)

func Pagination(value interface{}, meta *model.PaginationMetadata, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalItems int64
	db.Model(&value).Count(&totalItems)

	meta.TotalItem = totalItems
	totalPages := math.Ceil(float64(totalItems) / float64(meta.ItemsPerPage))
	meta.TotalPage = int64(totalPages)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(meta.GetOffset())).Limit(int(meta.ItemsPerPage))
	}
}

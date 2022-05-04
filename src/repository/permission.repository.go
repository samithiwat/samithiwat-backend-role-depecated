package repository

import (
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PermissionRepository struct {
	db gorm.DB
}

func NewPermissionRepository(db gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindAll(meta *proto.PaginationMetadata, perms *[]*model.Permission) error {
	return r.db.Scopes(Pagination(perms, meta)).Find(&perms).Count(&meta.ItemCount).Error
}

func (r *PermissionRepository) FindOne(id int, perm *model.Permission) error {
	return r.db.Preload(clause.Associations).First(&perm, id).Error
}

func (r *PermissionRepository) Create(perm *model.Permission) error {
	return r.db.Create(&perm).Error
}

func (r *PermissionRepository) Update(id int, perm *model.Permission) error {
	return r.db.Where(id).Updates(&perm).First(&perm).Error
}

func (r *PermissionRepository) Delete(id int, perm *model.Permission) error {
	return r.db.First(&perm, id).Delete(&model.Permission{}).Error
}

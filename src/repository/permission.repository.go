package repository

import (
	"github.com/fecamp-cu/fecamp-2022-user/src/database"
	"github.com/fecamp-cu/fecamp-2022-user/src/model"
	"github.com/fecamp-cu/fecamp-2022-user/src/proto"
	"gorm.io/gorm/clause"
)

type PermissionRepository struct {
	db database.Database
}

func NewPermissionRepository(db database.Database) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindAll(meta *proto.PaginationMetadata, perms *[]*model.Permission) error {
	return r.db.GetConnection().Scopes(Pagination(perms, meta)).Find(&perms).Count(&meta.ItemCount).Error
}

func (r *PermissionRepository) FindOne(id int, perm *model.Permission) error {
	return r.db.GetConnection().Preload(clause.Associations).First(&perm, id).Error
}

func (r *PermissionRepository) Create(perm *model.Permission) error {
	return r.db.GetConnection().Create(&perm).Error
}

func (r *PermissionRepository) Update(id int, perm *model.Permission) error {
	return r.db.GetConnection().Where(id).Updates(&perm).First(&perm).Error
}

func (r *PermissionRepository) Delete(id int, perm *model.Permission) error {
	return r.db.GetConnection().First(&perm, id).Delete(&model.Permission{}).Error
}

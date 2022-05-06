package repository

import (
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) FindAll(meta *proto.PaginationMetadata, roles *[]*model.Role) error {
	return r.db.Scopes(Pagination(roles, meta)).Find(&roles).Count(&meta.ItemCount).Error
}

func (r *RoleRepository) FindOne(id int, role *model.Role) error {
	return r.db.Preload(clause.Associations).First(&role, id).Error
}

func (r *RoleRepository) FindMulti(ids []uint32, roles *[]*model.Role) error {
	return r.db.Where("id IN ?", ids).Find(&roles).Error
}

func (r *RoleRepository) Create(role *model.Role) error {
	return r.db.Create(&role).Error
}

func (r *RoleRepository) Update(id int, role *model.Role) error {
	r.db.Model(&model.Role{}).Association("Permissions").Append(role.Permissions)
	return r.db.Where(id).Updates(&role).First(&role).Error
}

func (r *RoleRepository) Delete(id int, role *model.Role) error {
	return r.db.First(&role, id).Delete(&model.Role{}).Error
}

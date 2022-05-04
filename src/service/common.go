package service

import (
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
	"gorm.io/gorm"
)

func RawToDtoPermission(perm *model.Permission) *proto.Permission {
	return &proto.Permission{
		Id:   uint32(perm.ID),
		Name: perm.Name,
		Code: perm.Code,
	}
}

func RawToDtoRole(role *model.Role) *proto.Role {
	var permissions []*proto.Permission
	for _, permission := range role.Permissions {
		rolePerm := RawToDtoPermission(permission)
		permissions = append(permissions, rolePerm)
	}
	return &proto.Role{
		Id:          uint32(role.ID),
		Name:        role.Name,
		Description: role.Description,
		Permissions: permissions,
	}
}

func DtoToRawPermission(permission *proto.Permission) *model.Permission {
	return &model.Permission{
		Model: gorm.Model{
			ID: uint(permission.Id),
		},
		Name: permission.Name,
		Code: permission.Code,
	}
}

func DtoToRawRole(role *proto.Role) *model.Role {
	var perms []*model.Permission
	for _, perm := range role.Permissions {
		perms = append(perms, DtoToRawPermission(perm))
	}

	return &model.Role{
		Model:       gorm.Model{ID: uint(role.Id)},
		Name:        role.Name,
		Description: role.Description,
		Permissions: perms,
	}
}

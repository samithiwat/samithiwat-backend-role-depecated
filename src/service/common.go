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

func DtoToRawPermission(permission *proto.Permission) *model.Permission {
	return &model.Permission{
		Model: gorm.Model{
			ID: uint(permission.Id),
		},
		Name: permission.Name,
		Code: permission.Code,
	}
}

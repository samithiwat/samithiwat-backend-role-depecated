package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
	"gorm.io/gorm"
)

var Perm1 model.Permission
var Perm2 model.Permission
var Perm3 model.Permission
var Perm4 model.Permission
var Perms []*model.Permission
var CreatePermissionReqMock proto.CreatePermissionRequest
var UpdatePermissionReqMock proto.UpdatePermissionRequest

type PermissionMockRepo struct {
}

func (r *PermissionMockRepo) FindAll(meta *proto.PaginationMetadata, perms *[]*model.Permission) error {
	meta.CurrentPage = 1
	meta.TotalPage = 1
	meta.ItemCount = 4
	meta.TotalItem = 4
	meta.ItemsPerPage = 10

	*perms = Perms

	return nil
}

func (r *PermissionMockRepo) FindOne(_ int, perm *model.Permission) error {
	*perm = Perm1
	return nil
}

func (r *PermissionMockRepo) Create(perm *model.Permission) error {
	*perm = Perm1
	return nil
}

func (r *PermissionMockRepo) Update(_ int, perm *model.Permission) error {
	*perm = Perm1
	return nil
}

func (r *PermissionMockRepo) Delete(_ int, perm *model.Permission) error {
	*perm = Perm1
	return nil
}

type PermissionMockErrRepo struct {
}

func (r *PermissionMockErrRepo) FindAll(*proto.PaginationMetadata, *[]*model.Permission) error {
	return nil
}

func (r *PermissionMockErrRepo) FindOne(_ int, perm *model.Permission) error {
	return errors.New("Not found permission")
}

func (r *PermissionMockErrRepo) Create(*model.Permission) error {
	return errors.New("Duplicate permission")
}

func (r *PermissionMockErrRepo) Update(_ int, perm *model.Permission) error {
	return errors.New("Not found permission")
}

func (r *PermissionMockErrRepo) Delete(_ int, perm *model.Permission) error {
	return errors.New("Not found permission")
}

func InitializeMockPermission() (err error) {
	Perm1 = model.Permission{
		Model: gorm.Model{ID: 1},
		Name:  faker.Name(),
		Code:  faker.Word(),
	}

	Perm2 = model.Permission{
		Model: gorm.Model{ID: 2},
		Name:  faker.Name(),
		Code:  faker.Word(),
	}

	Perm3 = model.Permission{
		Model: gorm.Model{ID: 3},
		Name:  faker.Name(),
		Code:  faker.Word(),
	}

	Perm4 = model.Permission{
		Model: gorm.Model{ID: 4},
		Name:  faker.Name(),
		Code:  faker.Word(),
	}

	CreatePermissionReqMock = proto.CreatePermissionRequest{
		Permission: &proto.Permission{
			Name: Perm1.Name,
			Code: Perm1.Code,
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	UpdatePermissionReqMock = proto.UpdatePermissionRequest{
		Permission: &proto.Permission{
			Id:   uint32(Perm1.ID),
			Name: Perm1.Name,
			Code: Perm1.Code,
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	Perms = append(Perms, &Perm1, &Perm2, &Perm3, &Perm4)

	return nil
}

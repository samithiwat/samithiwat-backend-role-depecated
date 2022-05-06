package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
)

var Role1 model.Role
var Role2 model.Role
var Role3 model.Role
var Role4 model.Role
var Roles []*model.Role
var CreateRoleReqMock proto.CreateRoleRequest
var UpdateRoleReqMock proto.UpdateRoleRequest

type RoleMockRepo struct {
}

func (r *RoleMockRepo) FindAll(meta *proto.PaginationMetadata, roles *[]*model.Role) error {
	meta.CurrentPage = 1
	meta.TotalPage = 1
	meta.ItemCount = 4
	meta.TotalItem = 4
	meta.ItemsPerPage = 10

	*roles = Roles

	return nil
}

func (r *RoleMockRepo) FindOne(_ int, role *model.Role) error {
	*role = Role1
	return nil
}

func (r *RoleMockRepo) FindMulti(_ []uint32, roles *[]*model.Role) error {
	*roles = Roles
	return nil
}

func (r *RoleMockRepo) Create(role *model.Role) error {
	*role = Role1
	return nil
}

func (r *RoleMockRepo) Update(_ int, role *model.Role) error {
	*role = Role1
	return nil
}

func (r *RoleMockRepo) Delete(_ int, role *model.Role) error {
	*role = Role1
	return nil
}

type RoleMockErrRepo struct {
}

func (r *RoleMockErrRepo) FindAll(*proto.PaginationMetadata, *[]*model.Role) error {
	return nil
}

func (r *RoleMockErrRepo) FindOne(int, *model.Role) error {
	return errors.New("Not found role")
}

func (r *RoleMockErrRepo) FindMulti([]uint32, *[]*model.Role) error {
	return nil
}

func (r *RoleMockErrRepo) Create(*model.Role) error {
	return nil
}

func (r *RoleMockErrRepo) Update(int, *model.Role) error {
	return errors.New("Not found role")
}

func (r *RoleMockErrRepo) Delete(int, *model.Role) error {
	return errors.New("Not found role")
}

func InitializeMockRole() (err error) {
	Role1 = model.Role{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	Role2 = model.Role{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	Role3 = model.Role{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	Role4 = model.Role{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	CreateRoleReqMock = proto.CreateRoleRequest{
		Role: &proto.Role{
			Name:        Role1.Name,
			Description: Role1.Description,
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	UpdateRoleReqMock = proto.UpdateRoleRequest{
		Role: &proto.Role{
			Id:          uint32(Role1.ID),
			Name:        Role1.Name,
			Description: Role1.Description,
			Permissions: []*proto.Permission{{Id: 5}, {Id: 6}},
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	Roles = append(Roles, &Role1, &Role2, &Role3, &Role4)

	return nil
}

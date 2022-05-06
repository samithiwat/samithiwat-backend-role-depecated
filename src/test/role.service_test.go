package test

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
	"github.com/samithiwat/samithiwat-backend-role/src/service"
	"github.com/samithiwat/samithiwat-backend-role/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindAllRole(t *testing.T) {
	mock.InitializeMockRole()

	assert := assert.New(t)

	var result []*proto.Role
	for _, role := range mock.Roles {
		result = append(result, RawToDtoRole(role))
	}

	var errors []string

	want := &proto.RolePaginationResponse{
		Data: &proto.RolePagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    4,
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewRoleService(&mock.RoleMockRepo{})
	roleRes, err := roleService.FindAll(mock.Context{}, &proto.FindAllRoleRequest{Limit: 10, Page: 1})

	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes, fmt.Sprintf("Want %v but got %v", want, roleRes))
}

func TestFindOneRole(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       RawToDtoRole(&mock.Role1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewRoleService(&mock.RoleMockRepo{})
	roleRes, err := roleService.FindOne(mock.Context{}, &proto.FindOneRoleRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}

func TestFindOneErrNotFoundRole(t *testing.T) {
	mock.InitializeMockRole()

	errors := []string{"Not found role"}

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	roleService := service.NewRoleService(&mock.RoleMockErrRepo{})
	roleRes, err := roleService.FindOne(mock.Context{}, &proto.FindOneRoleRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, roleRes)
}

func TestFindMultiRole(t *testing.T) {
	mock.InitializeMockRole()

	var result []*proto.Role
	for _, role := range mock.Roles {
		result = append(result, RawToDtoRole(role))
	}

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleListResponse{
		Data:       result,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewRoleService(&mock.RoleMockRepo{})
	roleRes, err := roleService.FindMulti(mock.Context{}, &proto.FindMultiRoleRequest{Ids: []uint32{1, 2, 3, 4, 5}})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}

func TestCreateRole(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       RawToDtoRole(&mock.Role1),
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	roleService := service.NewRoleService(&mock.RoleMockRepo{})
	roleRes, err := roleService.Create(mock.Context{}, &mock.CreateRoleReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}

func TestUpdateRole(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       RawToDtoRole(&mock.Role1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewRoleService(&mock.RoleMockRepo{})
	roleRes, err := roleService.Update(mock.Context{}, &mock.UpdateRoleReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}

func TestUpdateErrNotFoundRole(t *testing.T) {
	mock.InitializeMockRole()

	errors := []string{"Not found role"}

	assert := assert.New(t)

	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	roleService := service.NewRoleService(&mock.RoleMockErrRepo{})
	roleRes, err := roleService.Update(mock.Context{}, &mock.UpdateRoleReqMock)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, roleRes)
}

func TestDeleteRole(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       RawToDtoRole(&mock.Role1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewRoleService(&mock.RoleMockRepo{})
	roleRes, err := roleService.Delete(mock.Context{}, &proto.DeleteRoleRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}

func TestDeleteErrNotFoundRole(t *testing.T) {
	mock.InitializeMockRole()

	errors := []string{"Not found role"}

	assert := assert.New(t)

	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	roleService := service.NewRoleService(&mock.RoleMockErrRepo{})
	roleRes, err := roleService.Delete(mock.Context{}, &proto.DeleteRoleRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, roleRes)
}

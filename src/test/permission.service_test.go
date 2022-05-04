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

func TestFindAllPermission(t *testing.T) {
	mock.InitializeMockPermission()

	assert := assert.New(t)

	var result []*proto.Permission
	for _, perm := range mock.Perms {
		result = append(result, RawToDtoPermission(perm))
	}

	var errors []string

	want := &proto.PermissionListResponse{
		Data: &proto.PermissionPagination{
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

	permService := service.NewPermissionService(&mock.PermissionMockRepo{})
	permRes, err := permService.FindAll(mock.Context{}, &proto.FindAllPermissionRequest{Limit: 10, Page: 1})

	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, permRes, fmt.Sprintf("Want %v but got %v", want, permRes))
}

func TestFindOnePermission(t *testing.T) {
	mock.InitializeMockPermission()

	var errors []string

	assert := assert.New(t)
	want := &proto.PermissionResponse{
		Data:       RawToDtoPermission(&mock.Perm1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	permService := service.NewPermissionService(&mock.PermissionMockRepo{})
	permRes, err := permService.FindOne(mock.Context{}, &proto.FindOnePermissionRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, permRes)
}

func TestFindOneErrNotFoundPermission(t *testing.T) {
	mock.InitializeMockPermission()

	errors := []string{"Not found permission"}

	assert := assert.New(t)
	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	permService := service.NewPermissionService(&mock.PermissionMockErrRepo{})
	permRes, err := permService.FindOne(mock.Context{}, &proto.FindOnePermissionRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, permRes)
}

func TestCreatePermission(t *testing.T) {
	mock.InitializeMockPermission()

	var errors []string

	assert := assert.New(t)
	want := &proto.PermissionResponse{
		Data:       RawToDtoPermission(&mock.Perm1),
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	permService := service.NewPermissionService(&mock.PermissionMockRepo{})
	permRes, err := permService.Create(mock.Context{}, &mock.CreatePermissionReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, permRes)
}

func TestCreateDuplicatedPermission(t *testing.T) {
	mock.InitializeMockPermission()

	errors := []string{"Duplicate permission"}

	assert := assert.New(t)

	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusUnprocessableEntity,
	}

	permService := service.NewPermissionService(&mock.PermissionMockErrRepo{})
	permRes, err := permService.Create(mock.Context{}, &mock.CreatePermissionReqMock)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, permRes)
}

func TestUpdatePermission(t *testing.T) {
	mock.InitializeMockPermission()

	var errors []string

	assert := assert.New(t)
	want := &proto.PermissionResponse{
		Data:       RawToDtoPermission(&mock.Perm1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	permService := service.NewPermissionService(&mock.PermissionMockRepo{})
	permRes, err := permService.Update(mock.Context{}, &mock.UpdatePermissionReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, permRes)
}

func TestUpdateErrNotFoundPermission(t *testing.T) {
	mock.InitializeMockPermission()

	errors := []string{"Not found permission"}

	assert := assert.New(t)

	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	permService := service.NewPermissionService(&mock.PermissionMockErrRepo{})
	permRes, err := permService.Update(mock.Context{}, &mock.UpdatePermissionReqMock)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, permRes)
}

func TestDeletePermission(t *testing.T) {
	mock.InitializeMockPermission()

	var errors []string

	assert := assert.New(t)
	want := &proto.PermissionResponse{
		Data:       RawToDtoPermission(&mock.Perm1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	permService := service.NewPermissionService(&mock.PermissionMockRepo{})
	permRes, err := permService.Delete(mock.Context{}, &proto.DeletePermissionRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, permRes)
}

func TestDeleteErrNotFoundPermission(t *testing.T) {
	mock.InitializeMockPermission()

	errors := []string{"Not found permission"}

	assert := assert.New(t)

	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	permService := service.NewPermissionService(&mock.PermissionMockErrRepo{})
	permRes, err := permService.Delete(mock.Context{}, &proto.DeletePermissionRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, permRes)
}

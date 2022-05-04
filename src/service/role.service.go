package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"github.com/samithiwat/samithiwat-backend-role/src/proto"
	"net/http"
)

type RoleService struct {
	repository RoleRepository
}

type RoleRepository interface {
	FindAll(*proto.PaginationMetadata, *[]*model.Role) error
	FindOne(int, *model.Role) error
	Create(*model.Role) error
	Update(int, *model.Role) error
	Delete(int, *model.Role) error
}

func NewRoleService(repository RoleRepository) *RoleService {
	return &RoleService{repository: repository}
}

func (s *RoleService) FindAll(_ context.Context, req *proto.FindAllRoleRequest) (res *proto.RoleListResponse, err error) {
	meta := proto.PaginationMetadata{
		ItemsPerPage: req.Limit,
		CurrentPage:  req.Page,
	}

	var roles []*model.Role
	var errors []string

	res = &proto.RoleListResponse{
		Data: &proto.RolePagination{
			Items: nil,
			Meta:  &meta,
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindAll(&meta, &roles)
	if err != nil {
		errors = append(errors, err.Error())
		res.StatusCode = http.StatusBadRequest
		return
	}

	var result []*proto.Role

	for _, role := range roles {
		result = append(result, RawToDtoRole(role))
	}

	res.Data.Items = result

	return
}

func (s *RoleService) FindOne(_ context.Context, req *proto.FindOneRoleRequest) (res *proto.RoleResponse, err error) {
	role := model.Role{}
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(int(req.Id), &role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoRole(&role)
	res.Data = result

	return
}

func (s *RoleService) Create(_ context.Context, req *proto.CreateRoleRequest) (res *proto.RoleResponse, err error) {
	role := DtoToRawRole(req.Role)
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return
	}

	result := RawToDtoRole(role)
	res.Data = result

	return
}

func (s *RoleService) Update(_ context.Context, req *proto.UpdateRoleRequest) (res *proto.RoleResponse, err error) {
	role := DtoToRawRole(req.Role)
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(int(role.ID), role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoRole(role)
	res.Data = result

	return
}

func (s *RoleService) Delete(_ context.Context, req *proto.DeleteRoleRequest) (res *proto.RoleResponse, err error) {
	role := model.Role{}
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(int(req.Id), &role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoRole(&role)
	res.Data = result

	return
}

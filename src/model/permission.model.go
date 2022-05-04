package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string `json:"name"`
	Code  string `json:"code" gorm:"index:,unique"`
	Roles uint   `json:"roles" gorm:"many2many:role_permission"`
}

package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserName string `json:"username"`
	PassWord bool `json:"password"`
}

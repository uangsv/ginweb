package models

import (
	"gorm.io/gorm"
	"ginweb/dao"
)

type MarkInfo struct {
	gorm.Model
	Title string `json:"title"`
	Url string `json:"url"`
	Remark string `json:"remark"`
}

func GetAllMarkInfo() (marklist []*MarkInfo, err error) {
	if err := dao.DB.Find(&marklist).Error; err != nil {
		return nil, err
	}
	return
}

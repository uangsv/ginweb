package models

import (
	"gorm.io/gorm"
	"ginweb/dao"
)

type FileInfo struct {
	gorm.Model
	Title string `json:"title"`
	Url string `json:"url"`
	Remark string `json:"remark"`
}


func GetAllFileInfo() (filelist []*FileInfo, err error) {
	if err := dao.DB.Find(&filelist).Error; err != nil {
		return nil, err
	}
	return
}

package controller

import (
	"encoding/json"
	"fmt"
	"ginweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetData(c *gin.Context) {
	dir, _ := os.Getwd()
	filepath := dir+"/static/data/data.json"
	filePtr, err := os.Open(filepath);
	if err != nil{
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)

	datainfo:=make(map[string]interface{})
	err = decoder.Decode(&datainfo)
	if err != nil {
		fmt.Printf("json decode has error:%v\n", err)
	}
	c.JSON(http.StatusOK, datainfo)
}

func GetFileInfo(c *gin.Context) {
	filelist, err := models.GetAllFileInfo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, filelist)
	}
}

func GetDocInfo(c *gin.Context) {
	marklist, err := models.GetAllMarkInfo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, marklist)
	}
}

func UserRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": "UserRegister",
	})
}

func UserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": "UserLogin",
	})
}

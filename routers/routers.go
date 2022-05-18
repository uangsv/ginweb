package routers

import (
	"ginweb/setting"

	"github.com/gin-gonic/gin"

	"ginweb/controller"
	"ginweb/middlewares"
)

func SetupRouter() *gin.Engine {
	// DebugMode indicates gin mode is debug.
	//DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	//ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	//TestMode = "test"

	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.Use(middlewares.Cors())//使用中间件
	// 静态 模板路径
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")

	r.GET("/", controller.IndexHandler)
	r.GET("/data", controller.GetData)

	// 文件
	fGroup := r.Group("/file")
	{
		// 查看所有的待办事项
		fGroup.GET("/info", controller.GetFileInfo)
	}

	// 文档
	dGroup := r.Group("/doc")
	{
		// 查看所有的待办事项
		dGroup.GET("/info", controller.GetDocInfo)
	}

	// 用户
	uGroup := r.Group("/user")
	{
		uGroup.POST("/register", controller.UserRegister)
		uGroup.POST("/login", controller.UserLogin)
		//uGroup.GET("/login", controller.UserLogin)
		//uGroup.DELETE("/login", controller.UserLogin)
	}

	return r
}

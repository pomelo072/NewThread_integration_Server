package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"integration_server/config"
	"integration_server/handlers"
	"integration_server/router"
)

func main() {
	app := gin.Default()
	// 跨域中间件
	app.Use(cors.Default())
	// 校验中间件
	app.Use(handlers.JwtVerity)
	// 用户组
	user := app.Group("/api/user")
	router.UserGroup(user)
	// 管理员组
	admin := app.Group("/api/admin")
	router.AdminGroup(admin)
	// 积分组
	integration := app.Group("/api/integration")
	router.IntegrationGroup(integration)
	// 奖品组
	award := app.Group("/api/exchange")
	router.AwardGroup(award)
	// 图片接口
	app.GET("/api/img/:imgPath", handlers.DownloadIMG)
	// 启动服务
	app.Run(config.Sysconfig.Port)
}

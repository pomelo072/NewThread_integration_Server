package router

import "github.com/gin-gonic/gin"

// UserGroup 用户路由组
func UserGroup(group *gin.RouterGroup) {
	// 登录接口
	group.POST("/login", controllers.UserLogin)
	// 个人信息接口
	group.GET("/info", controllers.UserInfo)
	// 修改个人信息接口
	group.POST("/edit", controllers.UserEdit)
	group.GET("/gethistory", controllers.GetIntegrationDetail)
	group.GET("/giftshistory", controllers.GetAwardDetail)
}

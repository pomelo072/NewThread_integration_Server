package router

import "github.com/gin-gonic/gin"

// IntegrationGroup 积分路由组
func IntegrationGroup(group *gin.RouterGroup) {
	group.POST("/apply", controllers.IntegrationApply)
	group.GET("/delapply", controllers.IntegrationApplyDel)
}

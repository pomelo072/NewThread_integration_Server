package router

import "github.com/gin-gonic/gin"

// AwardGroup 奖品路由组
func AwardGroup(group *gin.RouterGroup) {
	group.GET("/info", controllers.AwardInfo)
	group.GET("/list", controllers.AwardList)
	group.GET("/apply", controllers.AwardApply)
}

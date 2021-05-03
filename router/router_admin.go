package router

import "github.com/gin-gonic/gin"

// AdminGroup 管理路由组
func AdminGroup(group *gin.RouterGroup) {
	adminUser := group.Group("/user")
	AdminUserGroup(adminUser)
	adminAward := group.Group("/award")
	AdminAwardGroup(adminAward)
	awardDetail := group.Group("/adetail")
	AdminAwardDetailGroup(awardDetail)
	integrationDetail := group.Group("/idetail")
	AdminIntegrationDetailGroup(integrationDetail)
	integrationApply := group.Group("/iapply")
	AdminIntegrationApplyGroup(integrationApply)
}

// AdminUserGroup 用户管理子路由组
func AdminUserGroup(adminUser *gin.RouterGroup) {
	adminUser.POST("/add", controllers.AdminUserAdd)
	adminUser.POST("/edit", controllers.AdminUserEdit)
	adminUser.GET("/delete", controllers.AdminUserDel)
}

// AdminAwardGroup 奖品管理子路由组
func AdminAwardGroup(adminAward *gin.RouterGroup) {
	adminAward.POST("/add", controllers.AdminAwardAdd)
	adminAward.POST("/edit", controllers.AdminAwardEdit)
	adminAward.GET("/delete", controllers.AdminAwardDel)
}

// AdminAwardDetailGroup 奖品兑换明细子路由组
func AdminAwardDetailGroup(awardDetail *gin.RouterGroup) {
	awardDetail.GET("/download", controllers.AwardDetailDown)
}

// AdminIntegrationDetailGroup 积分兑换明细子路由组
func AdminIntegrationDetailGroup(integrationDetail *gin.RouterGroup) {
	integrationDetail.GET("/download", controllers.IntegrationDetailDown)
}

// AdminIntegrationApplyGroup 积分审核子路由组
func AdminIntegrationApplyGroup(integrationApply *gin.RouterGroup) {
	integrationApply.GET("/list", controllers.IntegrationApplyList)
	integrationApply.GET("/status", controllers.IntegrationApplyStatus)
}

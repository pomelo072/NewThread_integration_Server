package controllers

import (
	"github.com/gin-gonic/gin"
	"integration_server/handlers"
	"integration_server/json_struct"
	"integration_server/utils"
	"net/http"
)

// IntegrationApply 申请积分
func IntegrationApply(ctx *gin.Context) {
	// 获取用户信息
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 模型生成绑定
	if userclaims.Auth == 0 {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	} else {
		var applyInfo json_struct.IntegrationModel
		// 不允许字段为空
		if err := ctx.BindJSON(&applyInfo); err != nil {
			ctx.JSON(http.StatusNoContent, utils.GetReturnData(gin.H{"message": "Something Empty."}, "ERROR"))
			return
		}
		// 上传Handler
		result := handlers.ApplyIntegration(&applyInfo, userclaims)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	}
}

// IntegrationApplyDel 删除积分申请
func IntegrationApplyDel(ctx *gin.Context) {
	// 获取用户信息和申请记录ID
	user, _ := ctx.Get("user")
	id := ctx.Query("id")
	userclaims := user.(*json_struct.UserClaims)
	// Admin拒绝访问
	if userclaims.Auth == 0 {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	} else {
		result := handlers.DelApplyIntegration(id, userclaims)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	}
}

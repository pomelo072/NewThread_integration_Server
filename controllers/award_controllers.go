package controllers

import (
	"github.com/gin-gonic/gin"
	"integration_server/handlers"
	"integration_server/json_struct"
	"integration_server/utils"
	"net/http"
	"strconv"
)

// AwardInfo 获取奖品信息
func AwardInfo(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	numid, _ := strconv.Atoi(id)
	if numid == 0 {
		// 获取不到id
		ctx.JSON(http.StatusNoContent, utils.GetReturnData(gin.H{"message": "id Empty"}, "ERROR"))
	} else {
		// Handler查询
		result := handlers.InfoAward(numid)
		ctx.JSON(http.StatusOK, utils.GetReturnData(result, "SUCCESS"))
	}
}

// AwardList 获取奖品List
func AwardList(ctx *gin.Context) {
	tp := ctx.DefaultQuery("type", "all")
	result := handlers.ListAward(tp)
	ctx.JSON(http.StatusOK, utils.GetReturnData(result, "SUCCESS"))
}

// AwardApply 申请奖品
func AwardApply(ctx *gin.Context) {
	// 获取用户信息和奖品ID
	user, _ := ctx.Get("user")
	id := ctx.DefaultQuery("id", "0")
	userclaims := user.(*json_struct.UserClaims)
	// Admin拒绝访问
	if userclaims.Auth == 0 {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	} else if id == "0" {
		// 缺少奖品ID值
		ctx.JSON(http.StatusNoContent, utils.GetReturnData(gin.H{"message": "ID Empty"}, "ERROR"))
	} else {
		// 数据库操作
		// 返回码:
		// 200 兑换成功
		// 400 积分不足
		// 401 重复兑换
		// 500 库存不足
		result := handlers.ApplyAward(id, userclaims)
		if result == 200 {
			ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": "Apply Successfully"}, "SUCCESS"))
		} else if result == 500 {
			ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": "No More Stock"}, "ERROR"))
		} else if result == 400 {
			ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": "Need More Integration"}, "ERROR"))
		} else if result == 401 {
			ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": "Repeatedd change."}, "ERROR"))
		}
	}
}

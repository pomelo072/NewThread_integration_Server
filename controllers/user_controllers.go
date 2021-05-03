package controllers

import (
	"github.com/gin-gonic/gin"
	"integration_server/handlers"
	"integration_server/json_struct"
	"integration_server/utils"
	"net/http"
)

// UserLogin 登录接口
func UserLogin(ctx *gin.Context) {
	var userinfo json_struct.UserVerify
	// 获取登录信息JSON
	if err := ctx.ShouldBindJSON(&userinfo); err != nil {
		// 信息缺失报错
		ctx.JSON(http.StatusUnauthorized, utils.GetReturnData(gin.H{
			"message": err.Error(),
		}, "ERROR"))
		return
	}
	// 查询用户信息
	result := handlers.LoginUser(userinfo)
	// 校验信息成功则返回token
	if result != "error" || result != "UID error" {
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"token": result}, "SUCCESS"))
	} else {
		ctx.JSON(http.StatusUnauthorized, utils.GetReturnData(gin.H{"message": result}, "ERROR"))
	}
}

// UserInfo 获取用户信息
func UserInfo(ctx *gin.Context) {
	// 获取token
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 查询信息
	result := handlers.InfoUser(userclaims)
	// 返回信息
	ctx.JSON(http.StatusOK, result)
}

// UserEdit 用户信息更新接口
func UserEdit(ctx *gin.Context) {
	// 获取UID
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	var userEditor json_struct.UserEditor
	// 查询绑定的模型
	if ctx.BindJSON(&userEditor) == nil {
		ctx.JSON(http.StatusNoContent, utils.GetReturnData(gin.H{"message": "empty content"}, "ERROR"))
		return
	}
	// 调用处理程序
	result := handlers.EditUser(userclaims, &userEditor)
	ctx.JSON(http.StatusOK, result)
}

// GetIntegrationDetail 获取个人积分记录
func GetIntegrationDetail(ctx *gin.Context) {
	// 获取token
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// Admin拒绝访问
	if userclaims.Auth == 0 {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
		return
	}
	// 查询及返回
	result := handlers.PersonIntegrationDetail(userclaims)
	ctx.JSON(http.StatusOK, utils.GetReturnData(result, "SUCCESS"))
}

// GetAwardDetail 获取个人礼物兑换记录
func GetAwardDetail(ctx *gin.Context) {
	// 获取token
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// Admin拒绝访问
	if userclaims.Auth == 0 {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
		return
	}
	// 查询及返回
	result := handlers.PersonAwardDetail(userclaims)
	ctx.JSON(http.StatusOK, utils.GetReturnData(result, "SUCCESS"))
}

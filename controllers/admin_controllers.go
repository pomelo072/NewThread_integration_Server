package controllers

import (
	"github.com/gin-gonic/gin"
	"integration_server/handlers"
	"integration_server/json_struct"
	"integration_server/utils"
	"net/http"
)

// AdminUserAdd 新增用户
func AdminUserAdd(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	var group json_struct.UserAdd
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取JSON
		if ctx.ShouldBindJSON(&group) != nil {
			result := handlers.UserAddAdmin(&group)
			ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
		}
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AdminUserEdit 用户编辑
func AdminUserEdit(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取UID和新PassWord
		uid := ctx.Query("uid")
		psd := ctx.Query("password")
		result := handlers.UserEditAdmin(uid, psd)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AdminUserDel 删除用户
func AdminUserDel(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取UID
		uid := ctx.Query("uid")
		result := handlers.UserDelAdmin(uid)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AdminAwardAdd 奖品新增
func AdminAwardAdd(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取JSON
		var award json_struct.Award
		if ctx.ShouldBindJSON(&award) != nil {
			ctx.JSON(http.StatusNoContent, utils.GetReturnData(gin.H{"message": "Empty Content"}, "ERROR"))
			return
		}
		result := handlers.AwardAddAdmin(&award)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AdminAwardEdit 奖品编辑
func AdminAwardEdit(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取UID
		id := ctx.Query("id")
		var award json_struct.Award
		if ctx.ShouldBindJSON(&award) != nil {
			ctx.JSON(http.StatusNoContent, utils.GetReturnData(gin.H{"message": "Empty Content"}, "ERROR"))
			return
		}
		result := handlers.AwardEditAdmin(id, &award)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AdminAwardStatus 修改奖品状态
func AdminAwardStatus(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取UID和新PassWord
		id := ctx.Param("id")
		result := handlers.AwardStatusAdmin(id)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AdminAwardList 获取审核奖品列表
func AdminAwardList(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取List
		result := handlers.AwardListAdmin()
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AdminAwardDel 奖品删除
func AdminAwardDel(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取ID
		id := ctx.Query("id")
		result := handlers.AwardDelAdmin(id)
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": result}, "SUCCESS"))
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// IntegrationApplyList 获取审核列表
func IntegrationApplyList(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取List
		result := handlers.ApplyListIntegration()
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// IntegrationApplyStatus 审核积分状态
func IntegrationApplyStatus(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取ID和审核状态
		id := ctx.Query("id")
		status := ctx.Query("status")
		text := ctx.Query("text")
		result := handlers.ApplyStatusIntegration(id, status, text)
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// AwardDetailsDown 获取近三年奖品兑换记录下载
func AwardDetailsDown(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取List
		result := handlers.DownAwardDetails()
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

// IntegrationDetailsDown 获取近三年积分明细记录下载
func IntegrationDetailsDown(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userclaims := user.(*json_struct.UserClaims)
	// 管理组鉴权
	if userclaims.Auth == 0 {
		// 获取List
		result := handlers.DownIntegrationDetails()
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusForbidden, utils.GetReturnData(gin.H{"message": "Access Defined"}, "ERROR"))
	}
}

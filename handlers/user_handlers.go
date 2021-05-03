package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"integration_server/database"
	"integration_server/json_struct"
	"integration_server/models"
	"integration_server/utils"
)

// LoginUser 用户登录handler
func LoginUser(info json_struct.UserVerify) string {
	verify := new(models.UserInfo)
	// 根据uid查询用户信息
	count := database.Db.First(&verify, "uid = ?", info.Uid).RowsAffected
	// 校验
	if verify.UserVerification == utils.GetSHAEncode(info.Verification) {
		// 生成Token
		token := GenerateToken(&json_struct.UserClaims{
			Uid:            verify.UID,
			Verification:   verify.UserVerification,
			Auth:           verify.AUTH,
			StandardClaims: jwt.StandardClaims{},
		})
		return token
	} else if count == 0 {
		// 查询为空
		return "UID error"
	} else {
		// 查询错误
		return "error"
	}
}

// InfoUser 获取用户信息handler
func InfoUser(userclaims *json_struct.UserClaims) *gin.H {
	info := new(models.UserInfo)
	auth := userclaims.Auth
	// 查询信息
	if auth == 1 {
		database.Db.Select("username", "usergroup", "userhead_url",
			"userintegration", "teamintegration").First(&info, "uid = ?", userclaims.Uid)
		// 打包信息返回
		return utils.GetReturnData(&gin.H{
			"name":       info.UserName,
			"identity":   info.UserGroup,
			"points":     info.UserIntegration,
			"teampoints": info.TeamIntegration,
			"auth":       auth,
		}, "SUCCESS")
	} else {
		database.Db.Select("username", "userhead_url").First(&info, "uid = ?", userclaims.Uid)
		// 打包信息返回
		return utils.GetReturnData(&gin.H{
			"name":     info.UserName,
			"identity": info.UserGroup,
			"auth":     auth,
		}, "SUCCESS")
	}
}

// EditUser 用户修改信息handler
func EditUser(userclaims *json_struct.UserClaims, editor *json_struct.UserEditor) *gin.H {
	// 获取模型数据和uid
	uid := userclaims.Uid
	updatePwd := editor.UserPassword
	user := new(models.UserInfo)
	// 查询数据库并修改
	database.Db.First(&user, "uid = ?", uid)
	database.Db.Model(&user).Select("user_verification").Updates(&models.UserInfo{UserVerification: utils.GetSHAEncode(updatePwd)})
	// 返回结果
	return utils.GetReturnData(gin.H{"msg": "update successfully"}, "SUUCCESS")
}

// PersonIntegrationDetail 查询个人积分纪录
func PersonIntegrationDetail(userclaims *json_struct.UserClaims) *gin.H {
	uid := userclaims.Uid
	// 查询UID下的个人积分
	var PersonIntegration []map[string]interface{}
	database.Db.Model(&models.UserInfo{}).Where("uid = ?", uid).Find(&PersonIntegration)
	return &gin.H{"Data": PersonIntegration}
}

// PersonAwardDetail 查询个人兑换记录
func PersonAwardDetail(userclaims *json_struct.UserClaims) *gin.H {
	uid := userclaims.Uid
	// 查询UID下的个人兑换
	var PersonAward []map[string]interface{}
	database.Db.Model(&models.UserInfo{}).Where("uid = ?", uid).Find(&PersonAward)
	return &gin.H{"Data": PersonAward}
}

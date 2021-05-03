package handlers

import (
	"integration_server/database"
	"integration_server/json_struct"
	"integration_server/models"
	"strconv"
	"time"
)

// ApplyIntegration 申请积分
func ApplyIntegration(applyinfo *json_struct.IntegrationModel, userclaims *json_struct.UserClaims) string {
	// 获取UID
	uid := userclaims.Uid
	// 限定申请次数判断
	// 一周内
	timeFlag := time.Now().AddDate(0, 0, -7)
	// 搜索一周内申请次数
	list := new(models.IntegrationApply)
	// 判断是否超过10次
	if count := database.Db.Where("created_at > ?", timeFlag).Find(&list).RowsAffected; count > 10 {
		return "Apply Limited."
	}
	// 生成记录
	database.Db.Create(&models.IntegrationApply{
		IntegrationType: applyinfo.IntegrationType,
		ApplyUID:        uid,
		ApplyText:       applyinfo.ApplyText,
		ApplyType:       applyinfo.ApplyType,
		ApplyLevel:      applyinfo.ApplyLevel,
		ContestLevel:    applyinfo.ContestLevel,
		ApplyNumber:     0,
		ResponseReason:  "",
		ApplyStatus:     0,
	})
	// 获取申请ID
	newaward := new(models.IntegrationApply)
	database.Db.Model(&models.IntegrationApply{}).Where("uid = ?", uid).Last(&newaward)
	// 解包图片入库
	err := UnpackIMG(applyinfo.ApplyIMG, "积分", newaward.ID)
	if err != "SUCCESS" {
		return "Image Upload Error."
	}
	return "Upload successfully."
}

// DelApplyIntegration 删除申请记录
func DelApplyIntegration(id string, userclaims *json_struct.UserClaims) string {
	// 获取申请信息
	ApplyID, _ := strconv.Atoi(id)
	ApplyInfo := new(models.IntegrationApply)
	database.Db.Model(&models.IntegrationApply{}).Where("id = ? AND apply_uid = ?", ApplyID, userclaims.Uid).First(&ApplyInfo)
	// 判断申请是否已经通过
	if ApplyInfo.ApplyStatus == 1 {
		return "Passed."
	} else {
		// Apply_Status = 0 可以删除
		database.Db.Delete(&ApplyInfo)
		return "Delete successfully."
	}
}

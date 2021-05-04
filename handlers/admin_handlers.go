package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"integration_server/database"
	"integration_server/json_struct"
	"integration_server/models"
	"integration_server/utils"
	"strconv"
	"time"
)

// UserAddAdmin 新增用户Handler
func UserAddAdmin(user *json_struct.User) string {
	// 遍历JSON
	database.Db.Create(&models.UserInfo{
		UID:              user.UID,
		UserGroup:        user.Group,
		UserIntegration:  0,
		TeamIntegration:  0,
		UserVerification: utils.GetSHAEncode(utils.GetSHAEncode(user.UID)),
		AUTH:             1,
	})
	return "Add User successfully."
}

// UserEditAdmin 用户编辑Handler
func UserEditAdmin(uid string, psd string) string {
	// 数据库更新
	if err := database.Db.Model(&models.UserInfo{}).Where("uid = ?", uid).Update("user_verification", utils.GetSHAEncode(psd)).Error; err != nil {
		return "Edit User Error."
	}
	return "Edit User Successfully."
}

// UserDelAdmin 用户删除Handler
func UserDelAdmin(uid string) string {
	// 数据库删除
	if err := database.Db.Delete(&models.UserInfo{}, uid).Error; err != nil {
		return "Delete User Error."
	}
	return "Delete User successfully."
}

// AwardAddAdmin 奖品信息新增Handler
func AwardAddAdmin(awardInfo *json_struct.Award) string {
	// 根据模型和传入JSON建立记录
	if err := database.Db.Create(&models.AwardInfo{
		AwardType:         awardInfo.AwardType,
		AwardMenu:         awardInfo.AwardMenu,
		AwardName:         awardInfo.AwardName,
		AwardIntroduction: awardInfo.AwardIntroduction,
		NeedIntegration:   awardInfo.NeedIntegration,
		PeriodNumber:      1,
		InStock:           awardInfo.InStock,
		UsedNumber:        awardInfo.UsedNumber,
	}).Error; err != nil {
		return "Add Award Error."
	}
	// 获取新奖品ID
	newaward := new(models.AwardInfo)
	database.Db.Model(&models.AwardInfo{}).Where("award_name = ? AND award_type = ? AND award_menu = ?", awardInfo.AwardName, awardInfo.AwardType, awardInfo.AwardMenu).Last(&newaward)
	// 解包写入图片
	err := UnpackIMG(awardInfo.AwardIMG, "奖品", newaward.ID)
	if err != "SUCCESS" {
		return "Image Upload Error."
	}
	return "Add Award Successfully."
}

// AwardEditAdmin 奖品信息修改Handler
func AwardEditAdmin(id string, awardInfo *json_struct.Award) string {
	awardID, _ := strconv.Atoi(id)
	// 获取期数
	AwardPeriod := new(models.AwardInfo)
	database.Db.Model(&models.AwardInfo{}).Where("id = ?", awardID).Last(&AwardPeriod)
	// 数据库更新
	if err := database.Db.Model(&models.AwardInfo{}).Where("id = ?", awardID).Updates(models.AwardInfo{
		AwardType:         awardInfo.AwardType,
		AwardName:         awardInfo.AwardName,
		AwardIntroduction: awardInfo.AwardIntroduction,
		PeriodNumber:      AwardPeriod.PeriodNumber + 1,
		NeedIntegration:   awardInfo.NeedIntegration,
		InStock:           awardInfo.InStock,
		UsedNumber:        awardInfo.UsedNumber,
	}).Error; err != nil {
		return "Edit Award Error."
	}
	return "Edit Award Successfully."
}

// AwardStatusAdmin 奖品状态修改Handler
func AwardStatusAdmin(id string) string {
	detailID, _ := strconv.Atoi(id)
	if err := database.Db.Model(&models.AwardDetail{}).Where("id = ?", detailID).Update("award_status", 1).Error; err != nil {
		return "Status Edit Failed."
	}
	return "Status Edit Successfully."
}

// AwardListAdmin 获取审批奖品列表Handler
func AwardListAdmin() interface{} {
	var list []map[string]interface{}
	if err := database.Db.Model(&models.AwardDetail{}).Where("award_status", 0).Find(&list).Error; err != nil {
		return utils.GetReturnData(&gin.H{"message": "Get AwardDetails List Error."}, "ERROR")
	}
	return utils.GetReturnData(list, "SUCCESS")
}

// AwardDelAdmin 奖品信息删除Handler
func AwardDelAdmin(id string) string {
	awradID, _ := strconv.Atoi(id)
	// 数据库删除
	if err := database.Db.Delete(&models.AwardInfo{}, awradID).Error; err != nil {
		return "Delete Award Error."
	}
	return "Delete Award successfully."
}

// ApplyListIntegration 积分审核List Handler
func ApplyListIntegration() interface{} {
	// 开放模型
	var list []map[string]interface{}
	// 数据库查询, Status为0的数据
	if err := database.Db.Model(&models.IntegrationApply{}).Where("apply_status = ?", 0).Find(&list).Error; err != nil {
		return utils.GetReturnData(&gin.H{"message": "Get Apply List Error."}, "ERROR")
	}
	for _, v := range list {
		v["apply_img"] = ImportIMG("积分", v["id"], 0)
	}
	return utils.GetReturnData(list, "SUCCESS")
}

// ApplyStatusIntegration 审核积分状态Handler
func ApplyStatusIntegration(id string, status string, text string) interface{} {
	// 审核记录查找
	applyID, _ := strconv.Atoi(id)
	applyInfo := new(models.IntegrationApply)
	// 查询记录
	if err := database.Db.Model(&models.IntegrationApply{}).Where("id = ?", applyID).First(&applyInfo).Error; err != nil {
		return utils.GetReturnData(gin.H{"message": "NoRecord."}, "ERROR")
	}
	// 如果已经审核, 则不再审核
	if applyInfo.ApplyStatus != 0 {
		return utils.GetReturnData(gin.H{"message": "Checked."}, "ERROR")
	}
	// 不批准通过
	if status == "2" {
		s, _ := strconv.Atoi(status)
		// 更新Apply_Status为2
		database.Db.Model(&applyInfo).Updates(models.IntegrationApply{ResponseReason: text, ApplyStatus: s})
		return utils.GetReturnData(gin.H{"message": "Failed Successfully."}, "SUCCESS")
	}
	// 允许通过
	if status == "1" {
		s, _ := strconv.Atoi(status)
		integration, _ := strconv.Atoi(text)
		database.Db.Transaction(func(tx *gorm.DB) error {
			// 更新Apply_Status为1, 申请积分数值
			if err := tx.Model(&applyInfo).Updates(models.IntegrationApply{
				ApplyNumber: integration,
				ApplyStatus: s,
			}).Error; err != nil {
				return err
			}
			// 用户信息更新
			user := new(models.UserInfo)
			if err := tx.Model(&models.UserInfo{}).Where("uid = ?", applyInfo.ApplyUID).First(&user).Error; err != nil {
				return err
			}
			// 新增积分记录
			if applyInfo.IntegrationType == 1 {
				// 个人积分记录
				if err := tx.Model(&user).Update("user_integration", user.UserIntegration+integration).Error; err != nil {
					return err
				}
				detail := models.IntegrationDetail{
					IntegrationType: applyInfo.IntegrationType,
					IntegrationUID:  applyInfo.ApplyUID,
					OperationDetail: applyInfo.ApplyText,
					OperationType:   "积分申请",
					OperationNumber: integration,
					AfterOperation:  user.UserIntegration + integration,
				}
				if err := tx.Create(&detail).Error; err != nil {
					return err
				}
			} else if applyInfo.IntegrationType == 0 {
				// 集体积分记录
				if err := tx.Model(&user).Update("team_integration", user.TeamIntegration+integration).Error; err != nil {
					return err
				}
				detail := models.IntegrationDetail{
					IntegrationType: applyInfo.IntegrationType,
					IntegrationUID:  applyInfo.ApplyUID,
					OperationDetail: applyInfo.ApplyText,
					OperationType:   "积分申请",
					OperationNumber: integration,
					AfterOperation:  user.TeamIntegration + integration,
				}
				if err := tx.Create(&detail).Error; err != nil {
					return err
				}
			}
			return nil
		})
	}
	return utils.GetReturnData(gin.H{"message": "Passed."}, "SUCCESS")
}

// DownAwardDetails 获取近三年奖品兑换记录Handler
func DownAwardDetails() interface{} {
	today := time.Now()
	timeFlag := today.AddDate(-3, 0, 0)
	var list []map[string]interface{}
	if err := database.Db.Model(&models.IntegrationDetail{}).Where("created_at > ?", timeFlag).Find(&list).Error; err != nil {
		return utils.GetReturnData(&gin.H{"message": "Download IntegrationDetails List Error."}, "ERROR")
	}
	return utils.GetReturnData(list, "SUCCESS")
}

// DownIntegrationDetails 获取近三年积分明细记录Handler
func DownIntegrationDetails() interface{} {
	today := time.Now()
	timeFlag := today.AddDate(-3, 0, 0)
	var list []map[string]interface{}
	if err := database.Db.Model(&models.AwardDetail{}).Where("created_at > ?", timeFlag).Find(&list).Error; err != nil {
		return utils.GetReturnData(&gin.H{"message": "Download AwardDetails List Error."}, "ERROR")
	}
	return utils.GetReturnData(list, "SUCCESS")
}

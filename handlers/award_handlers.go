package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"integration_server/database"
	"integration_server/json_struct"
	"integration_server/models"
	"strconv"
)

// InfoAward 获取奖品详细信息和推荐
func InfoAward(id int) *gin.H {
	// 定义推荐商品模型
	var recommendGoods1, recommendGoods2 models.AwardInfo
	// 数据库模型
	info := new(models.AwardInfo)
	recommend := new(models.AwardInfo)
	// 查询对应奖品
	database.Db.Model(&info).Where("id = ?", id).First(&info)
	// 查询加入图片
	infoMap := gin.H{
		"name":    info.AwardName,
		"desc":    info.AwardIntroduction,
		"type":    info.AwardMenu,
		"points":  info.NeedIntegration,
		"stock":   info.InStock,
		"imgList": ImportIMG("奖品", id, 0),
	}
	// 查询按库存排序, 按时间排序
	database.Db.Model(&recommend).Order("instock desc").Limit(1).Find(&recommendGoods1)
	database.Db.Model(&recommend).Order("created_at desc").Limit(1).Find(&recommendGoods2)
	// 查询加入图片
	var recommendGoods []interface{}
	recommendGoods = append(recommendGoods, gin.H{
		"id":     recommendGoods1.ID,
		"img":    ImportIMG("奖品", recommendGoods1.ID, 1),
		"title":  recommendGoods1.AwardName,
		"desc":   recommendGoods1.AwardIntroduction,
		"points": recommendGoods1.NeedIntegration,
		"type":   recommendGoods1.AwardMenu,
		"stock":  recommendGoods1.InStock,
	}, gin.H{
		"id":     recommendGoods2.ID,
		"img":    ImportIMG("奖品", recommendGoods2.ID, 1),
		"title":  recommendGoods2.AwardName,
		"desc":   recommendGoods2.AwardIntroduction,
		"points": recommendGoods2.NeedIntegration,
		"type":   recommendGoods2.AwardMenu,
		"stock":  recommendGoods2.InStock,
	})
	return &gin.H{"goodInfo": infoMap, "recommendGoods": recommendGoods}
}

// ListAward 获取奖品列表
func ListAward(tp string) interface{} {
	// 定义开放模型
	var list []map[string]interface{}
	award := new(models.AwardInfo)
	// 获取数据
	if tp == "all" {
		database.Db.Model(&models.AwardInfo{}).Find(&list)
	} else {
		database.Db.Model(&award).Where("award_menu", tp).Find(&list)
	}
	for _, v := range list {
		v["img"] = ImportIMG("奖品", v["id"], 1)
	}
	return list
}

// ApplyAward 申请兑换奖品
func ApplyAward(id string, userclaims *json_struct.UserClaims) int {
	// 查询对应用户数据和奖品数据和兑换记录查询
	awardId, _ := strconv.Atoi(id)
	userinfo := new(models.UserInfo)
	awardinfo := new(models.AwardInfo)
	database.Db.Model(&userinfo).Where("uid = ?", userclaims.Uid).Select("userintegration", "teamintegration").First(&userinfo)
	database.Db.Model(&awardinfo).Where("id = ?", awardId).First(&awardinfo)
	var change []map[string]interface{}
	changeCount := database.Db.Model(&models.AwardDetail{}).Where("award_id = ? AND apply_period = ?", awardId, awardinfo.PeriodNumber).Find(&change).RowsAffected
	// 库存为0判断
	if awardinfo.InStock == 0 {
		return 500
	}
	// 积分数值判断
	var uidtype int
	if awardinfo.AwardType == "个人" {
		uidtype = 1
		if userinfo.UserIntegration < awardinfo.NeedIntegration {
			return 400
		}
	} else {
		uidtype = 0
		if userinfo.TeamIntegration < awardinfo.NeedIntegration {
			return 400
		}
	}
	// 重复兑换判断
	if changeCount > 0 {
		return 401
	}
	// 开始事务
	database.Db.Transaction(func(tx *gorm.DB) error {
		// 扣减库存
		if err := tx.Model(&models.AwardInfo{}).Where("id = ?", awardId).Updates(models.AwardInfo{InStock: awardinfo.InStock - 1, UsedNumber: awardinfo.UsedNumber + 1}).Error; err != nil {
			return err
		}
		// 扣减相应积分
		var afterNumber int
		if awardinfo.AwardType == "个人" {
			afterNumber = userinfo.UserIntegration - awardinfo.NeedIntegration
			if err := tx.Model(&models.UserInfo{}).Where("uid = ?", userclaims.Uid).Update("userintegration", userinfo.UserIntegration-awardinfo.NeedIntegration).Error; err != nil {
				return err
			}
		} else {
			afterNumber = userinfo.TeamIntegration - awardinfo.NeedIntegration
			if err := tx.Model(&models.UserInfo{}).Where("uid = ?", userclaims.Uid).Update("teamintegration", userinfo.TeamIntegration-awardinfo.NeedIntegration).Error; err != nil {
				return err
			}
		}
		// 创建积分变更记录
		if err := tx.Create(&models.IntegrationDetail{
			IntegrationType: uidtype,
			IntegrationUID:  userclaims.Uid,
			OperationDetail: awardinfo.AwardName,
			OperationType:   "奖品兑换",
			OperationNumber: awardinfo.NeedIntegration,
			AfterOperation:  afterNumber,
		}).Error; err != nil {
			return err
		}
		// 创建奖品兑换记录
		if err := tx.Create(&models.AwardDetail{
			AwardUID:        userclaims.Uid,
			AwardID:         awardId,
			AwardName:       awardinfo.AwardName,
			NeedIntegration: awardinfo.NeedIntegration,
			AwardStatus:     0,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	// 兑换成功
	return 200
}

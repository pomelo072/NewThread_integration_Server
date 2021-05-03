package handlers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"integration_server/database"
	"integration_server/models"
	"os"
	"strconv"
	"strings"
)

// ImportIMG("奖品", v["id"], 1)

// 图片base路径
var basePath = "/integration-app/image/"

// UnpackIMG 解包Base64并写入数据库
func UnpackIMG(imgGroup []string, tp string, id uint) string {
	// 类型文件名词缀
	var tpFlag string
	if tp == "奖品" {
		tpFlag = "award"
	} else if tp == "积分" {
		tpFlag = "integration"
	} else {
		return "Type Error."
	}
	// 父ID
	fatherID := strconv.Itoa(int(id))
	// 遍历文件
	for k, v := range imgGroup {
		numString := strconv.Itoa(k)
		// 解包base64并写入文件
		filename := UnpackFunc(v, tpFlag, fatherID, numString)
		if filename == "ERROR" {
			return "OS Write Error."
		}
		// 写入数据库
		if err := database.Db.Create(&models.Image{
			ImageType: tp,
			FatherID:  id,
			ImageURL:  filename,
		}).Error; err != nil {
			return "DB Write Error."
		}
	}
	return "SUCCESS"
}

// UnpackFunc 解包Base64写入文件
func UnpackFunc(imgBase64 string, tp string, fatherID string, num string) string {
	// 图片为JPG时
	if strings.Contains(imgBase64[:20], "image/png") {
		// 创建Reader
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imgBase64[22:]))
		// 图片解码
		img, _, _ := image.Decode(reader)
		filename := tp + fatherID + num + ".png"
		// 打开文件
		file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		defer file.Close()
		png.Encode(file, img)
		// 返回文件名
		return filename
	}
	// 图片为PNG时
	if strings.Contains(imgBase64[:20], "image/jpeg") {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imgBase64[23:]))
		// 图片解码
		img, _, _ := image.Decode(reader)
		filename := tp + fatherID + num + ".jpg"
		file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		defer file.Close()
		jpeg.Encode(file, img, &jpeg.Options{Quality: 50})
		return filename
	}
	return "ERROR"
}

// ImportIMG 查询图片并打包返回
func ImportIMG(tp string, id interface{}, flag int) interface{} {
	// 判断返回类型:
	// 0 数组
	// 1 字符串
	if flag == 0 {
		// 定义返回数组
		var list []string
		var urlList []models.Image
		database.Db.Model(&models.Image{}).Where("image_type = ? AND father_id = ?", tp, id).Find(&urlList)
		// 遍历添加至List
		for _, v := range urlList {
			list = append(list, v.ImageURL)
		}
		return list
	} else {
		// 查询某张图片返回
		urlList := new(models.Image)
		database.Db.Model(&models.Image{}).Where("image_type = ? AND father_id = ?", tp, id).First(&urlList)
		return urlList.ImageURL
	}
}

// DownloadIMG 图片接口
func DownloadIMG(ctx *gin.Context) {
	imgPath := ctx.Param("imgPath")
	ctx.File(basePath + imgPath)
}

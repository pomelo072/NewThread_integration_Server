package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"strconv"
)

//GetSHAEncode 数据加密
func GetSHAEncode(str string) string {
	w := sha256.New()
	io.WriteString(w, str)            //将str写入到w中
	bw := w.Sum(nil)                  //w.Sum(nil)将w的hash转成[]byte格式
	shastr2 := hex.EncodeToString(bw) //将 bw 转成字符串
	return shastr2
}

// GetReturnData 打包数据和Code,Msg
func GetReturnData(dt interface{}, msgstring string) *gin.H {
	var flag bool
	if msgstring == "SUCCESS" {
		flag = true
	} else {
		flag = false
	}
	result := gin.H{"isSuccess": flag, "Data": dt}
	return &result
}

// Paginate 分页器
// 分页数据请求分页大小为pagesize
func Paginate(pages string, pagesize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(pages)
		if page <= 0 {
			page = 1
		}

		offset := (page - 1) * pagesize
		return db.Offset(offset).Limit(pagesize)
	}
}

// IsContainArr 判断路由是否需要校验Token
func IsContainArr(noVerityAddr []interface{}, c string) bool {
	for _, noAddr := range noVerityAddr {
		if c == noAddr {
			return true
		}
	}
	return false
}


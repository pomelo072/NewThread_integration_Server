package json_struct

import "github.com/dgrijalva/jwt-go"

// UserVerify 用户登录校验
type UserVerify struct {
	Uid          string `json:"username" binding:"required"`
	Verification string `json:"password" binding:"required"`
}

// UserEditor 用户修改信息
type UserEditor struct {
	UserPassword string `json:"userPassword" binding:"required"`
}

// UserClaims 用户JWT的payload结构
type UserClaims struct {
	Uid          string `json:"uid"`
	Verification string `json:"verification"`
	Auth         int    `json:"auth"`
	jwt.StandardClaims
}

// IntegrationModel 用户提交申请JSON模型
type IntegrationModel struct {
	IntegrationType int      `json:"integration_type" binding:"required"`
	ApplyText       string   `json:"apply_text" binding:"required"`
	ApplyType       string   `json:"apply_type" binding:"required"`
	ApplyLevel      string   `json:"apply_level" binding:"required"`
	ContestLevel    string   `json:"contest_level" binding:"required"`
	ApplyIMG        []string `json:"apply_img" binding:"required"`
}

// User 用户JSON结构
type User struct {
	UID   string
	Group string
}

// UserAdd 新增用户JSON模型
type UserAdd struct {
	Users []User `json:"users" binding:"required"`
}

// Award 奖品JSON模型
type Award struct {
	AwardType         string   `json:"award_type" binding:"required"`
	AwardMenu         string   `json:"award_menu" binding:"required"`
	AwardName         string   `json:"award_name" binding:"required"`
	AwardIntroduction string   `json:"award_introduction" binding:"required"`
	AwardIMG          []string `json:"award_img" binding:"required"`
	NeedIntegration   int      `json:"need_integration" binding:"required"`
	InStock           int      `json:"in_stock" binding:"required"`
	UsedNumber        int      `json:"used_number" binding:"required"`
}

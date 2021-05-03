package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"integration_server/json_struct"
	"integration_server/utils"
	"time"
)

// 自定义私钥, 不校验token的页面, token有效期
var (
	secret     = []byte("Pomelo")
	noVerity   = []interface{}{"/api/user/login", "/api/award/info", "/api/award/list"}
	effectTime = 2 * time.Hour
)

// GenerateToken 生成Token
func GenerateToken(claims *json_struct.UserClaims) string {
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		panic(err)
	}
	return sign
}

// JwtVerity 校验Token
func JwtVerity(ctx *gin.Context) {
	// 判断是否需要校验
	if utils.IsContainArr(noVerity, ctx.Request.RequestURI) {
		return
	}
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(200, utils.GetReturnData(gin.H{"error": "Token Need"}, "ERROR"))
		return
	}
	ctx.Set("user", ParseToken(token))
}

// ParseToken 解析Token
func ParseToken(tokenString string) *json_struct.UserClaims {
	token, err := jwt.ParseWithClaims(tokenString, &json_struct.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*json_struct.UserClaims)
	if !ok {
		panic("token valid")
	}
	return claims
}

// Refresh 更新Token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &json_struct.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*json_struct.UserClaims)
	if !ok {
		panic("token valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}

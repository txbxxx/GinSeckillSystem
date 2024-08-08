/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 20:09
 * @File:  Token
 * @Software: GoLand
 **/

package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

// UserClaims 定义Claims的结构体，作用是存储用户信息
type UserClaims struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"is_admin"`
	jwt.RegisteredClaims
}

// myKey 密钥
var myKey = []byte("golangLearn")

// GetMd5 讲密码转换为md值
func GetMd5(s string) string {
	//直接使用md5的Sum方法
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken 生成Token
func GenerateToken(identity, name string, isAdmin int) (string, error) {
	//创建一个Claims,将传入到的用户信息放入Claims
	userClaim := &UserClaims{
		Identity: identity,
		Name:     name,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)), //Token超时时间为60分钟
		},
	}
	//NewWithClaims使用指定的签名方法和声明创建一个新令牌，使用SigningMethodHS256方法对Claims进行签名
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	//SignedString创建并返回一个完整的、有符号的JWT。使用令牌中指定的SigningMethod对令牌进行签名
	token, err := claim.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return token, err
}

// AnalyseToken 解析Token
func AnalyseToken(token string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	//ParseWithClaims像Parse一样解析、验证和验证，但提供一个实现Claims接口的默认对象。这提供了可以覆盖的默认值，并允许调用者使用自己的类型，而不是Claims的默认MapClaims实现
	claims, err := jwt.ParseWithClaims(token, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, err
	}
	return userClaim, err
}

// GenerateUUID 生成UUID
func GenerateUUID() string {
	return uuid.NewV4().String()
}

// GenerateCode 随机生成code验证码
func GenerateCode() string {
	//时间戳作为随机数种子
	rand.NewSource(time.Now().UnixNano())
	//生成随机数
	var code string
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%x", rand.Intn(10))
	}
	return code
}

/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 21:59
 * @File:  UserLoginSvc
 * @Software: GoLand
 **/

package userSvc

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-RocketMQConsumer/model"
	"go-RocketMQConsumer/utils"
	"gorm.io/gorm"
)

// UserLoginService 登录验证器
type UserLoginService struct {
	Name     string `form:"name" json:"name" binding:"required,min=5,max=10" `
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 登录service
func (service *UserLoginService) Login() gin.H {
	//获取用户信息
	var data model.User
	var cunt int64
	errData := utils.DB.Model(&model.User{}).Where("name = ? and password = ?", service.Name, utils.GetMd5(service.Password)).Count(&cunt).Find(&data).Error
	if errData != nil {
		logrus.Error("获取用户信息失败: " + errData.Error())
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			}
		}
		return gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errData.Error(),
		}
	}

	if cunt == 0 {
		return gin.H{
			"code": -1,
			"msg":  "用户名不存在",
		}
	}

	//生成token
	token, errToken := utils.GenerateToken(data.Identity, data.Name, data.IsAdmin)
	if errToken != nil {
		logrus.Error("生成Token失败: " + errToken.Error())
		return gin.H{
			"code": "-1",
			"msg":  "生成Token失败",
		}
	}

	return gin.H{
		"code": 200,
		"msg":  "登录成功！！",
		"data": gin.H{
			"token":    token,
			"is_admin": data.IsAdmin,
		},
	}
}

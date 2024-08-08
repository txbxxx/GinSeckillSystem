/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 21:40
 * @File:  UserSvc
 * @Software: GoLand
 **/

package userSvc

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-RocketMQProducer/model"
	serializes "go-RocketMQProducer/serialized"
	"go-RocketMQProducer/utils"
)

// UserRegisterService 注册用户验证器
type UserRegisterService struct {
	Name     string `form:"name" json:"name" binding:"required,min=5,max=10" `
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	Mail     string `form:"mail" json:"mail" binding:"required,email"`
}

// Register 注册用户service
func (service *UserRegisterService) Register() gin.H {
	//查找是否存在用户
	var userData model.User
	errSearchUser := utils.DB.Model(&model.User{}).Where("name = ? ", service.Name).Find(&userData).Error
	if errSearchUser != nil {
		logrus.Error("查找用户失败" + errSearchUser.Error())
		return gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue",
		}

	}
	if userData != (model.User{}) {
		return gin.H{
			"code": -1,
			"msg":  "用户已存在",
		}
	}

	//创建用户对象
	user := &model.User{
		Name:     service.Name,
		Password: utils.GetMd5(service.Password),
		Identity: utils.GenerateUUID(),
		Phone:    service.Phone,
		Mail:     service.Mail,
	}
	//插入数据库
	errInsert := utils.DB.Model(&model.User{}).Create(user).Error
	if errInsert != nil {
		return gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errInsert.Error(),
		}
	}

	//生成token
	token, errToken := utils.GenerateToken(user.Identity, user.Name, user.IsAdmin)
	if errToken != nil {
		return gin.H{
			"code": -1,
			"msg":  "生成Token失败" + errToken.Error(),
		}
	}

	return gin.H{
		"code": 200,
		"msg":  "注册成功！！",
		"data": gin.H{
			"data":  serializes.UserSerializeSingle(*user),
			"token": token,
		},
	}
}

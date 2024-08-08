/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 21:12
 * @File:  User
 * @Software: GoLand
 **/

package control

import (
	"github.com/gin-gonic/gin"
	"go-RocketMQProducer/service/userSvc"
)

// Login 登录
func Login(c *gin.Context) {
	var service userSvc.UserLoginService
	err := c.ShouldBind(&service)
	if err == nil {
		login := service.Login()
		c.JSON(200, login)
	} else {
		c.JSON(200, gin.H{"err": err})
	}
}

// Register 注册
func Register(c *gin.Context) {
	var service userSvc.UserRegisterService
	err := c.ShouldBind(&service)
	if err == nil {
		register := service.Register()
		c.JSON(200, register)
	} else {
		c.JSON(200, gin.H{"err": err})
	}
}

/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 15:28
 * @File:  route
 * @Software: GoLand
 **/

package router

import (
	"github.com/gin-gonic/gin"
	"go-RocketMQProducer/control"
	"go-RocketMQProducer/middleware"
)

func Router() *gin.Engine {
	httpServer := gin.Default()
	//跨域
	httpServer.Use(middleware.Cors())

	user := httpServer.Group("/user")
	{
		user.POST("/login", control.Login)
		user.POST("/register", control.Register)
	}

	httpServer.GET("/doseckill", control.DoSecKill)
	return httpServer
}

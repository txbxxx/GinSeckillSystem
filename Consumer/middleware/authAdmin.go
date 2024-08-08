/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 20:12
 * @File:  authAdmin
 * @Software: GoLand
 **/

package middleware

import (
	"github.com/gin-gonic/gin"
	"go-RocketMQConsumer/utils"
)

// 判断是否为管理员
func authAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("token")
		//解析token返回claims，从claims中获取
		claims, err := utils.AnalyseToken(auth)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "token认证失败",
			})
			c.Abort()
			return
		}
		if claims.IsAdmin != 1 || claims == nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

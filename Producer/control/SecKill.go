/**
 * @Author admin
 * @Description //TODO
 * @Date 2024/8/5 17:08
 * @File:  SecKill
 * @Software: GoLand
 **/

package control

import (
	"github.com/gin-gonic/gin"
	"go-RocketMQProducer/service/seckillSvc"
)

func DoSecKill(c *gin.Context) {
	var svc seckillSvc.SecKillService
	err := c.ShouldBind(&svc)
	if err != nil {
		c.JSON(200, gin.H{"err": err})
	}
	c.JSON(200, svc.Seckill())
}

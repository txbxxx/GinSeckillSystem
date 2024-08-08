/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 14:52
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/gin-gonic/gin"
	"go-RocketMQProducer/conf"
	"go-RocketMQProducer/router"
	"os"
)

func main() {
	//初始化配置
	conf.Init()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := router.Router()
	//启动http服务
	err := r.Run(os.Getenv("GIN_PORT"))
	if err != nil {
		return
	}
	//进程结束关闭Producer
	defer func(PD rocketmq.Producer) {
		err := PD.Shutdown()
		if err != nil {

		}
	}(conf.PD)
}

/**
 * @Author admin
 * @Description 连接RocketMQ代码
 * @Date 2024/8/5 20:12
 * @File:  RocketMQ
 * @Software: GoLand
 **/

package utils

import (
	rocketmq "github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	producerSvc "github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"os"
)

func QueueUtil(ProducerGroup string) (rocketmq.Producer, error) {
	// 设置日志级别为 Error 级别（关闭大多数日志）
	rlog.SetLogLevel("error")
	return rocketmq.NewProducer(
		//设置组名
		producerSvc.WithGroupName(ProducerGroup),
		//设置nameserver地址
		producerSvc.WithNameServer([]string{os.Getenv("NameSvcAddr")}),
		//设置了连接验证需要配置，如果没有设置则忽略
		producerSvc.WithCredentials(primitive.Credentials{
			AccessKey: os.Getenv("RMQ_AKEY"),
			SecretKey: os.Getenv("RMQ_SKEY"),
		}),
	)
}

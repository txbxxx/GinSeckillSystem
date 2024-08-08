/**
 * @Author admin
 * @Description //TODO
 * @Date 2024/8/6 16:52
 * @File:  RocketMQ
 * @Software: GoLand
 **/

package utils

import (
	rocketmq "github.com/apache/rocketmq-client-go/v2"
	consumerSvc "github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	producerSvc "github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"os"
)

func Producer(ProducerGroup string) (rocketmq.Producer, error) {
	// 设置日志级别为 Error 级别（关闭大多数日志）
	rlog.SetLogLevel("error")
	return rocketmq.NewProducer(
		producerSvc.WithGroupName(ProducerGroup),
		producerSvc.WithNameServer([]string{os.Getenv("NameSvcAddr")}),
		producerSvc.WithCredentials(primitive.Credentials{
			AccessKey: os.Getenv("RMQ_AKEY"),
			SecretKey: os.Getenv("RMQ_SKEY"),
		}),
	)
}

func Consumer(ConsumerGroup string) (rocketmq.PushConsumer, error) {
	// 设置日志级别为 Error 级别（关闭大多数日志）
	rlog.SetLogLevel("error")
	return rocketmq.NewPushConsumer(
		consumerSvc.WithGroupName(ConsumerGroup),
		consumerSvc.WithNameServer([]string{os.Getenv("NameSvcAddr")}),
		consumerSvc.WithCredentials(primitive.Credentials{
			AccessKey: os.Getenv("RMQ_AKEY"),
			SecretKey: os.Getenv("RMQ_SKEY"),
		}),
	)

}

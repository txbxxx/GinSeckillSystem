/**
 * @Author admin
 * @Description //TODO
 * @Date 2024/8/6 17:42
 * @File:  SeckillConsumer
 * @Software: GoLand
 **/

package Consumer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	consumerSvc "github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/sirupsen/logrus"
	"go-RocketMQConsumer/conf"
	"go-RocketMQConsumer/model"
	"go-RocketMQConsumer/utils"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Redis_time 次数时间
var Redis_time = 20

func ReceptionMsg() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//创建Consumer
	consumer, err := utils.Consumer("SeckillConumserGoup3")
	if err != nil {
		logrus.Error("创建消费者错误: ", err.Error())
		return
	}

	//订阅Topic开始消费
	err = Sub(err, consumer)
	if err != nil {
		logrus.Error("订阅主题错误", err)
		return
	}
	//启动消费者
	err = consumer.Start()
	if err != nil {
		logrus.Error("启动消费者错误", err.Error())
		return
	}
	//挂起消费者
	<-sig
	err = consumer.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
		return
	}
}

// Sub 订阅主题开始消费工作
func Sub(err error, consumer rocketmq.PushConsumer) error {
	// 订阅Topic开始消费
	err = consumer.Subscribe("SecKillTopic3", consumerSvc.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumerSvc.ConsumeResult, error) {
			for i := range msgs {
				//开启事物
				tx := utils.DB.Begin()
				// 分割队列中的消息，0为GoodsID 1为UserID
				fmt.Println("收到消息:", string(msgs[i].Body))
				split := strings.Split(string(msgs[i].Body), "-")
				err := processMsg(split, tx)
				if err != nil {
					logrus.Error("处理消息:", string(msgs[i].Body), " 错误", err)
					return consumerSvc.ConsumeRetryLater, err
				}
			}
			return consumerSvc.ConsumeSuccess, nil
		},
	)
	return err
}

func processMsg(split []string, tx *gorm.DB) error {

	for current := 0; current < Redis_time; current++ {
		//开启Redis锁使用SETNX
		if flag := conf.RDB.SetNX(context.Background(), "goods_lock:"+split[0], "", time.Second*20).Val(); flag {
			//减库存
			//如果事物报错则回滚
			if txerr := tx.Model(&model.Goods{}).Where("goods_id = ?", split[0]).Update("stock", gorm.Expr("stock - ?", 1)).Error; txerr != nil {
				return txerr
			}
			// 将这两项插入数据库中
			if txerr := tx.Create(&model.Seckill{GoodsID: split[0], UserID: split[1]}).Error; txerr != nil {
				return txerr
			}
			// 提交事务
			if err := tx.Commit().Error; err != nil {
				// 如果提交失败，则回滚事务
				tx.Rollback()
				return err
			}
			// 操作完成没有报错就删除锁
			conf.RDB.Del(context.Background(), "goods_lock:"+split[0])
			return nil
		} else {
			//如果没拿到锁就等待2秒
			time.Sleep(time.Second * 5)
		}
	}
	//return后判断是回滚还是提交
	return fmt.Errorf("未能获取锁 goods_lock:%s", split[0])
}

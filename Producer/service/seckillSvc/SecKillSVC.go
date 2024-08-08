/**
 * @Author admin
 * @Description //TODO
 * @Date 2024/8/5 17:19
 * @File:  SecKillSVC
 * @Software: GoLand
 **/

package seckillSvc

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-RocketMQProducer/conf"
	"strconv"
	"sync"
	"sync/atomic"
)

var counter int64

type SecKillService struct {
	GoodsID string `form:"goods_id" json:"goods_id"`
	//UserID  string `form:"user_id" json:"user_id"`
}

// MQ报错返回
var mqErr = gin.H{
	"code": -1,
	"msg":  "抢购人数太火爆啦，请稍后再试",
}

func (service *SecKillService) Seckill() gin.H {
	// 初始化计数器
	atomic.AddInt64(&counter, 1)

	//生成uk
	ukey := service.GoodsID + "-" + strconv.FormatInt(atomic.LoadInt64(&counter), 10)
	//如果是第一抢这个商品，就生成的uk加入到redis里面，使用setnx
	nx := conf.RDB.SetNX(context.Background(), "uk:"+ukey, "", 0)
	if !nx.Val() {
		return gin.H{
			"code": -1,
			"msg":  "您已经来过了，去看看的商品吧",
		}
	}

	// 将redis内的库存-1，如果为0就表示商品已经卖完了
	decr := conf.RDB.Decr(context.Background(), "goodsID:"+service.GoodsID)
	if decr.Val() <= 0 {
		return gin.H{
			"code": -1,
			"msg":  "该商品已经被抢购完毕",
		}
	}

	err := service.ASyncSendProducer(ukey)
	if err != nil {
		return mqErr
	}
	// 创建一个通道用于接收异步发送的结果
	//resultChan := make(chan error, 1)
	//go service.ASyncSendProducer(ukey, resultChan)
	//
	//// 在这里你可以选择等待结果，或者做其他事情
	//// 如果需要等待结果，可以使用 select 或 time.After
	//select {
	//case err := <-resultChan:
	//	if err != nil {
	//		logrus.Error(err.Error())
	//		// 处理发送错误
	//		return mqErr
	//	}
	//case <-time.After(time.Second * 5): // 设置超时时间
	//	// 超时处理
	//	logrus.Error("发送请求超时！")
	//	return mqErr
	//}

	// 返回gin.H
	return gin.H{
		"code": 200,
		"msg":  "抢购成功，请稍后去订单详情页查看",
	}
}

func (service *SecKillService) ASyncSendProducer(ukey string) gin.H {
	err := conf.PD.Start()
	if err != nil {
		logrus.Error("启动Producer失败", err)
		return mqErr
	}
	var wg sync.WaitGroup
	wg.Add(1)
	err = conf.PD.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			logrus.Error("发送失败 UserID:"+strconv.FormatInt(atomic.LoadInt64(&counter), 10)+" GoodsID: "+service.GoodsID, " err:", err, " result:", result)
		} else {
			fmt.Println("发送成功")
		}
		wg.Done()
	}, primitive.NewMessage("SecKillTopic3", []byte(ukey)))
	if err != nil {
		logrus.Error("发送至RocketMQ失败: ", err.Error(), "msg: ", ukey)
	}
	wg.Wait()

	return nil
}

//func (service *SecKillService) ASyncSendProducer(ukey string, res chan<- error) {
//	producer, err := utils.QueueUtil("SeckillProducerGroup")
//	if err != nil {
//		logrus.Error("连接RocketMQ失败")
//		res <- err
//		return
//	}
//
//	err = producer.Start()
//	if err != nil {
//		logrus.Error("启动RocketMQ失败")
//		res <- err
//		return
//	}
//
//	err = producer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
//		if err != nil {
//			logrus.Error("发送失败 UserID:"+service.UserID+" GoodsID: "+service.GoodsID, "err:", err)
//			res <- err // 将错误发送到通道
//			return
//		} else {
//			logrus.Info("发送成功")
//			res <- nil // 发送nil表示无错误
//		}
//	}, primitive.NewMessage("SecKillTopic", []byte(ukey)))
//	if err != nil {
//		logrus.Error("发送至RocketMQ失败: ", err.Error(), "msg: ", ukey)
//		res <- err // 如果发送异步消息时出错，也发送到通道
//	}
//
//	err = producer.Shutdown()
//	if err != nil {
//		res <- err
//	}
//}

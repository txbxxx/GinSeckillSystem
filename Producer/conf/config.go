/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 16:14
 * @File:  config
 * @Software: GoLand
 **/

package conf

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go-RocketMQProducer/utils"
	"os"
	"strconv"
)

// RDB PD 添加两个全局变量一个是Redis，一个是Procuder的
var RDB *redis.Client
var PD rocketmq.Producer

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("读取配置文件环境失败" + err.Error())
	}

	//连接数据库
	//utils.DBUntil(os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_ADDR"), os.Getenv("DB_NAME"), os.Getenv("TABLE_NAME"))

	//连接redis
	RDB = utils.RedisUtils(os.Getenv("RDB_ADDR"), os.Getenv("RDB_PWD"), os.Getenv("RDB_DEFAULT_DB"))

	//logrus配置
	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	logrus.SetLevel(logrus.Level(logLevel))
	logrus.SetReportCaller(true)

	//初始化producer
	PD, err = utils.QueueUtil("SeckillProducerGroup")
	if err != nil {
		logrus.Error("连接RocketMQ失败")
		return
	}

}

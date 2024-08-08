/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 16:14
 * @File:  config
 * @Software: GoLand
 **/

package conf

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go-RocketMQConsumer/model"
	"go-RocketMQConsumer/utils"
	"os"
	"strconv"
)

var RDB *redis.Client

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("读取配置文件环境失败" + err.Error())
	}

	//连接数据库
	utils.DBUntil(os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_ADDR"), os.Getenv("DB_NAME"), os.Getenv("TABLE_NAME"))

	//连接redis
	RDB = utils.RedisUtils(os.Getenv("RDB_ADDR"), os.Getenv("RDB_PWD"), os.Getenv("RDB_DEFAULT_DB"))

	//logrus配置
	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	logrus.SetLevel(logrus.Level(logLevel))
	logrus.SetReportCaller(true)

	//同步数据库内的商品数据到redis中
	var goods model.Goods
	err = utils.DB.Model(&model.Goods{}).Where("goods_id = ?", "10086").Take(&goods).Error
	logrus.Info("正在同步至redis;商品:", goods.GoodsName, ";库存:", goods.Stocks)
	if err != nil {
		logrus.Error("查询商品错误" + err.Error())
	}
	RDB.Set(context.Background(), "goodsID:"+goods.GoodsID, goods.Stocks, 0)
}

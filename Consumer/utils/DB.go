/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 15:57
 * @File:  DB
 * @Software: GoLand
 **/

package utils

import (
	"go-RocketMQConsumer/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// DBUntil 用于连接数据库
func DBUntil(DBUser, DBPwd, DBAddr, DBName, TablePrefix string) {
	//定义gorm的日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	databases := DBUser + ":" + DBPwd + "@tcp(" + DBAddr + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//配置数据库
	db, err := gorm.Open(mysql.Open(databases), &gorm.Config{
		SkipDefaultTransaction: false, //禁用事务
		NamingStrategy: schema.NamingStrategy{ //命名策略
			TablePrefix:   TablePrefix,
			SingularTable: true, //禁用复数名称
		},
		Logger: newLogger,
	})

	sqlDB, err := db.DB()

	// 设置连接池
	sqlDB.SetMaxIdleConns(10)

	//  设置最大打开连接数
	sqlDB.SetMaxOpenConns(20)
	if err != nil {
		log.Println("数据库连接失败", err.Error())
	}

	DB = db
	//创建表
	CreateTable()

}

// CreateTable 使用自动迁移创建表
func CreateTable() {
	err := DB.AutoMigrate(&model.Goods{}, &model.Seckill{})
	if err != nil {
		log.Println("创建表失败", err.Error())
		return
	}
}

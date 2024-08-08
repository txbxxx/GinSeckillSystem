/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 22:59
 * @File:  gormtest
 * @Software: GoLand
 **/

package test

import (
	"go-RocketMQProducer/utils"
	"testing"
)

func TestCreateUser(t *testing.T) {
	md5 := utils.GetMd5("123456")
	println(md5)
}

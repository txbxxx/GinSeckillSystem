/**
 * @Author admin
 * @Description //TODO
 * @Date 2024/8/6 16:28
 * @File:  Seckill
 * @Software: GoLand
 **/

package model

import "gorm.io/gorm"

// Seckill 抢购商品和用户的关联表
type Seckill struct {
	gorm.Model
	UserID  string `gorm:"column:user_id;type:varchar(50);unique" json:"user_id"` //用户id
	GoodsID string `gorm:"column:goods_id;type:varchar(50);" json:"goods_id"`     //商品id
}

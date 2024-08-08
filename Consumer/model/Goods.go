/**
 * @Author admin
 * @Description //TODO
 * @Date 2024/8/5 21:58
 * @File:  Goods
 * @Software: GoLand
 **/

package model

import "gorm.io/gorm"

// Goods 商品表(在这里就是需要秒杀的商品表)
type Goods struct {
	gorm.Model
	GoodsID   string `gorm:"column:goods_id;type:varchar(50);unique" json:"goods_id"`     //商品id
	GoodsName string `gorm:"column:goods_name;type:varchar(50);unique" json:"goods_name"` //商品名
	Price     int    `gorm:"column:price;type:int(8);" json:"price"`                      //商品单价
	Stocks    int    `gorm:"column:stock;type:int(8);unique" json:"stocks"`               // 商品库存
}

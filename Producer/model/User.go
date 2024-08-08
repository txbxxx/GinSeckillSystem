/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 15:56
 * @File:  User
 * @Software: GoLand
 **/

package model

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);unique" json:"identity"`  // 用户的唯一标识
	Name     string `gorm:"column:name;type:varchar(100);unique" json:"name"`         // 用户名
	Password string `gorm:"column:password;type:varchar(32);unique" json:"password"`  // 用户密码
	Phone    string `gorm:"column:phone;varchar(20)" json:"phone"`                    // 电话
	Mail     string `gorm:"column:mail;varchar(100)" json:"mail"`                     // 电子邮件
	IsAdmin  int    `gorm:"column:isadmin;type:tinyint(1);default:2" json:"is_admin"` // 是否为管理员1为是2为不是
}

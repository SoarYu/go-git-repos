package models

import "gorm.io/gorm"

// 用户表
type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`  // 唯一标识
	UserName string `gorm:"column:username;type:varchar(255);" json:"username"` // 唯一标识
	Password string `gorm:"column:password;type:varchar(36);" json:"password"`  // 唯一标识
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

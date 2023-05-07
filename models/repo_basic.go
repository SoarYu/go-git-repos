package models

import "gorm.io/gorm"

// 仓库表
type RepoBasic struct {
	gorm.Model
	Uid      int    `gorm:"column:uid;type:int(11);" json:"uid"`
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 唯一标识
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`
	Desc     string `gorm:"column:desc;type:varchar(255);" json:"desc"`
	Star     int    `gorm:"column:star;type:int(11);default:0;" json:"star"`
}

package model

import (
    "github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model

    Content string `gorm:"type:text"`
    UserId int `gorm:"index"`
    PostId int `gorm:"index"`
}

package model

import (
    "github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model

    Title string
    Content string `gorm:"type:text"`
    Pv int
    UserId int `gorm:"index;not null"` // 为该键创建索引
    Tags []Tag `gorm:"many2many:post_tags;"` // 多对多关系
}

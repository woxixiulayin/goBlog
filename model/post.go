package model

import (
    "github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model

    Title string
    Content string `gorm:"type:text"`
    Pv int
    UserId uint
    Tags []Tag `gorm:"many2many:post_tags;"` // 多对多关系
    Comment []Comment // 包含多个comment
}

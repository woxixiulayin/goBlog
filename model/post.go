package model

import (
    "github.com/jinzhu/gorm"
    "goBlog/modules/log"
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


func GetPostsByUserId(userId uint, page int, size int) (*[]Post, error) {
    posts := []Post{}

    if err := DB().Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).Find(&posts).Error; err != nil {
        log.Debugf("Get posts error %v", err)
        return nil, err
    }

    return &posts, nil
}
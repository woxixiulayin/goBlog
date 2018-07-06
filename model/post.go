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

func (p *Post) GetPostById(postId uint) (*Post, error) {
    post := &Post{}
    if err := DB().Where("id = ?", postId).Find(post).Error; err != nil {
        log.Debugf("get post by id error: %v", err)
        return nil, err
    }

    return post, nil
}

func (p *Post) GetPostsByUserId(userId uint, page int, size int) (*[]Post, error) {
    posts := []Post{}

    if err := DB().Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).First(&posts).Error; err != nil {
        log.Debugf("Get posts error %v", err)
        return nil, err
    }

    return &posts, nil
}

func (p *Post) AddPost(title string, content string, userId uint, tags []Tag) (*Post, error){
    tx := DB().Begin()

    post := &Post{
        Title: title,
        Content: content,
        UserId: userId,
        Tags: tags,
    }

    if err := tx.Create(post).Error; err != nil {
        return nil, err
    }

    tx.Commit()

    return post, nil
}
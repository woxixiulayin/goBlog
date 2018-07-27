package model

import (
    "goBlog/modules/log"
)

type Post struct {
    DBModel
    Title     string    `json:"title"`
    Content   string  `gorm:"type:text" json:"content"`
    Pv        int      `json:"pv"`
    UserId    uint     `json:"userId"`
    Tags      []Tag     `gorm:"many2many:post_tags;" json:"tags"` // 多对多关系
    Comments   []Comment  `json:"comments"`// 包含多个comment
}

func (p *Post) GetPostById(postId uint) (*Post, error) {
    post := &Post{}
    if err := DB().Where("id = ?", postId).First(post).Error; err != nil {
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

func (p *Post) CreatePost(title string, content string, userId uint, tags []Tag) (*Post, error){
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

func (p *Post) UpdatePost(postParam map[string]interface{}) (*Post, error) {

    post := &Post{}

    log.Debugf("Post UpdatePost with %v", postParam)
    // find
    if err := DB().First(post, postParam["id"]).Error; err != nil {
        log.Debugf("Post UpdatePost error1: %v", err)
        return nil, err
    }

    delete(postParam, "Id")
    if err := DB().Model(post).Update(postParam).Error; err != nil {
        log.Debugf("Post UpdatePost error2: %v", err)
        return nil, err
    }

    return post, nil
}
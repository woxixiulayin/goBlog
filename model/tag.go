package model

import (
    "goBlog/modules/log"
)

type Tag struct {
    DBModel

    Name string `gorm:"unique"`
}

func (t *Tag) GetTagByName(name string) (*Tag, error) {
    tag := Tag{}

    if err := DB().Where("name = ?", name).First(&tag).Error; err != nil {
        log.Debugf("GetTagByName error %v", err)
        return nil, err
    }

    return &tag, nil
}

func (t *Tag) CreateTagByName(name string) (*Tag, error) {
    tag := Tag{Name: name}

    if err := DB().Create(&tag).Error; err != nil {
        log.Debugf("CreateTagByName error %v", err)
        return nil, err
    }

    return &tag, nil
}
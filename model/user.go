package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

    Name string
    Password string
    Info string
    Post []Post // 一对多
    Comment []Comment //
}

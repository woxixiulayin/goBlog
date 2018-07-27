package model

type User struct {
    DBModel

    Name string
    Password string
    Info string
    Post []Post // 一对多
    Comment []Comment //
}

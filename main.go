package main

import (
	"net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    // "github.com/jinzhu/gorm"
    model "goBlog/model"
)

type Person struct {
    ID uint `json:”id”`
    FirstName string `json:”firstname”`
    LastName string `json:”lastname”`
}

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // 连接数据库
    // if db.HasTable(&Person{}) == false {
    //     db.CreateTable(&Person{})
    // }
    //

    db := model.DB()

    db.Create(&model.User{
        Name: "main",
        Password: "asdddd",
        Info: "is root",
    })

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello world \n")
    })

    e.Logger.Fatal(e.Start(":1234"))

}

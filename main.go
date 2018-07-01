package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    // "github.com/jinzhu/gorm"
    model "goBlog/model"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    model.DB()

    // 静态资源
    e.Static("/assets", "./assets")

    e.Logger.Fatal(e.Start(":1234"))
}

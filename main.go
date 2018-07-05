package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "goBlog/router"
    myTemplate "goBlog/template"
    model "goBlog/model"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    model.DB()

    // 静态资源
    e.Static("/assets", "./assets")

    // 注册renderer
    e.Renderer = myTemplate.Renderer()

    // 注册路由
    router.RegisterRouters(e)
    
    e.Logger.Fatal(e.Start(":1234"))
}

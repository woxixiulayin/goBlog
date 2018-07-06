package router

import (
    "github.com/labstack/echo"
)

func RegisterRouters(e *echo.Echo) error {
    // home
    e.GET("/", HomeHandler)

    // post
    post := e.Group("/posts")
    {
        // 通过get的queryparam获取参数
        post.GET("", GetPostsByUserId)
    }

    return nil
}
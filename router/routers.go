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
        post.GET("", GetPostsByUserId)
        post.GET("/:postId", GetPostsById)
        post.POST("/create", CreatePost)
    }

    return nil
}

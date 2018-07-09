package router

import (
    "net/http"
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
        post.POST("/update", UpdatePost)
    }

    return nil
}

// 设置返回信息的方法
func RequestResult(c echo.Context, code int, data map[string]interface{}, msg string) {
    result := map[string]interface{}{}

    result["code"] = code
    result["data"] = data
    result["msg"] = msg


    c.JSON(http.StatusOK, result)
}
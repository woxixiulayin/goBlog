package router

import (
    "strconv"
    "github.com/labstack/echo"
    "goBlog/model"
)

func PostsHandler(c echo.Context) error {

    userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
    if err != nil {
        panic(err)
    }

    page, err := strconv.Atoi(c.Param("page"))
    if err != nil {
        panic(err)
    }

    pageSize, err := strconv.Atoi(c.Param("page_size"))
    if err != nil {
        panic(err)
    }

    var Post model.Post
    posts, err := getPostsByUserId(userId, page, pageSize)
    if err != nil {
        panic(err)
    }

    return nil
}
package router

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "goBlog/model"
    "goBlog/modules/log"
)

// posts?user=1&page=1&page_size=10
func GetPostsByUserId(c echo.Context) error {

    userId, err := strconv.ParseUint(c.QueryParam("user"), 10, 64)
    if err != nil {
        panic(err)
    }

    page, err := strconv.Atoi(c.QueryParam("page"))
    if err != nil {
        panic(err)
    }

    pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
    if err != nil {
        panic(err)
    }

    log.Debugf("userid: %v, page: %v, page_size: %v", userId, page, pageSize)

    var Post model.Post

    posts, err := Post.GetPostsByUserId(uint(userId), page, pageSize)
    if err != nil {
        panic(err)
    }

    c.JSON(http.StatusOK, map[string]interface{} {
        "code": http.StatusOK,
        "posts": posts,
    })
    return nil
}
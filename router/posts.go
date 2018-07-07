package router

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "goBlog/model"
    "github.com/jinzhu/gorm"
    "goBlog/modules/log"
)

var Post model.Post

// posts/:postId
func GetPostsById(c echo.Context) error {

    result := make(map[string]interface{})

    postId, err := strconv.ParseUint(c.Param("postId"), 10, 64)

    if err != nil {
        panic(err)
    }

    post, err := Post.GetPostById(uint(postId))
    if err != nil {
        switch err {
        case gorm.ErrRecordNotFound:
            result["code"] = http.StatusNotFound
        default:
            result["code"] = http.StatusInternalServerError
        }

        result["post"] = nil
    } else {
        result["code"] = http.StatusOK
        result["post"] = post
    }

    log.Debugf("GetPostsById result is %v", result)

    c.JSON(http.StatusOK, result)

    return nil
}

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
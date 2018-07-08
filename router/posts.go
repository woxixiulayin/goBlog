package router

import (
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/labstack/echo"
    "goBlog/model"
    "github.com/jinzhu/gorm"
    "goBlog/modules/log"
)

var Post model.Post

type BingedCreatePostParam struct {
    Title string `form:"title"`
    Content string `form:"content"`
    Jsontags string `form:"tags"`
}

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

// method： Post， 参数在form中
func CreatePost(c echo.Context) error {
    result := make(map[string]interface{})

    // title := c.FormValue("title")
    // content := c.FormValue("content")
    // jsonTags := c.FormValue("tags")

    // log.Debugf("title is %v, content is %v, jsonTags is %v", title, content, jsonTags)

    // var tags []string
    // json.Unmarshal([]byte(jsonTags), &tags)
    // log.Debugf("postCreate get tags: %v", tags)

    post := &BingedCreatePostParam{}
    if err:= c.Bind(post); err != nil {
        log.Debugf("postCreate bind error %v", err)
        result["code"] = http.StatusBadRequest
        c.JSON(http.StatusOK, result)
        return err
    }

    log.Debugf("post create bind: %v", post)

    userId := uint(1)

    var tags []model.Tag

    var tagStrs []string
    json.Unmarshal([]byte(post.Jsontags), &tagStrs)

    log.Debugf("------------%v", post.Jsontags)
    log.Debugf("------------%v", tagStrs)
    for _, tagStr := range tagStrs{
        var tag = &model.Tag{}

        log.Debugf("tagstr is %v", tagStr)

        if tag, err := tag.GetTagByName(tagStr); err != nil {
            tag.CreateTagByName(tagStr)
            tags = append(tags, *tag)
        } else {
            tags = append(tags, *tag)
        }
    }

    if _, err := Post.CreatePost(post.Title, post.Content, userId, tags); err != nil {
        log.Debugf("postCreate error $v", err)
        result["code"] = http.StatusInternalServerError
        c.JSON(http.StatusOK, result)
        return err
    }

    result["code"] = http.StatusOK
    c.JSON(http.StatusOK, result)

    return nil
}

// func postUpdate(c echo.Context) error {

// }
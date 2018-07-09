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

type BingedCreatePostParam struct {
    Title string `form:"title"`
    Content string `form:"content"`
    Jsontags string `form:"tags"`
}

type BingedUpdatePostParam struct {
    Title string `form:"title"`
    Content string `form:"content"`
    Id string `form:"id"`
}

// posts/:postId
func GetPostsById(c echo.Context) error {

    result := make(map[string]interface{})

    postId, err := strconv.ParseUint(c.Param("postId"), 10, 64)

    if err != nil {
        panic(err)
    }

    var Post model.Post
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

// method： Post， 参数在form中
func CreatePost(c echo.Context) error {
    result := make(map[string]interface{})

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

    var Post model.Post

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

func UpdatePost(c echo.Context) error {
    result := make(map[string]interface{})
    post := model.Post{}

    postParam := BingedUpdatePostParam{}

    if err := c.Bind(&postParam); err != nil {
        log.Debugf("[UpdatePost]bind data error: %v", err)
        result["code"] = http.StatusBadRequest
        c.JSON(http.StatusOK, result)
        return err
    }

    postId, err := strconv.ParseUint(postParam.Id, 10, 64)
    if err != nil {
        RequestResult(c, http.StatusBadRequest, nil, "wrong id")
        return err
    }

    if _, err := post.GetPostById(uint(postId)); err != nil {
        RequestResult(c, http.StatusNotFound, nil, "post not found")
        return err
    }

    post.ID = uint(postId)

    if err := model.DB().Model(&post).Update(map[string]interface{}{
        "Title": postParam.Title,
        "Content": postParam.Content,
    }).Error; err != nil {
        log.Debugf("[UpdatePost] update error: %v", err)
        var code int
        switch err {
        case gorm.ErrRecordNotFound:
            code = http.StatusNotFound
        default:
            code = http.StatusInternalServerError
        }
        RequestResult(c, code, nil, "update post error")
        return err
    }

    RequestResult(c, http.StatusOK, nil, "ok")

    return nil

}
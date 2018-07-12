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

// 绑定只能为字符串类型
type BingedCreatePostParam struct {
    Title string `form:"title"`
    Content string `form:"content"`
    JsonTags string `form:"tags"`
}

type BingedUpdatePostParam struct {
    Title string `form:"title"`
    Content string `form:"content"`
    Id string `form:"id"`
    Pv string `form:"pv"`
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
    json.Unmarshal([]byte(post.JsonTags), &tagStrs)

    log.Debugf("------------%v", post.JsonTags)
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

    RequestResult(c, http.StatusOK, nil, "ok")

    return nil
}

/**
* param: id, title, content, pv, tags
*/
func UpdatePost(c echo.Context) error {
    postParam := BingedUpdatePostParam{}

    if err := c.Bind(&postParam); err != nil {
        log.Debugf("[UpdatePost]bind data error: %v", err)
        RequestResult(c, http.StatusBadRequest, nil, "error")
        return err
    }

    s := c.FormValue("pv")
    v, _ := c.FormParams()
    log.Debugf("update post with postParam: %v, %v, %v", postParam, v, s)

    postId, err := strconv.ParseUint(postParam.Id, 10, 64)
    if err != nil {
        RequestResult(c, http.StatusBadRequest, nil, "wrong id")
        return err
    }

    log.Debugf("%v", postId)

    params := make(map[string]interface{})

    params["id"] = postId
    if len(postParam.Title) > 0 {
        params["title"] = postParam.Title
    }
    if len(postParam.Content) > 0 {
        params["content"] = postParam.Content
    }
    if len(postParam.Pv) > 0 {
        if pv, err := strconv.Atoi(postParam.Pv); err != nil {
            RequestResult(c, http.StatusBadRequest, nil, "wrong pv")
            } else {
                params["pv"] = pv
            }
        }
        
    var post *model.Post
    if _, err := post.UpdatePost(params); err != nil {
        return err
    }

    RequestResult(c, http.StatusOK, nil, "ok")
    return nil
}
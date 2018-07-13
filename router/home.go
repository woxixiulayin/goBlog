package router

import (
    "net/http"
    "github.com/labstack/echo"
    assets "goBlog/util/assets"
)

var jsFiles = assets.GetJsFiles()

func HomeHandler(c echo.Context) error {

    return c.Render(http.StatusOK, "home", map[string]interface{} {
        "name": "哈哈哈",
        "title": "goBlog",
        "jsHome": jsFiles["home"],
    })
}
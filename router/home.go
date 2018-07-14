package router

import (
    "net/http"
    "github.com/labstack/echo"
    assets "goBlog/util/assets"
)


func HomeHandler(c echo.Context) error {

    jsFiles := assets.GetJsFiles()

    return c.Render(http.StatusOK, "home", map[string]interface{} {
        "name": "哈哈哈",
        "title": "goBlog",
        "jsHome": jsFiles.Home,
    })
}
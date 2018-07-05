package router

import (
    "net/http"
    "github.com/labstack/echo"
)

func HomeHandler(c echo.Context) error {
    return c.Render(http.StatusOK, "home", map[string]interface{} {
        "name": "哈哈哈",
        "title": "goBlog",
    })
}
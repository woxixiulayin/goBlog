package router

import (
    "net/http"
    "github.com/labstack/echo"
)

/**
    return login html
*/
func LoginHandler(c echo.Context) error {

    return c.Render(http.StatusOK, "login", map[string]interface{} {})
}

func LoginPostHandler(c echo.Context) error {

    return c.String(http.StatusOK, "yes")
}
package router

import (
    "github.com/labstack/echo"
)

func RegisterRouters(e *echo.Echo) error {
    // home
    e.GET("/", HomeHandler)

    return nil
}
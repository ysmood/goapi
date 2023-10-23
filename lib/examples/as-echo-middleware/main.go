// Package main .
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ysmood/goapi"
)

func main() {
	e := echo.New()

	g := goapi.New()

	goapi.Add(g, hello)

	e.Use(echo.WrapMiddleware(g.Handler))

	_ = e.Start(":3000")
}

func hello(any) string {
	return "World"
}

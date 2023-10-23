// Package main .
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysmood/goapi"
)

func main() {
	e := gin.New()

	g := goapi.New()

	goapi.Add(g, hello)

	e.Use(func(ctx *gin.Context) {
		g.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx.Next()
		})).ServeHTTP(ctx.Writer, ctx.Request)
	})

	_ = e.Run(":3000")
}

func hello(any) string {
	return "World"
}

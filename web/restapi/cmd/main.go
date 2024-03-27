package main

import (
	authGame "restapi/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/scripts", "../../frontend/scripts")
	r.Static("/styles", "../../frontend/styles")

	r.LoadHTMLGlob("../../frontend/templates/*.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	authGroup := r.Group("/auth")

	authGame.AuthGameRoute(r, authGroup)

	r.Run() // слушаем и обслуживаем на 0.0.0.0:8080
}

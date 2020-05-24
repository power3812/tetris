package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Static("/assets", "./assets")
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "tetris.html", gin.H{})
    })

    router.Run()
}
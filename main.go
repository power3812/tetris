package main

import (
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    router := gin.Default()
    router.Static("/assets", "./assets")
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "top.html", gin.H{})
    })
    router.GET("/tetris", func(ctx *gin.Context){
        time.Sleep(1 * time.Second)
        ctx.HTML(200, "tetris.html", gin.H{})
    })

    router.Run()
}

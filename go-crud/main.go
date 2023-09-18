package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usukhbayar/go-crud/connect"
	"github.com/usukhbayar/go-crud/controllers"
)

func init() {
	connect.LoadEnvVariables()
	connect.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}

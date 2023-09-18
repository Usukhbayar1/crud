package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/usukhbayar/go-crud/connect"
	"github.com/usukhbayar/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := connect.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsIndex(c *gin.Context) {
	var posts []models.Post
	connect.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	connect.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}
func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	connect.DB.First(&post, id)

	connect.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsDelete(C *gin.Context) {
	id := C.Param("id")

	connect.DB.Delete(&models.Post{}, id)

	C.Status(200)
}

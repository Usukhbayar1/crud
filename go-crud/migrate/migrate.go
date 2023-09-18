package main

import (
	"github.com/usukhbayar/go-crud/connect"
	"github.com/usukhbayar/go-crud/models"
)

func init() {
	connect.LoadEnvVariables()
	connect.ConnectToDB()
}

func main() {
	connect.DB.AutoMigrate(&models.Post{})
}

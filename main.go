package main

import (
	"go-buscador/config"
	"go-buscador/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.InitElasticsearch()
	r.POST("/api/search/documents", controllers.SearchDocuments)
	r.POST("/api/search/word", controllers.SearchByWord)
	r.GET("/api/search/all", controllers.GetAllDocuments) // Novo endpoint
	r.Run()
}

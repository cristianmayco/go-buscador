package controllers

import (
	"go-buscador/models"
	"go-buscador/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchDocuments(c *gin.Context) {
	var searchParams models.GlobalSearchElastic
	if err := c.ShouldBindJSON(&searchParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	results, err := services.SearchDocuments(searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func SearchByWord(c *gin.Context) {
	var request struct {
		Word string `json:"word"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	results, err := services.SearchByWord(request.Word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func GetAllDocuments(c *gin.Context) {
	results, err := services.GetAllDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

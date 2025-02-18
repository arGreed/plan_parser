package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	getQueryRoute string = "/getQuery"
)

var (
	plan  string
	query string
)

func getQuery(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input Input
		err := c.ShouldBindJSON(&input)
		if err != nil || input.Query == "" {
			log.Println("Ошибка работы с json:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Передан некорректный файл"})
			return
		}
		query = input.Query
	}
}

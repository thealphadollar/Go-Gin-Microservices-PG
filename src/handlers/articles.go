package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/models"
)

func ShowIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": models.GetAllArticles(),
	})
}

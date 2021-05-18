package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/helpers"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/models"
)

func ShowIndexPage(c *gin.Context) {
	helpers.RenderResponse(c, gin.H{
		"title":   "Home Page",
		"payload": models.GetAllArticles(),
	}, "index.html")
}

func GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("article_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		article, err := models.GetArticleByID(articleID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		} else {
			helpers.RenderResponse(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		}
	}
}

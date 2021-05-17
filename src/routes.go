package main

import (
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/handlers"
)

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
}

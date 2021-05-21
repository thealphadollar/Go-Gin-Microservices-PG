package main

import (
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/handlers"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/middleware"
)

func initializeRoutes() {
	router.Use(middleware.SetUserStatus())
	router.GET("/", handlers.ShowIndexPage)

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", handlers.GetArticle)
		articleRoutes.GET("/create", middleware.EnsureLoggedIn(), handlers.ShowArticleCreationPage)
		articleRoutes.POST("/create", middleware.EnsureLoggedIn(), handlers.CreateArticle)
	}

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), handlers.ShowRegisterPage)
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), handlers.RegisterNewUser)
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.PerformLogout)
	}
}

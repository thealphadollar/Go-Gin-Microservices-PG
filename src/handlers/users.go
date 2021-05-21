package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/helpers"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/models"
)

func ShowLoginPage(c *gin.Context) {
	helpers.RenderResponse(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func ShowRegisterPage(c *gin.Context) {
	helpers.RenderResponse(c, gin.H{
		"title": "Register",
	}, "register.html")
}

func PerformLogin(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	if models.IsUserValid(username, password) {
		setLoginToken(c)
		helpers.RenderResponse(c, gin.H{
			"title": "Successful Login",
		}, "login-successful.html")
	} else {
		helpers.RenderResponse(c, gin.H{
			"title":        "Login Failed",
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid Credentials",
		}, "login.html")
	}
}

func PerformLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func RegisterNewUser(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	if _, err := models.RegisterNewUser(username, password); err == nil {
		setLoginToken(c)
		helpers.RenderResponse(c, gin.H{
			"title": "Login Successful",
		}, "login-successful.html")
	} else {
		helpers.RenderResponse(c, gin.H{
			"ErrorTitle":   "Failed Registration",
			"ErrorMessage": err.Error(),
		}, "register.html")
	}
}

func setLoginToken(c *gin.Context) {
	token := helpers.GenerateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)
}

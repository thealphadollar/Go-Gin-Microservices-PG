package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderResponse(c *gin.Context, data gin.H, template string) {
	data["is_logged_in"], _ = c.Get("is_logged_in")
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, template, data)
	}
}

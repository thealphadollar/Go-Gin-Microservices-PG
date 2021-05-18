package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/models"
)

var tmpArticleList []models.Article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

// create router for testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../templates/*")
	}
	return r
}

// helper to test response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

// create temporary list
func saveLists() {
	tmpArticleList = models.GetAllArticles()
}

// retrieve original article list
func restoreLists() {
	models.SetAllArticles(tmpArticleList)
}

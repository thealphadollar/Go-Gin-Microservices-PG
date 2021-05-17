package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/handlers"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handlers.ShowIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowArticlePageUnauthenticatedSuccess(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handlers.GetArticle)

	req, _ := http.NewRequest("GET", "/article/view/0", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		page, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(page), "Article 1") > 0

		return statusOK && pageOK
	})
}

func TestShowArticlePageUnauthenticatedWrongID(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handlers.GetArticle)

	req, _ := http.NewRequest("GET", "/article/view/2", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusNotFound
		return statusOK
	})
}

func TestShowArticlePageUnauthenticatedStringID(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handlers.GetArticle)

	req, _ := http.NewRequest("GET", "/article/view/string", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusBadRequest
		return statusOK
	})
}

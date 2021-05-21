package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/handlers"
	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/middleware"
)

func TestShowLoginPage(t *testing.T) {
	r := getRouter(true)

	r.GET("/user/login", handlers.ShowLoginPage)

	req, _ := http.NewRequest("GET", "/user/login", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowRegisterPage(t *testing.T) {
	r := getRouter(true)

	r.GET("/user/register", handlers.ShowRegisterPage)

	req, _ := http.NewRequest("GET", "/user/register", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})
}

func TestPerformLoginSuccess(t *testing.T) {
	r := getRouter(true)

	r.POST("/user/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(getExistingUserPayload()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(getExistingUserPayload())))

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Successful Login</title>") > 0

		var cookieOK bool
		for _, cookie := range w.Result().Cookies() {
			if cookie.Name == "token" && cookie.Value != "" {
				cookieOK = true
				break
			}
		}

		return statusOK && pageOK && cookieOK
	})
}

func TestPerformLoginInvalidCreds(t *testing.T) {
	r := getRouter(true)

	r.POST("/user/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(getUserPayload()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(getUserPayload())))

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login Failed</title>") > 0

		return statusOK && pageOK
	})
}

func TestPerformLoginAuthenticatedFail(t *testing.T) {
	r := getRouter(true)
	r.POST("/user/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(getExistingUserPayload()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(getExistingUserPayload())))
	req.AddCookie(&http.Cookie{Name: "token", Value: "123"})

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusUnauthorized
		return statusOK
	})
}

func getUserPayload() string {
	params := url.Values{}
	params.Add("username", "user")
	params.Add("password", "pass")
	return params.Encode()
}

func getExistingUserPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")
	return params.Encode()
}

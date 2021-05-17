package tests

import "testing"
import "github.com/thealphadollar/Go-Gin-Microservices-PG/src/models"

func TestGetAllArticles(t *testing.T) {
	alist := models.GetAllArticles()
	articleList := models.GetAllArticles()
	// check length
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check values are same
	for i, v := range alist {
		if v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title ||
			v.Content != articleList[i].Content {
			t.Fail()
			break
		}
	}
}

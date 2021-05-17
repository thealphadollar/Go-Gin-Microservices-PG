package tests

import (
	"testing"

	"github.com/thealphadollar/Go-Gin-Microservices-PG/src/models"
)

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

func TestGetArticleByIDExists(t *testing.T) {
	alist := models.GetAllArticles()
	articleID1, err := models.GetArticleByID(1)

	if err != nil || alist[1].ID != articleID1.ID ||
		alist[1].Content != articleID1.Content || alist[1].Title != articleID1.Content {
		t.Fail()
	}
}

func TestGetArticleByIDNotExists(t *testing.T) {
	articleID1, err := models.GetArticleByID(2)

	if err == nil || articleID1 != nil {
		t.Fail()
	}
}

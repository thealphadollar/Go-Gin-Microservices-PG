package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	{ID: 0, Title: "Article 1", Content: "Some body 1"},
	{ID: 1, Title: "Article 2", Content: "Some body 2"},
}

func GetAllArticles() []Article {
	return articleList
}

func SetAllArticles(aList []Article) {
	articleList = aList
}

func GetArticleByID(id int) (*Article, error) {
	for _, article := range articleList {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Article not found!")
}

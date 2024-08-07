package db

import (
	"log"

	"github.com/football-news-backend/pkg/models"
)

const (
	getNewsQuery      = "SELECT title, link, source, creation_date FROM news;"
	getDatedNewsQuery = "SELECT title, link, source, creation_date FROM news WHERE creation_date = $1;"
	createNewsQuery   = "INSERT INTO news(title, link, source, creation_date) VALUES ($1, $2, $3, $4);"
)

func (d Database) GetDatedNews(date string) ([]models.NewsItem, error) {
	news := []models.NewsItem{}
	rows, err := d.DB.Query(getDatedNewsQuery, date)
	if err != nil {
		log.Println("Failed to get dated news from database.")
		return news, err
	}

	for rows.Next() {
		var n models.NewsItem
		rows.Scan(&n.Title, &n.URL, &n.Source, &n.Creation)
		news = append(news, n)
	}

	return news, nil
}

func (d Database) GetNews() ([]models.NewsItem, error) {
	news := []models.NewsItem{}
	rows, err := d.DB.Query(getNewsQuery)
	if err != nil {
		log.Println("Failed to get news from database.")
		return news, err
	}

	for rows.Next() {
		var n models.NewsItem
		rows.Scan(&n.Title, &n.URL, &n.Source, &n.Creation)
		news = append(news, n)
	}

	return news, nil
}

func (d Database) CreateNews(newsItem models.NewsItem) error {
	_, err := d.DB.Exec(createNewsQuery, newsItem.Title, newsItem.URL, newsItem.Source, newsItem.Creation)
	if err != nil {
		log.Println("Failed to create news item in database.")
		return err
	}

	return nil
}

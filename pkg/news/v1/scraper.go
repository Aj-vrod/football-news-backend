package newsv1

import (
	"fmt"
	"time"

	"github.com/football-news-backend/pkg/models"
	"github.com/gocolly/colly"
)

const (
	ninetyMinURL = "https://www.90min.com/categories/football-news"
	goalNewsURL  = "https://www.goal.com/en-in/news"
)

func ScraperNews() []models.NewsItem {
	var news []models.NewsItem
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := scrapeNinetyMin(c, &news) // TODO: check link's relationt title
	if err != nil {
		fmt.Println("Error scrapping 90Min website", err)
	}
	err = scrapeGoalNewsURL(c, &news)
	if err != nil {
		fmt.Println("Error scrapping Goal website", err)
	}

	return news
}

func scrapeNinetyMin(c *colly.Collector, news *[]models.NewsItem) error {
	var err error
	c.OnError(func(_ *colly.Response, e error) {
		err = e
	})

	c.OnHTML("article > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link == "" {
			return
		}
		title := e.ChildText("h3")
		if title == "" {
			return
		}
		newsPiece := models.NewsItem{
			URL:      link,
			Title:    title,
			Source:   ninetyMinURL,
			Creation: setCreationDate(),
		}
		*news = append(*news, newsPiece)
	})
	if err != nil {
		return err
	}
	c.Visit(ninetyMinURL)
	return nil
}

func scrapeGoalNewsURL(c *colly.Collector, news *[]models.NewsItem) error {
	var err error
	c.OnError(func(_ *colly.Response, e error) {
		err = e
	})

	c.OnHTML("article", func(e *colly.HTMLElement) {
		link := e.ChildAttr("div.poster-wrapper > a", "href")
		if link == "" {
			return
		}
		title := e.ChildText("div.content-wrapper > div.content-body > p")
		if title == "" {
			return
		}
		newsPiece := models.NewsItem{
			URL:      fmt.Sprintf("https://www.goal.com%s", link),
			Title:    title,
			Source:   goalNewsURL,
			Creation: setCreationDate(),
		}
		*news = append(*news, newsPiece)
	})

	if err != nil {
		return err
	}

	c.Visit(goalNewsURL)
	return nil
}

func setCreationDate() string {
	currentTime := time.Now()

	return fmt.Sprintf("%d-%d-%d %d:%d:%d",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())
}

package newsv1

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

const (
	ninetyMinURL   = "https://www.90min.com/categories/football-news"
	oneFootballURL = "https://onefootball.com/en/home"
	goalNewsURL    = "https://www.goal.com/en-in/news"
)

type NewsItem struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Source   string `json:"source"`
	Creation string `json:"creation_date"`
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

func Scraper() {
	var news []NewsItem
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})
	c.OnResponse(func(res *colly.Response) {
		fmt.Println("Got a response from ", res.Request.URL.String())
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
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
		newsPiece := NewsItem{
			URL:      link,
			Title:    title,
			Source:   ninetyMinURL,
			Creation: setCreationDate(),
		}
		news = append(news, newsPiece)
	})

	c.Visit(ninetyMinURL)

	out, _ := json.Marshal(news)
	fmt.Println(string(out))
}

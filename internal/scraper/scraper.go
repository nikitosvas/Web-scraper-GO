package scraper

import (
	"Web-parsing/internal/storage"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/spf13/viper"
	"log"
)

type Site struct {
	Name      string `mapstructure:"name"`
	URL       string `mapstructure:"url"`
	Selectors struct {
		Item  string `mapstructure:"item"`
		Title string `mapstructure:"title"`
	} `mapstructure:"selectors"`
}

func Run() {
	var sites []Site
	err := viper.UnmarshalKey("sites", &sites)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	for _, site := range sites {

		c := colly.NewCollector(
			colly.AllowedDomains("www.rbc.ru"),
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
		)

		c.OnHTML(site.Selectors.Item, func(e *colly.HTMLElement) {
			title := e.ChildText(site.Selectors.Title)
			fmt.Println("Title:", title)
			data := map[string]string{"title": title}
			storage.Save(data)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		err := c.Visit(site.URL)
		if err != nil {
			log.Fatal("Failed to visit", site.URL, err)
		}
	}
}

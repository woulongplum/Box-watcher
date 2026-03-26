package scraper

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/woulongplum/Box-watcher/internal/model"
)

type SurugayaScraper struct{}

func (s SurugayaScraper) Parse(url string) (model.Item, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.Item{}, err
	}

	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	name := strings.TrimSpace(doc.Find("#item_title").Text())

	statusText := doc.Find(".text-red").Text()

	inStock := true
	if strings.Contains(statusText, "品切れ") || name == "" {
		inStock = false
	}

	finalStatus := "在庫あり"
	if !inStock {
		finalStatus = "在庫なし"
	}

	return model.Item{
		Name:    name,
		URL: url,
		InStock: inStock,
		Source:  "Surugaya",
		Status: finalStatus,
	}, nil
}

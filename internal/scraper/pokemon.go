package scraper

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PokemonScraper struct {
	URL string
}

func (p PokemonScraper) CheckStock() (StockResult, error) {
	
	resp , err := http.Get(p.URL)
	if err != nil {
		return StockResult{}, err
	}
	defer resp.Body.Close()

	doc , err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return  StockResult{}, err
	}

	name := doc.Find("h2.heading_01").Text()

	var status string
	var inStock = false

	switch { 
		case strings.Contains(name, "在庫品"): 
		status = "在庫あり" 
		inStock = true 
		case strings.Contains(name, "在庫切れ"): 
		status = "在庫切れ" 
		inStock = false 
		case strings.Contains(name, "予約"): 
		status = "予約商品" 
		inStock = true 
		case strings.Contains(name, "再販"): 
		status = "再販待ち" 
		inStock = true 
		default: 
		status = "不明" 
		inStock = false 
	} 
	return StockResult{ 
		ProductName: name, 
		Price: 0, 
		InStock: inStock, 
		Status:status,
		URL: p.URL, 
	}, nil
}

package main

import (
	"fmt"

	"github.com/woulongplum/Box-watcher/internal/scraper"
)

func main() {
	scrapers := []scraper.Scraper{
		scraper.PokemonScraper{
			URL: "https://www.amiami.jp/top/detail/detail?gcode=CARD-00022985",
		},
	}

	for _, s := range scrapers {
		result , err := s.CheckStock()
		if err != nil {
			fmt.Printf("Error checking stock: %v\n", err)
			continue
		}

		fmt.Println("商品名:", result.ProductName) 
		fmt.Println("価格:", result.Price) 
		fmt.Println("在庫:", result.InStock) 
		fmt.Println("URL:", result.URL) 
		fmt.Println("------")
	}
}

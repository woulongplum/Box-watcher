package scraper

import (
	"time"
)

type PokemonScraper struct {}

func (p PokemonScraper) CheckStock() (StockResult, error) {
	time.Sleep(1 * time.Second)

	return StockResult {
		ProductName: "Pokemon Box",
		Price : 5500,
		InStock: true,
		URL: "https://example.com",
	},nil
}

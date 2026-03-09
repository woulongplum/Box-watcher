package main

import (
	"log"

	"github.com/woulongplum/Box-watcher/internal/scraper"
	"github.com/woulongplum/Box-watcher/internal/service"
)

func main() {

	surugayaScraper := scraper.SurugayaScraper{}

	pokemonService := service.PokemonService{
		Scraper: surugayaScraper,
	}

	targetUrls := []string {
		"https://www.suruga-ya.jp/product/detail/630028922",
		"https://www.suruga-ya.jp/product/detail/630028446",
		"https://www.suruga-ya.jp/product/detail/630027321",
	}

	_, err := pokemonService.Execute(targetUrls)

	if err != nil {
		log.Fatalf("エラーが発生しました: %v", err)
	}
}

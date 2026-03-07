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

	_, err := pokemonService.Execute()

	if err != nil {
		log.Fatalf("エラーが発生しました: %v", err)
	}
}

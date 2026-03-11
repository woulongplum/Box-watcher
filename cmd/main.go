package main

import (
	"fmt"
	"log"
	"time"

	"github.com/woulongplum/Box-watcher/internal/scraper"
	"github.com/woulongplum/Box-watcher/internal/service"
)

func main() {

	for {
		surugayaScraper := scraper.SurugayaScraper{}

		pokemonService := service.PokemonService{
			Scraper: surugayaScraper,
		}

		targetUrls := []string{
			"https://www.suruga-ya.jp/product/detail/630028922",
			"https://www.suruga-ya.jp/product/detail/630028446",
			"https://www.suruga-ya.jp/product/detail/630027321",
		}

		results, err := pokemonService.Execute(targetUrls)

		if err != nil {
			log.Printf("調査中にエラーが発生しましたが続行します: %v", err)
		} else {
			log.Printf("調査完了: %d 件のアイテムを確認", len(results))
		}

		fmt.Println("５分間休憩します...")
		time.Sleep(5 * time.Minute)
	}

}

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/woulongplum/Box-watcher/internal/notifier"
	"github.com/woulongplum/Box-watcher/internal/scraper"
	"github.com/woulongplum/Box-watcher/internal/service"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("エラー：.envファイルが見つかりません。")
	}

	for {
		webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

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
		} else if len(results) > 0 {
			msg := fmt.Sprintf("【在庫あり！】%d 件のアイテムが見つかりました", len(results))
			notifier.SendDiscordNotification(webhookURL,msg)
			log.Printf("通知を送信しました")
		}

		fmt.Println("５分間休憩します...")
		time.Sleep(5 * time.Minute)
	}

}

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/woulongplum/Box-watcher/internal/notifier"
	"github.com/woulongplum/Box-watcher/internal/repository"
	"github.com/woulongplum/Box-watcher/internal/scraper"
	"github.com/woulongplum/Box-watcher/internal/service"
	"gorm.io/gorm/clause"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("エラー：.envファイルが見つかりません。")
	}

	db , err := repository.InitDB()
	if err != nil {
		log.Fatalf("DBの初期化に失敗しました: %v", err)
	}
	log.Printf("データベースの準備が整いました: %T", db)

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

			for _ , item := range results {


				result := db.Clauses(clause.OnConflict{

					Columns: []clause.Column{{Name: "url"}},

					DoUpdates: clause.AssignmentColumns([]string{
						"name", "price", "in_stock", "updated_at",
					}),
				}).Create(&item)

				if result.Error != nil {
					log.Printf("DB保存エラー: %v", result.Error)
				}
			}
			msg := fmt.Sprintf("【在庫あり！】%d 件のアイテムが見つかりました", len(results))
			notifier.SendDiscordNotification(webhookURL,msg)
			log.Printf("通知を送信しました")
		}

		fmt.Println("５分間休憩します...")
		time.Sleep(5 * time.Minute)
	}

}

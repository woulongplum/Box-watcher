package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/woulongplum/Box-watcher/internal/model"
	"github.com/joho/godotenv"
	"github.com/woulongplum/Box-watcher/internal/notifier"
	"github.com/woulongplum/Box-watcher/internal/repository"
	"github.com/woulongplum/Box-watcher/internal/scraper"
	"github.com/woulongplum/Box-watcher/internal/service"
	
)


func main() {

	// 1. 環境設定の読み込み (.env)
	err := godotenv.Load()
	if err != nil {
		log.Printf("警告：.envファイルが見つかりません。環境変数から直接読み込みます。")
	}

	// 2. データベースの初期化
	db, err := repository.InitDB()
	if err != nil {
		log.Fatalf("DBの初期化に失敗しました: %v", err)
	}
	itemRepo := repository.NewItemRepository(db)
	log.Printf("データベースの準備が整いました: %T", db)

	// 3. 調査員（スクレイパー）を一人ずつ雇っておく（ループの外で1回だけ！）
	surugayaScraper := scraper.NewSurugayaScraper()
	rakutenScraper := scraper.NewRakutenMarketScraper()

	for {
		webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

		// --- 調査対象URLの設定 ---
		surugayaUrls := []string{
			"https://www.suruga-ya.jp/product/detail/630028922",
			"https://www.suruga-ya.jp/product/detail/630028446",
			"https://www.suruga-ya.jp/product/detail/630027321",
		}
		rakutenUrls := []string{
			"https://item.rakuten.co.jp/digitamin/yc172764/",
		}

		// --- 駿河屋の調査実行 ---
		sService := service.PokemonService{Scraper: surugayaScraper}
		sResults, sErr := sService.Execute(surugayaUrls)
		if sErr != nil {
			log.Printf("駿河屋の調査中にエラーが発生しました: %v", sErr)
		}

		// --- 楽天の調査実行 ---
		rService := service.PokemonService{Scraper: rakutenScraper}
		rResults, rErr := rService.Execute(rakutenUrls)
		if rErr != nil {
			log.Printf("楽天の調査中にエラーが発生しました: %v", rErr)
		}

		// --- 結果の合体と処理 ---
		allResults := append(sResults, rResults...)

		// 在庫があるアイテムだけをピックアップする
		var foundItems []model.Item
		for _, item := range allResults {
			// DBを更新（在庫の有無に関わらず最新状態を保存）
			if err := itemRepo.Upsert(&item); err != nil {
				log.Printf("DB保存エラー [%s]: %v", item.Name, err)
			}

			// 在庫ありなら通知用リストに追加
			if item.InStock {
				foundItems = append(foundItems, item)
			}
		}

		// 在庫ありの商品が1つでもあればDiscordに通知
		if len(foundItems) > 0 {
			msg := fmt.Sprintf("【在庫あり速報！】\n%d 件のアイテムが入荷しています！確認してください。", len(foundItems))
			notifier.SendDiscordNotification(webhookURL, msg)
			log.Printf("Discordに通知を送信しました（対象: %d件）", len(foundItems))
		}

		fmt.Println("--- 今回の巡回を終了しました。5分間休憩します ---")
		time.Sleep(5 * time.Minute)
	}

}

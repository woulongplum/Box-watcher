package service

import (
	"fmt"
	"time"

	"github.com/woulongplum/Box-watcher/internal/model"
	"github.com/woulongplum/Box-watcher/internal/scraper"
)

type PokemonService struct {
	Scraper scraper.SurugayaScraper
}

func (s PokemonService) Execute(urls []string) ([]model.Item, error) {

	var results []model.Item

	for _, targetURL := range urls {
		item, err := s.Scraper.Parse(targetURL)
		if err != nil {
			fmt.Printf("調査スキップ [%s]: %v\n", targetURL, err)
			continue
		}

		if item.InStock {
			fmt.Printf("【速報】%s の在庫があります！\n", item.Name)
		} else {
			fmt.Printf("%s は在庫切れです。\n", item.Name)
		}

		results = append(results, item)

		time.Sleep(1 * time.Second) // サーバーへの負荷を避けるために少し待機
	}

	return results, nil
}

package service

import (
	
	"fmt"

	"github.com/woulongplum/Box-watcher/internal/model"
	"github.com/woulongplum/Box-watcher/internal/scraper"
)

type PokemonService struct {
	Scraper scraper.SurugayaScraper
}

func (s PokemonService) Execute() (model.Item, error) {

	targetURL := "https://www.suruga-ya.jp/product/detail/630028446"

	item,err := s.Scraper.Parse(targetURL)
	if err != nil {
		return model.Item{}, fmt.Errorf("在庫チェックに失敗しました: %w", err)
	}

	if item.InStock {
		fmt.Printf("【速報】%s の在庫があります！\n", item.Name)
	} else {
		fmt.Printf("%s は在庫切れです。\n", item.Name)
	}
	return item, nil
}

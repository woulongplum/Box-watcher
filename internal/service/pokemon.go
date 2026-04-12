package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/woulongplum/Box-watcher/internal/model"
	"github.com/woulongplum/Box-watcher/internal/scraper"
)

type ScrapeTask struct {
	Scraper scraper.Scraper
	Urls []string
}

type PokemonService struct {
	Scraper scraper.Scraper
	
}

func (s PokemonService) FetchAllParallel(tasks []ScrapeTask) []model.Item {
	var wg sync.WaitGroup

	resultChan := make(chan []model.Item , len(tasks))

	for _ , task := range tasks {
		wg.Add(1)
		go func (t ScrapeTask)  {
			defer wg.Done()

			results , err := PokemonService{Scraper: t.Scraper}.Execute(t.Urls)
			if err == nil {
				resultChan <- results
			}
		}(task)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var allResults []model.Item
	for res := range resultChan {
		allResults = append(allResults, res...)
	}

	return allResults
}

func (s PokemonService) Execute(urls []string) ([]model.Item, error) {

	var results []model.Item

	for _, targetURL := range urls {
		item, err := s.Scraper.CheckStock(targetURL)
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

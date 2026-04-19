package service

import (
	"fmt"
	"testing"

	"github.com/woulongplum/Box-watcher/internal/model"
	"github.com/woulongplum/Box-watcher/internal/scraper"
)

type MockScraper struct {
	Fail bool
}

func (m *MockScraper) Scrape(url string) ([]model.Item,error) {
	return  []model.Item{{Name: "テスト商品",Price: 1000,InStock: true}},nil
}

func (m *MockScraper) CheckStock(url string) (model.Item, error) {
	if m.Fail && url == "url_error" {
		return model.Item{}, fmt.Errorf("接続失敗")
	}
	return model.Item{Name: "テスト商品", Price: 1000, InStock: true}, nil
}

func TestFetchAllParallel(t *testing.T) {

	var s scraper.Scraper = &MockScraper{}
	
	tasks := []ScrapeTask {
		{Scraper:s, Urls: []string{"url1","url2"}},
		{Scraper:s,Urls:[]string{"url3"}},
	}

	svc := PokemonService{}
	results := svc.FetchAllParallel(tasks)

	expectedCount := 3
	if len(results) != expectedCount {
		t.Errorf("期待した件数は %d ですが、実際は %d でした", expectedCount, len(results))
	}
}

func TestFetchAllParallel_WithError(t *testing.T) {
	s := &MockScraper{Fail: true}

	tasks := []ScrapeTask {
		{Scraper:s, Urls: []string{"url1","url_error"}},	
		{Scraper:s, Urls: []string{"url3"}},
	}

	svc := PokemonService{}
	results := svc.FetchAllParallel(tasks)

	expectedCount := 2
	if len(results) != expectedCount {
		t.Errorf("エラー発生時に2件回収できるはずが、実際は %d 件でした。並列処理が途中で止まっている可能性があります", len(results))
	}
}

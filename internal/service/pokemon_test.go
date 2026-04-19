package service

import (
	"testing"

	"github.com/woulongplum/Box-watcher/internal/model"
	"github.com/woulongplum/Box-watcher/internal/scraper"
)

type MockScraper struct {}

func (m *MockScraper) Scrape(url string) ([]model.Item,error) {
	return  []model.Item{{Name: "テスト商品",Price: 1000,InStock: true}},nil
}

func (m *MockScraper) CheckStock(url string) (model.Item, error) {
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

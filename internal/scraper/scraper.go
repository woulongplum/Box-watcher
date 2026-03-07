package scraper

import (
	"github.com/woulongplum/Box-watcher/internal/model"
)

type Scraper interface {
	CheckStock(url string)(model.Item, error)
}



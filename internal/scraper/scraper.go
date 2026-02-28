package scraper

type StockResult struct {
	ProductName string 
	Price int
	InStock bool
	Status string
	URL  string
}

type Scraper interface {
	CheckStock()(StockResult, error)
}



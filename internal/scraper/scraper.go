package scraper

type StockResult struct {
	ProductName string 
	Price int
	InStock bool
	URL  string
}

type Scraper interface {
	CheckStock()(StockResult, error)
}



package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/woulongplum/Box-watcher/internal/model"
	"golang.org/x/net/html/charset"
)

type RakutenMarketScraper struct {}

func NewRakutenMarketScraper() *RakutenMarketScraper {
	return  &RakutenMarketScraper{}
}

func (s *RakutenMarketScraper) CheckStock(url string) (model.Item, error) {
	
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	client := &http.Client{}

	res,err := client.Do(req)

	if err != nil {
		return  model.Item{},err
	}
	
	defer res.Body.Close()

	utf8Reader, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
    if err != nil {
        return model.Item{}, err
    }

	if res.StatusCode != 200 {
		return  model.Item{}, fmt.Errorf("status code error: %d", res.StatusCode)
	}

	doc , err := goquery.NewDocumentFromReader(utf8Reader)

	if err != nil {
		return  model.Item{}, err
	}

	var item model.Item

	item.URL = url

	item.Source = "楽天市場"

	name := doc.Find(".normal_reserve_item_name").First().Text()

	if name == ""{
		name , _ = doc.Find("meta[itemprop='name']").Attr("content") 
	}
	item.Name = strings.TrimSpace(name)

	fullText := doc.Text()

	isSoldOut := strings.Contains(fullText,"売り切れ")|| strings.Contains(fullText, "在庫切れ") || 	strings.Contains(fullText, "この商品は販売期間外です")


	hasCartButton := strings.Contains(fullText, "かごに追加") || strings.Contains(fullText, "ご購入手続きへ")
	
	if isSoldOut || !hasCartButton {
		item.InStock = false
		item.Status = "在庫なし"
	} else {
		item.InStock = true
		item.Status = "在庫あり"
	}

	return item, nil
	
}

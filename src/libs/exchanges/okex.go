package exchanges

import (
	"log"
	"strconv"
	"strings"
	"tickers/src/models"
	"time"
)

var okexBaseURL = "https://www.okex.com"

// GetOkexTickers ...
func GetOkexTickers () []models.Ticker {
	
	var tickers []models.Ticker

	url := okexBaseURL + "/api/spot/v3/instruments/ticker"

	data := makeRequest(url)

	parsed := data.([]interface{})

	var coin string
	var cur string

	for n := range parsed {
		d := parsed[n].(map[string]interface{})
		coin = strings.Split(d["product_id"].(string), "-")[0]
		cur = strings.Split(d["product_id"].(string), "-")[1]
		bp, err := strconv.ParseFloat(d["best_bid"].(string), 64)
		bps, err := strconv.ParseFloat(d["best_bid_size"].(string), 64)
		ap, err := strconv.ParseFloat(d["best_ask"].(string), 64)
		aps, err := strconv.ParseFloat(d["best_ask_size"].(string), 64)

		if err != nil {
			log.Println(err)
		}

		tickers = append(tickers, models.Ticker{
			coin,
			cur,
			coin+cur,
			bp,
			bps,
			ap,
			aps,
			"okex",
			int(time.Now().Unix())})
	}
	return tickers
}
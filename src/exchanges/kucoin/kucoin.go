package kucoin

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var baseURL = "https://api.kucoin.com"

// GetTickers ...
func GetTickers () []models.Ticker {

	var tickers []models.Ticker

	url := baseURL + "/api/v1/market/allTickers"

	data := utils.MakeRequest(url)

	parsed := data.(map[string]interface{})
	
	for k, v := range parsed {
		if k == "data" {
			for m, n := range v.(map[string]interface{}) {
				if m == "ticker" {
					for _, i := range n.([]interface{}) {
						coin := strings.Split(i.(map[string]interface{})["symbol"].(string), "-")[0]
						cur := strings.Split(i.(map[string]interface{})["symbol"].(string), "-")[1]
						bidPrice, _ := strconv.ParseFloat(i.(map[string]interface{})["buy"].(string), 64)
						askPrice, _ := strconv.ParseFloat(i.(map[string]interface{})["sell"].(string), 64)
						tickers = append(tickers, models.Ticker{
							Coin: coin,
							Currency: cur,
							Symbol: coin + cur,
							BidPrice: bidPrice,
							BidQty: 0.0,
							AskPrice: askPrice,
							AskQty: 0.0,
							Exchange: "kucoin",
							Timestamp: int(time.Now().Unix())})
					}
				}
			}
		}
	}
			
	return tickers
}


package exchanges

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"time"
)

var kucoinBaseURL = "https://api.kucoin.com"

// GetKucoinTickers ...
func GetKucoinTickers () []models.Ticker {

	var tickers []models.Ticker

	url := kucoinBaseURL + "/api/v1/market/allTickers"

	data := makeRequest(url)

	parsed := data.(map[string]interface{})
	
	for k, v := range parsed {
		if k == "data" {
			for m, n := range v.(map[string]interface{}) {
				if m == "ticker" {
					for _, i := range n.([]interface{}) {
						bidPrice, _ := strconv.ParseFloat(i.(map[string]interface{})["buy"].(string), 64)
						askPrice, _ := strconv.ParseFloat(i.(map[string]interface{})["sell"].(string), 64)
						tickers = append(tickers, models.Ticker{
							strings.Split(i.(map[string]interface{})["symbol"].(string), "-")[0],
							strings.Split(i.(map[string]interface{})["symbol"].(string), "-")[1],
							strings.ReplaceAll(i.(map[string]interface{})["symbol"].(string), "-", ""),
							bidPrice,
							0.0,
							askPrice,
							0.0,
							"kucoin",
							int(time.Now().Unix())})
					}
				}
			}
		}
	}
			
	return tickers
}


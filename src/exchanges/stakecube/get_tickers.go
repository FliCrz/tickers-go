package stakecube

import (
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {
	baseUrl := "https://stakecube.io/api/v2/exchange/spot/markets?baseMarket="
	
	currencies := []string{
		"BTC", "SCC", "LTC", "DASH", "DOGE",
	}

	for _, c := range currencies {
		url := baseUrl + c

		data := utils.MakeRequest(url)

		parsed := data.(map[string]interface{})["result"]

		reparsed := parsed.(map[string]interface{})

		for k, v := range reparsed {
			
			coin := strings.Split(k, "_")[0]
			cur := strings.Split(k, "_")[1]

			d := v.(map[string]interface{})

			new_ticker := models.Ticker {
				Coin: coin,
				Currency: cur,
				Symbol: coin + cur,
				BidPrice: d["bestBid"].(float64),
				BidQty: 0.0,
				AskPrice: d["bestAsk"].(float64),
				AskQty: 0.0,
				Exchange: "stakecube",
				Timestamp: int(time.Now().Unix()),
			}

			tickers = append(tickers, new_ticker)
		}
	}

	return tickers
}
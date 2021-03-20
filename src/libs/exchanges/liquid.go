package exchanges

import (
	"tickers/src/models"
	"time"
)

var liquidBaseURL = "https://api.liquid.com"

// GetLiquidTickers ...
func GetLiquidTickers () []models.Ticker {

	var tickers []models.Ticker

	url := liquidBaseURL + "/products"

	data := makeRequest(url)

	parsed := data.([]interface{})

	for _, i := range parsed {
		coin := i.(map[string]interface{})["base_currency"].(string)
		cur := i.(map[string]interface{})["quoted_currency"].(string)
		bp := i.(map[string]interface{})["market_bid"]
		if bp == nil {
			bp = 0.0
		}
		ap := i.(map[string]interface{})["market_ask"]
		if ap == nil {
			ap = 0.0
		}
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bp.(float64),
			BidQty: 0.0,
			AskPrice: ap.(float64),
			AskQty: 0.0,
			Exchange: "liquid",
			Timestamp: int(time.Now().Unix())})
	}
	return tickers
}
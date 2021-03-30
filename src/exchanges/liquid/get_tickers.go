package liquid

import (
	"strconv"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var baseURL = "https://api.liquid.com"

// GetTickers ...
func GetTickers () []models.Ticker {

	var tickers []models.Ticker

	url := baseURL + "/products"

	data := utils.MakeRequest(url)

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

		var vol float64
		v := i.(map[string]interface{})["volume_24h"]
		if v == nil {
			v = 0.0
		} else {
			vol, _ = strconv.ParseFloat(v.(string), 64)
		}
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bp.(float64),
			BidQty: 0.0,
			AskPrice: ap.(float64),
			AskQty: 0.0,
			BaseVolume: vol,
			QuoteVolume: 0.0,
			Exchange: "liquid",
			Timestamp: int(time.Now().Unix())})
	}
	return tickers
}
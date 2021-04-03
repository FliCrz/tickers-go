package cryptocom

import (
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var baseUrl = "https://api.crypto.com/v2"

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {

	raw := utils.MakeRequest(baseUrl + "/public/get-ticker")

	data := raw.(map[string]interface{})["result"].(map[string]interface{})["data"]

	for _, i := range data.([]interface{}) {
		t := models.Ticker {
			Coin: strings.Split(i.(map[string]interface{})["i"].(string), "_")[0],
			Currency: strings.Split(i.(map[string]interface{})["i"].(string), "_")[1],
			Symbol: strings.Replace(i.(map[string]interface{})["i"].(string), "_", "", 1),
			BidPrice: i.(map[string]interface{})["b"].(float64),
			BidQty: 0.0,
			AskPrice: i.(map[string]interface{})["a"].(float64),
			AskQty: 0.0,
			BaseVolume: i.(map[string]interface{})["v"].(float64),
			QuoteVolume: 0.0,
			Exchange: "crypto.com",
			Timestamp: int(time.Now().Unix()),
		}
		tickers = append(tickers, t)
	}
	return tickers
}
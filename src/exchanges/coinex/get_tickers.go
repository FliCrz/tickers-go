package coinex

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {
	url :=  baseUrl + "/market/ticker/all"

	symbols := getSymbols()
	
	raw := utils.MakeRequest(url)

	data := raw.(map[string]interface{})["data"].(map[string]interface{})["ticker"]

	for k, v := range data.(map[string]interface{}) {
		var coin string
		var cur string
		for _, s := range symbols {
			if k == strings.Replace(s, "-", "", 1) {
				coin = strings.Split(s, "-")[0]
				cur = strings.Split(s, "-")[1]
				break
			}
		}

		bid, _ := strconv.ParseFloat(v.(map[string]interface{})["buy"].(string), 64)
		bidQty, _ := strconv.ParseFloat(v.(map[string]interface{})["buy_amount"].(string), 64)
		ask, _ := strconv.ParseFloat(v.(map[string]interface{})["sell"].(string), 64)
		askQty, _ := strconv.ParseFloat(v.(map[string]interface{})["sell_amount"].(string), 64)
		vol, _ := strconv.ParseFloat(v.(map[string]interface{})["vol"].(string), 64)

		t := models.Ticker {
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bid,
			BidQty: bidQty,
			AskPrice: ask,
			AskQty: askQty,
			BaseVolume: vol,
			QuoteVolume: 0.0,
			Exchange: "coinex",
			Timestamp: int(time.Now().Unix()),
		}
		tickers = append(tickers, t)
	}

	return tickers
}

package exchanges

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"time"
)

var poloniexBaseURL = "https://poloniex.com/public"

// GetPoloniexTickers ...
func GetPoloniexTickers () []models.Ticker {

	var tickers []models.Ticker

	url := poloniexBaseURL + "?command=returnTicker"

	data := makeRequest(url)

	parsed := data.(map[string]interface{})

	var coin string
	var cur string

	for k, v := range parsed {
		coin = strings.Split(k, "_")[1]
		cur = strings.Split(k, "_")[0]
		bp, _ := strconv.ParseFloat(v.(map[string]interface{})["highestBid"].(string), 64)
		ap, _ := strconv.ParseFloat(v.(map[string]interface{})["lowestAsk"].(string), 64)
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bp,
			BidQty: 0.0,
			AskPrice: ap,
			AskQty: 0.0,
			Exchange: "poloniex",
			Timestamp: int(time.Now().Unix())})
	}
	return tickers
}
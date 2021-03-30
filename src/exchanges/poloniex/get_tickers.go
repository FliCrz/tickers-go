package poloniex

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var baseURL = "https://poloniex.com/public"

// GetTickers ...
func GetTickers () []models.Ticker {

	var tickers []models.Ticker

	url := baseURL + "?command=returnTicker"

	data := utils.MakeRequest(url)

	parsed := data.(map[string]interface{})

	var coin string
	var cur string

	for k, v := range parsed {
		coin = strings.Split(k, "_")[1]
		cur = strings.Split(k, "_")[0]
		bp, _ := strconv.ParseFloat(v.(map[string]interface{})["highestBid"].(string), 64)
		ap, _ := strconv.ParseFloat(v.(map[string]interface{})["lowestAsk"].(string), 64)
		baseVol, _ := strconv.ParseFloat(v.(map[string]interface{})["baseVolume"].(string), 64)
		quoteVol, _ := strconv.ParseFloat(v.(map[string]interface{})["quoteVolume"].(string), 64)
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bp,
			BidQty: 0.0,
			AskPrice: ap,
			AskQty: 0.0,
			BaseVolume: baseVol,
			QuoteVolume: quoteVol,
			Exchange: "poloniex",
			Timestamp: int(time.Now().Unix())})
	}
	return tickers
}
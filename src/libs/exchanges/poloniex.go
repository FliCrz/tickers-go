package exchanges

import (
	"log"
	"strconv"
	"strings"
	"tickers/src/models"
	"time"
)

var poloniexBaseURL = "https://poloniex.com/public"

// GetPoloniexTickers ...
func GetPoloniexTickers () []models.Ticker {

	log.Println("Getting data from poloniex")

	var tickers []models.Ticker

	url := poloniexBaseURL + "?command=returnTicker"

	data := makeRequest(url)

	parsed := data.(map[string]interface{})

	var coin string
	var cur string

	for k, v := range parsed {
		coin = strings.Split(k, "_")[1]
		cur = strings.Split(k, "_")[0]
		bp, err := strconv.ParseFloat(v.(map[string]interface{})["highestBid"].(string), 64)
		ap, err := strconv.ParseFloat(v.(map[string]interface{})["lowestAsk"].(string), 64)

		if err != nil {
			log.Println(err)
		}
		
		tickers = append(tickers, models.Ticker{
			coin,
			cur,
			coin+cur,
			bp,
			0.0,
			ap,
			0.0,
			"poloniex",
			int(time.Now().Unix())})
	}
	return tickers
}
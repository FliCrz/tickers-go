package hitbtc

import (
	"log"
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var APIURL =  "https://api.hitbtc.com/api/2"

// GetTickers ...
func GetTickers() []models.Ticker {

	url := APIURL + "/public/ticker"
	
	data := utils.MakeRequest(url)

	var tickers []models.Ticker

	parsed := data.([]interface{})

	for _, i := range parsed {

		reparsed := i.(map[string]interface{})

		// log.Println(reparsed)

		symbol := reparsed["symbol"].(string)
		coin := strings.SplitN(reparsed["symbol"].(string), "", 2)[0]
		cur := strings.SplitN(reparsed["symbol"].(string), "", 2)[1]
		
		var bidPrice float64
		var askPrice float64

		if reparsed["bid"] == nil {
			bidPrice = 0.0
		} else {
			bidPrice, _ = strconv.ParseFloat(reparsed["bid"].(string), 64)
		}

		if reparsed["ask"] == nil {
			askPrice = 0.0
		} else {
			askPrice, _ = strconv.ParseFloat(reparsed["ask"].(string), 64)
		}
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: symbol,
			BidPrice: bidPrice,
			BidQty: 0.0,
			AskPrice: askPrice,
			AskQty: 0.0,
			Exchange: "hitbtc",
			Timestamp: int(time.Now().Unix())})
	}
	if len(tickers) == 0 {
		log.Println("Could not get tickers from hitbtc.")
	}
	return tickers
}

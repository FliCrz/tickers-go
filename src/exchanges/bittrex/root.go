package bittrex

import (
	"log"
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers() []models.Ticker {

	url := "https://api.bittrex.com/v3/markets/tickers"
	
	data := utils.MakeRequest(url)

	var tickers []models.Ticker

	parsed := data.([]interface{})

	for _, i := range parsed {

		reparsed := i.(map[string]interface{})

		symbol := reparsed["symbol"].(string)
		coin := strings.Split(symbol, "-")[0]
		cur := strings.Split(symbol, "-")[1]
		bidPrice, _ := strconv.ParseFloat(reparsed["bidRate"].(string), 64)
		askPrice, _ := strconv.ParseFloat(reparsed["askRate"].(string), 64)
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bidPrice,
			BidQty: 0.0,
			AskPrice: askPrice,
			AskQty: 0.0,
			Exchange: "bittrex",
			Timestamp: int(time.Now().Unix())})
	}

	if len(tickers) == 0 {
		log.Println("Could not get tickers from bittrex.")
	}

	return tickers
}
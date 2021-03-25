package bitfinex

import (
	"log"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var APIURL = "https://api-pub.bitfinex.com/v2"


// GetTickers ...
func GetTickers() []models.Ticker {

	url := APIURL + "/tickers?symbols=ALL"
	
	data := utils.MakeRequest(url)
	
	var tickers []models.Ticker
	
	parsed := data.([]interface{})

	symbolsList := getBitfinexRawSymbols()

	for _, i := range parsed {

		reparsed := i.([]interface{})

		symbol := reparsed[0].(string)

		if strings.Index(symbol, "t") == 0 && strings.Contains(symbol, ":") && strings.Contains(symbol, "-") {

			symbol = symbol[1:]
			rawCoin := symbol[:len(symbol) / 2]
			rawCur := symbol[len(symbol) / 2:]

			var coin string
			var cur string

			for _, s := range symbolsList {
				i := s.([]interface{})[0].(string)
				j := s.([]interface{})[1].(string)

				if rawCoin == i {
					coin = strings.ToUpper(j)
				} else if rawCur == i {
					cur = strings.ToUpper(j)
				} else {
					cur = rawCur
					coin = rawCoin
				}
			}

			bidPrice := reparsed[1].(float64)
			bidQty := reparsed[2].(float64)
			askPrice := reparsed[3].(float64)
			askQty := reparsed[4].(float64)
		
			tickers = append(tickers, models.Ticker{
				Coin: coin,
				Currency: cur,
				Symbol: symbol,
				BidPrice: bidPrice,
				BidQty: bidQty,
				AskPrice: askPrice,
				AskQty: askQty,
				Exchange: "bitfinex",
				Timestamp: int(time.Now().Unix())})
		}
	}

	if len(tickers) == 0 {
		log.Println("Could not get tickers from bitfinex.")
	}

	return tickers
}

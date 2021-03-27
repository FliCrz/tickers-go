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

		if strings.Index(symbol, "t") == 0 && strings.Contains(symbol, ":") {
			
			symbol = symbol[1:]
			rawCoin := strings.Split(symbol, ":")[0]
			rawCur := strings.Split(symbol, ":")[1]
			
			var coin string
			var cur string
			
			
			for _, s := range symbolsList {
				i := s.([]interface{})[0]
				j := s.([]interface{})[1]
				
				// log.Println(i, j)

				if rawCoin == i {
					coin = strings.ToUpper(j.(string))
				} else if rawCur == i {
					cur = strings.ToUpper(j.(string))
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

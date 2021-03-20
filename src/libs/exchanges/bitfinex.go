package exchanges

import (
	"log"
	"strings"
	"tickers/src/models"
	"time"
)

var bitfinexAPIURL = "https://api-pub.bitfinex.com/v2"


// GetBitfinexTickers ...
func GetBitfinexTickers() []models.Ticker {

	url := bitfinexAPIURL + "/tickers?symbols=ALL"
	
	data := makeRequest(url)
	
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


func getBitfinexRawSymbols () []interface{} {

	log.Println("Getting symbols data from bitfinex.")

	url := bitfinexAPIURL + "/conf/pub:map:currency:sym"

	data := makeRequest(url)

	var symbols []interface{}
	// log.Println(data)
	parsed := data.([]interface{})
	for _, i := range parsed {
		symbols = append(symbols, i.([]interface{}))
	}

	return symbols
}

package bitfinex

import (
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var APIURL = "https://api-pub.bitfinex.com/v2"


// GetTickers ...
func GetTickers() (tickers []models.Ticker) {

	url := APIURL + "/tickers?symbols=ALL"
	
	data := utils.MakeRequest(url)
	
	parsed := data.([]interface{})

	symbolsList := getBitfinexRawSymbols()

	symbols := symbolsList[0]

	// log.Println(symbolsList)

	for _, i := range parsed {
		
		reparsed := i.([]interface{})
		
		rawSymbol := reparsed[0].(string)
		
		var coin string
		var cur string
		var symbol string
		
		if strings.Index(rawSymbol, "t") == 0 {
			
			symbol = rawSymbol[1:]

			if strings.Contains(symbol, ":") {
				coin = strings.Split(symbol, ":")[0]
				cur = strings.Split(symbol, ":")[1]
			} else {
				for _, s := range symbols.([]interface{}) {
					// log.Println(s)
					if strings.Index(symbol, s.(string)) == 0 {
						coin = s.(string)
					} else if strings.Index(symbol, s.(string)) > 2 {
						cur = s.(string)
					}
				}
			}
			
			if coin == "DSH" {
				coin = "DASH"
			}

			bidPrice := reparsed[1].(float64)
			bidQty := reparsed[2].(float64)
			askPrice := reparsed[3].(float64)
			askQty := reparsed[4].(float64)
			baseVol := reparsed[8].(float64)
			quoteVol := 0.0
		
			tickers = append(tickers, models.Ticker{
				Coin: coin,
				Currency: cur,
				Symbol: coin + cur,
				BidPrice: bidPrice,
				BidQty: bidQty,
				AskPrice: askPrice,
				AskQty: askQty,
				BaseVolume: baseVol,
				QuoteVolume: quoteVol,
				Exchange: "bitfinex",
				Timestamp: int(time.Now().Unix())})
		}
	}
	return tickers
}

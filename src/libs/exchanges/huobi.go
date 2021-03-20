package exchanges

import (
	"log"
	"strings"
	"tickers/src/models"
	"time"
)

// GetHuobiTickers ...
func GetHuobiTickers() []models.Ticker {

	url := "https://api.huobi.pro/market/tickers"
	
	data := makeRequest(url)

	
	var tickers []models.Ticker
	
	parsed := data.(map[string]interface{})
	parsedData := parsed["data"].([]interface{})

	symbolsList := getHuobiRawSymbols()
	
	for _, i := range parsedData {
		var coin string
		var cur string
		
		reparsed := i.(map[string]interface{})
		// fmt.Println(reparsed)

		symbol := reparsed["symbol"].(string)

		for n := range symbolsList {
			s := symbolsList[n].(map[string]interface{})
			if s["symbol"].(string) == symbol {
				coin = strings.ToUpper(s["base-currency"].(string))
				cur = strings.ToUpper(s["quote-currency"].(string))
			}
		}
		// symbol = strings.ToUpper(symbol)

		bidPrice := reparsed["bid"].(float64)
		bidQty := reparsed["bidSize"].(float64)
		askPrice := reparsed["ask"].(float64)
		askQty := reparsed["askSize"].(float64)
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: cur + coin,
			BidPrice: bidPrice,
			BidQty: bidQty,
			AskPrice: askPrice,
			AskQty: askQty,
			Exchange: "huobi",
			Timestamp: int(time.Now().Unix())})
	}

	if len(tickers) == 0 {
		log.Println("Could not get tickers from huobi.")
	}

	return tickers
}

func getHuobiRawSymbols() []interface{} {
	url := "https://api.huobi.pro/v1/common/symbols"
	data := makeRequest(url)
	parsed := data.(map[string]interface{})
	reparsed := parsed["data"].([]interface{})
	var symbols []interface{}
	symbols = append(symbols, reparsed...)
	return symbols
}


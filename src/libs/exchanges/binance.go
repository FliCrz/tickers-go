package exchanges

import (
	"log"
	"strconv"
	"tickers/src/models"
	"time"
)

var binanceAPIURL =  "https://api.binance.com/api/v3"

// GetBinanceTickers ...
func GetBinanceTickers() []models.Ticker {

	url := binanceAPIURL + "/ticker/bookTicker"
	
	data := makeRequest(url)

	var tickers []models.Ticker

	parsed := data.([]interface{})

	symbolsList := getSymbolsData()

	var coin string
	var cur string

	for _, i := range parsed {

		reparsed := i.(map[string]interface{})

		symbol := reparsed["symbol"].(string)

		for n := range symbolsList{
			i := symbolsList[n].(map[string]interface{})
			if i["symbol"] == symbol {
				coin = i["baseAsset"].(string)
				cur = i["quoteAsset"].(string)
			}
		}

		bidPrice, _ := strconv.ParseFloat(reparsed["bidPrice"].(string), 64)
		bidQty, _ := strconv.ParseFloat(reparsed["bidQty"].(string), 64)
		askPrice, _ := strconv.ParseFloat(reparsed["askPrice"].(string), 64)
		askQty, _ := strconv.ParseFloat(reparsed["askQty"].(string), 64)
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: symbol,
			BidPrice: bidPrice,
			BidQty: bidQty,
			AskPrice: askPrice,
			AskQty: askQty,
			Exchange: "binance",
			Timestamp: int(time.Now().Unix())})
	}
	if len(tickers) == 0 {
		log.Println("Could not get tickers from binance.")
	}
	return tickers
}


func getSymbolsData() []interface{} {
	log.Println("Getting symbols from binance")

	url := binanceAPIURL + "/exchangeInfo"
	data := makeRequest(url)
	parsed := data.(map[string]interface{})
	symbolsList := parsed["symbols"].([]interface{})

	return symbolsList
}

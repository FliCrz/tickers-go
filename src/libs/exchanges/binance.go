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

	log.Println("Getting data from binance.")

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

		bidPrice, err := strconv.ParseFloat(reparsed["bidPrice"].(string), 64)
		bidQty, err := strconv.ParseFloat(reparsed["bidQty"].(string), 64)
		askPrice, err := strconv.ParseFloat(reparsed["askPrice"].(string), 64)
		askQty, err := strconv.ParseFloat(reparsed["askQty"].(string), 64)
		
		if err != nil {
			log.Fatal(err)
		}
		
		tickers = append(tickers, models.Ticker{
			coin,
			cur,
			symbol,
			bidPrice,
			bidQty,
			askPrice,
			askQty,
			"binance",
			int(time.Now().Unix())})
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

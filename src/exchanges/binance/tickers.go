package binance

import (
	"log"
	"strconv"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers() []models.Ticker {

	url := APIURL + "/ticker/bookTicker"
	
	data := utils.MakeRequest(url)

	var tickers []models.Ticker

	parsed := data.([]interface{})

	symbolsList := getSymbols()

	var coin string
	var cur string

	volumeData := get24HData()

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
			BaseVolume: volumeData[symbol][0],
			QuoteVolume: volumeData[symbol][1],
			Exchange: "binance",
			Timestamp: int(time.Now().Unix())})
	}
	if len(tickers) == 0 {
		log.Println("Could not get tickers from binance.")
	}
	return tickers
}
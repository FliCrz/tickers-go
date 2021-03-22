package exchanges

import (
	"log"
	"strings"
	"tickers/src/models"
	"time"
)

var crex24APIUrl =  "https://api.crex24.com/v2"

// GetBinanceTickers ...
func GetCrex24Tickers() []models.Ticker {

	url := crex24APIUrl + "/public/tickers"
	
	data := makeRequest(url)

	var tickers []models.Ticker

	parsed := data.([]interface{})

	for _, i := range parsed {

		reparsed := i.(map[string]interface{})

		// log.Println(reparsed)

		symbol := strings.Replace(reparsed["instrument"].(string), "-", "", -1)
		coin := strings.SplitN(reparsed["instrument"].(string), "-", 2)[0]
		cur := strings.SplitN(reparsed["instrument"].(string), "-", 2)[1]
		
		var bidPrice float64
		var askPrice float64

		if reparsed["bid"] == nil {
			bidPrice = 0.0
		} else {
			bidPrice = reparsed["bid"].(float64)
		}

		if reparsed["ask"] == nil {
			askPrice = 0.0
		} else {
			askPrice = reparsed["ask"].(float64)
		}
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: symbol,
			BidPrice: bidPrice,
			BidQty: 0.0,
			AskPrice: askPrice,
			AskQty: 0.0,
			Exchange: "crex24",
			Timestamp: int(time.Now().Unix())})
	}
	if len(tickers) == 0 {
		log.Println("Could not get tickers from crex24.")
	}
	return tickers
}

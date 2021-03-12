package exchanges

import (
	"log"
	"strconv"
	"strings"
	"tickers/src/models"
	"time"
)

// GetBittrexTickers ...
func GetBittrexTickers() []models.Ticker {

	log.Println("Getting data from bittrex.")

	url := "https://api.bittrex.com/v3/markets/tickers"
	
	data := makeRequest(url)

	var tickers []models.Ticker

	parsed := data.([]interface{})

	for _, i := range parsed {

		reparsed := i.(map[string]interface{})

		symbol := reparsed["symbol"].(string)
		coin := strings.Split(symbol, "-")[0]
		cur := strings.Split(symbol, "-")[1]
		bidPrice, err := strconv.ParseFloat(reparsed["bidRate"].(string), 64)
		bidQty := 0.0
		askPrice, err := strconv.ParseFloat(reparsed["askRate"].(string), 64)
		askQty := 0.0
		
		if err != nil {
			log.Fatal(err)
		}
		
		tickers = append(tickers, models.Ticker{
			coin,
			cur,
			coin + cur,
			bidPrice,
			bidQty,
			askPrice,
			askQty,
			"bittrex",
			int(time.Now().Unix())})
	}

	if len(tickers) == 0 {
		log.Println("Could not get tickers from bittrex.")
	}

	return tickers
}
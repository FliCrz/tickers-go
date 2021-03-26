package bitstamp

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {
	
	symbols := getSymbols()
	curs := getCurrencies()
	
	for _, s := range symbols {
		
		url := "https://www.bitstamp.net/api/ticker/" + s
		data := utils.MakeRequest(url)
		parsed := data.(map[string]interface{})
		
		// log.Println(parsed)

		var coin string
		var cur string

		for _, c := range curs {
			if strings.Contains(s, c) {
				coin = strings.ToUpper(strings.Split(s, c)[0])
				cur = strings.ToUpper(c)
			}
		}
		
		bid, _ := strconv.ParseFloat(parsed["bid"].(string), 64)
		ask, _ := strconv.ParseFloat(parsed["ask"].(string), 64)

		newTicker := models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bid,
			BidQty: 0.0,
			AskPrice: ask,
			AskQty: 0.0,
			Exchange: "bitstamp",
			Timestamp: int(time.Now().Unix()),
		}

		tickers = append(tickers, newTicker)
		time.Sleep((1/700)* time.Second)
	}
	return tickers
}
package exchanges

import (
	"log"
	"strconv"
	"strings"
	"tickers/src/models"
	"time"
)

var hitbtcAPIURL =  "https://api.hitbtc.com/api/2"

// GetBinanceTickers ...
func GetHitbtcTickers() []models.Ticker {

	url := hitbtcAPIURL + "/public/ticker"
	
	data := makeRequest(url)

	var tickers []models.Ticker

	parsed := data.([]interface{})

	for _, i := range parsed {

		reparsed := i.(map[string]interface{})

		symbol := reparsed["symbol"].(string)
		coin := strings.SplitN(reparsed["symbol"].(string), "", 2)[0]
		cur := strings.SplitN(reparsed["symbol"].(string), "", 2)[1]
		

		bidPrice, _ := strconv.ParseFloat(reparsed["bid"].(string), 64)
		askPrice, _ := strconv.ParseFloat(reparsed["ask"].(string), 64)
		
		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: symbol,
			BidPrice: bidPrice,
			BidQty: 0.0,
			AskPrice: askPrice,
			AskQty: 0.0,
			Exchange: "hitbtc",
			Timestamp: int(time.Now().Unix())})
	}
	if len(tickers) == 0 {
		log.Println("Could not get tickers from hitbtc.")
	}
	return tickers
}

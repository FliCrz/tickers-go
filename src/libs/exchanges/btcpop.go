package exchanges

import (
	"log"
	"strconv"
	"tickers/src/models"
	"time"
)

// GetBtcPopTickers ...
func GetBtcPopTickers() []models.Ticker {

	url := "https://btcpop.co/api/market-public.php"
	
	data := makeRequest(url)

	var tickers []models.Ticker
	
	parsed := data.([]interface{})
	
	for _, i := range parsed {
		
		reparsed := i.(map[string]interface{})

		if reparsed["buyPrice"] != nil && reparsed["sellPrice"] != nil{
			coin := reparsed["ticker"].(string)
			bidPrice, _ := strconv.ParseFloat(reparsed["buyPrice"].(string), 64)
			askPrice, _ := strconv.ParseFloat(reparsed["sellPrice"].(string), 64)
			
			tickers = append(tickers, models.Ticker{
				Coin: coin,
				Currency: "BTC",
				Symbol: coin + "BTC",
				BidPrice: bidPrice,
				BidQty: 0.0,
				AskPrice: askPrice,
				AskQty: 0.0,
				Exchange: "btcpop",
				Timestamp: int(time.Now().Unix())})
		}
	}

	if len(tickers) == 0 {
		log.Println("Could not get tickers from btcpop.")
	}

	return tickers
}
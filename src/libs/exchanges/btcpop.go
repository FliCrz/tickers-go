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
			bidPrice, err := strconv.ParseFloat(reparsed["buyPrice"].(string), 64)
			bidQty := 0.0
			askPrice, err := strconv.ParseFloat(reparsed["sellPrice"].(string), 64)
			askQty := 0.0

			if err != nil {
				log.Fatal(err)
			}
			
			tickers = append(tickers, models.Ticker{
				coin,
				"BTC",
				coin + "BTC",
				bidPrice,
				bidQty,
				askPrice,
				askQty,
				"btcpop",
				int(time.Now().Unix())})
		}
	}

	if len(tickers) == 0 {
		log.Println("Could not get tickers from btcpop.")
	}

	return tickers
}
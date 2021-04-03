package cexio

import (
	"fmt"
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {

	curs := getCurrencies()
	
	// log.Println(curs)

	for k := range curs {

		raw := utils.MakeRequest(fmt.Sprintf("https://cex.io/api/tickers/%s", k))

		data := raw.(map[string]interface{})["data"]

		for _, i := range data.([]interface{})  {

			vol, _ := strconv.ParseFloat(i.(map[string]interface{})["volume"].(string), 64)

			t := models.Ticker {
				Coin: strings.Split(i.(map[string]interface{})["pair"].(string), ":")[0],
				Currency: strings.Split(i.(map[string]interface{})["pair"].(string), ":")[0],
				Symbol: strings.Replace(i.(map[string]interface{})["pair"].(string), ":", "", 1),
				BidPrice: i.(map[string]interface{})["bid"].(float64),
				BidQty: 0.0,
				AskPrice:  i.(map[string]interface{})["ask"].(float64),
				AskQty: 0.0,
				BaseVolume:  vol,
				QuoteVolume: 0.0,
				Exchange: "cexio",
				Timestamp: int(time.Now().Unix()),
			}
			tickers = append(tickers, t)
		}
	}

	return tickers
}
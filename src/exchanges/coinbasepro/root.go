package coinbasepro

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var baseUrl = "https://api.pro.coinbase.com"

// GetTickers ...
func GetTickers() []models.Ticker {
	symbols := getSymbols()
	var tickers []models.Ticker
	for _, s := range symbols {
		url := baseUrl + "/products/" + s + "/book"
		data := utils.MakeRequest(url)
		parsed := data.(map[string]interface{})
		
		
		bids := parsed["bids"].([]interface{})[0].([]interface{})
		asks := parsed["asks"].([]interface{})[0].([]interface{})
		
		// log.Println(bids, asks)

		bidPrice, _ := strconv.ParseFloat(bids[0].(string), 64)
		bidQty, _ := strconv.ParseFloat(bids[1].(string), 64)
		askPrice, _ := strconv.ParseFloat(asks[0].(string), 64)
		askQty, _ := strconv.ParseFloat(asks[1].(string), 64)
		new_ticker := models.Ticker{
			Coin: strings.Split(s, "-")[0],
			Currency: strings.Split(s, "-")[1],
			Symbol: strings.Replace(s, "-", "", -1),
			BidPrice: bidPrice,
			BidQty: bidQty,
			AskPrice: askPrice,
			AskQty: askQty,
			Exchange: "coinbasepro",
			Timestamp: int(time.Now().Unix())}
		
		tickers = append(tickers, new_ticker)
		time.Sleep(time.Duration(1000000000 / 6))
	}

	return tickers
}
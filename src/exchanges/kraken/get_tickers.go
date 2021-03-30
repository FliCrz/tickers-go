package kraken

import (
	"net/url"
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers() []models.Ticker {
	
	var tickers []models.Ticker

	symbolsRequest := url.Values{}
	symbols := getKrakenRawSymbols()
	var s string
	for n := range symbols {
		s += symbols[n].([]interface{})[0].(string) + ","
	}
	symbolsRequest.Add("pair", s[:len(s) - 1])
	url := "https://api.kraken.com/0/public/Ticker?" + symbolsRequest.Encode()
	data := utils.MakeRequest(url)
	parsed := data.(map[string]interface{})["result"]
	// log.Println(parsed)
	for k, v := range parsed.(map[string]interface{}) {
		var coin string
		var cur string
		for n := range symbols {
			if k == symbols[n].([]interface{})[0] {
				coin = strings.Split(symbols[n].([]interface{})[1].(string), "-")[0]
				cur = strings.Split(symbols[n].([]interface{})[1].(string), "-")[1]
			}
		}
		d := v.(map[string]interface{})
		bidData := d["b"].([]interface{})
		askData := d["a"].([]interface{})

		bidPrice, _ := strconv.ParseFloat(bidData[0].(string), 64)
		bidQty, _ := strconv.ParseFloat(bidData[2].(string), 64)
		askPrice, _ := strconv.ParseFloat(askData[0].(string), 64)
		askQty, _ := strconv.ParseFloat(askData[2].(string), 64)

		baseVol, _ := strconv.ParseFloat(d["v"].([]interface{})[1].(string), 64)
		quoteVol, _ := strconv.ParseFloat(d["p"].([]interface{})[1].(string), 64)

		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: k,
			BidPrice: bidPrice,
			BidQty: bidQty,
			AskPrice: askPrice,
			AskQty: askQty,
			BaseVolume: baseVol,
			QuoteVolume: quoteVol,
			Exchange: "kraken",
			Timestamp: int(time.Now().Unix())})
	}

	return tickers
}
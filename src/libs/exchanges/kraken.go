package exchanges

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"tickers/src/models"
	"time"
)

// GetKrakenTickers ...
func GetKrakenTickers() []models.Ticker {
	
	var tickers []models.Ticker
	
	fmt.Println("Getting data from kraken.")

	symbolsRequest := url.Values{}
	symbols := getKrakenRawSymbols()
	var s string
	for n := range symbols {
		s += symbols[n].([]interface{})[0].(string) + ","
	}
	symbolsRequest.Add("pair", s[:len(s) - 1])
	url := "https://api.kraken.com/0/public/Ticker?" + symbolsRequest.Encode()
	data := makeRequest(url)
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

		bidPrice, err := strconv.ParseFloat(bidData[0].(string), 64)
		bidQty, err := strconv.ParseFloat(bidData[2].(string), 64)
		askPrice, err := strconv.ParseFloat(askData[0].(string), 64)
		askQty, err := strconv.ParseFloat(askData[2].(string), 64)

		if err != nil {
			log.Fatalln(err)
		}

		tickers = append(tickers, models.Ticker{
			coin,
			cur,
			k,
			bidPrice,
			bidQty,
			askPrice,
			askQty,
			"kraken",
			int(time.Now().Unix())})
	}

	return tickers
}


func getKrakenRawSymbols () []interface{} {

	var symbols []interface{}
	url := "https://api.kraken.com/0/public/AssetPairs"
	data := makeRequest(url)
	parsed := data.(map[string]interface{})["result"]

	for k, v := range parsed.(map[string]interface{}) {
		pair :=  v.(map[string]interface{})["base"].(string) + "-" + v.(map[string]interface{})["quote"].(string)
		var pairData []interface{}
		pairData = append(pairData, k, pair)
		symbols = append(symbols, pairData)
	} 
	return symbols
}
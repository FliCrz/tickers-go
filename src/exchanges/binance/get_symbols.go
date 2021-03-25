package binance

import (
	"log"
	"tickers/src/utils"
)

func getSymbols() []interface{} {
	log.Println("Getting symbols from binance")

	url := APIURL + "/exchangeInfo"
	data := utils.MakeRequest(url)
	parsed := data.(map[string]interface{})
	symbolsList := parsed["symbols"].([]interface{})

	return symbolsList
}
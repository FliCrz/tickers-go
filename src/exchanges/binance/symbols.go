package binance

import (
	"tickers/src/utils"
)

func getSymbols() []interface{} {

	url := APIURL + "/exchangeInfo"
	data := utils.MakeRequest(url)
	parsed := data.(map[string]interface{})
	symbolsList := parsed["symbols"].([]interface{})

	return symbolsList
}
package bitfinex

import (
	"tickers/src/utils"
)

func getBitfinexRawSymbols () []interface{} {

	url := APIURL + "/conf/pub:list:currency"

	data := utils.MakeRequest(url)

	return data.([]interface{})
}
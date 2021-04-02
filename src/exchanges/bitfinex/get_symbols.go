package bitfinex

import (
	"log"
	"tickers/src/utils"
)

func getBitfinexRawSymbols () []interface{} {

	log.Println("Getting symbols from bitfinex.")

	url := APIURL + "/conf/pub:list:currency"

	data := utils.MakeRequest(url)

	return data.([]interface{})
}
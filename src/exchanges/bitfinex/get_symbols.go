package bitfinex

import (
	"log"
	"tickers/src/utils"
)

func getBitfinexRawSymbols () []interface{} {

	log.Println("Getting symbols data from bitfinex.")

	url := APIURL + "/conf/pub:map:currency:sym"

	data := utils.MakeRequest(url)

	var symbols []interface{}
	// log.Println(data)
	parsed := data.([]interface{})
	for _, i := range parsed {
		symbols = append(symbols, i.([]interface{}))
	}

	return symbols
}
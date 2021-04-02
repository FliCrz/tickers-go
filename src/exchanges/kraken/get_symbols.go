package kraken

import (
	"strings"
	"tickers/src/utils"
)


func getKrakenRawSymbols () []interface{} {

	var symbols []interface{}
	url := "https://api.kraken.com/0/public/AssetPairs"
	data := utils.MakeRequest(url)
	parsed := data.(map[string]interface{})["result"]

	for k, v := range parsed.(map[string]interface{}) {
		raw := v.(map[string]interface{})["wsname"]
		if raw != nil {
			pair := strings.Replace(raw.(string), "/", "-", 1)
			var pairData []interface{}
			pairData = append(pairData, k, pair)
			symbols = append(symbols, pairData)
		}
	} 
	return symbols
}
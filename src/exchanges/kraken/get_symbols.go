package kraken

import "tickers/src/utils"


func getKrakenRawSymbols () []interface{} {

	var symbols []interface{}
	url := "https://api.kraken.com/0/public/AssetPairs"
	data := utils.MakeRequest(url)
	parsed := data.(map[string]interface{})["result"]

	for k, v := range parsed.(map[string]interface{}) {
		pair :=  v.(map[string]interface{})["base"].(string) + "-" + v.(map[string]interface{})["quote"].(string)
		var pairData []interface{}
		pairData = append(pairData, k, pair)
		symbols = append(symbols, pairData)
	} 
	return symbols
}
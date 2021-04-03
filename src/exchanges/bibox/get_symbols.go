package bibox

import "tickers/src/utils"


func getSymbols () (symbols []string) {
	
	raw := utils.MakeRequest("https://api.bibox.com/v3/mdata/pairList")

	data := raw.(map[string]interface{})["result"]

	for _, i := range data.([]interface{}) {
		symbols = append(symbols, i.(map[string]interface{})["pair"].(string))
	}
	return symbols
}
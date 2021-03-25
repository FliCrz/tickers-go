package huobi

import "tickers/src/utils"


func getSymbols() []interface{} {
	url := "https://api.huobi.pro/v1/common/symbols"
	data := utils.MakeRequest(url)
	parsed := data.(map[string]interface{})
	reparsed := parsed["data"].([]interface{})
	var symbols []interface{}
	symbols = append(symbols, reparsed...)
	return symbols
}
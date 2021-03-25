package coinbasepro

import "tickers/src/utils"

func getSymbols () []string {
	url := baseUrl + "/products"
	data := utils.MakeRequest(url)
	var symbols []string
	for _, i := range data.([]interface{}) {
		d := i.(map[string]interface{})
		symbols = append(symbols, d["id"].(string))
	}
	return symbols
}
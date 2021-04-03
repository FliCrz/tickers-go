package cexio

import "tickers/src/utils"


func getCurrencies () (map[string]struct{}) {
	
	raw := utils.MakeRequest("https://cex.io/api/currency_limits")

	data := raw.(map[string]interface{})["data"].(map[string]interface{})["pairs"]

	curs := make(map[string]struct{})
	for _, i := range data.([]interface{}) {
		curs[i.(map[string]interface{})["symbol2"].(string)] = struct{}{}
	}
	return curs
}
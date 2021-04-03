package coinex

import (
	"fmt"
	"tickers/src/utils"
)


func getSymbols () (symbols []string) {
	url := baseUrl + "/market/info"

	raw := utils.MakeRequest(url)

	data := raw.(map[string]interface{})["data"]

	for _, v := range data.(map[string]interface{}) {
		symbols = append(symbols, fmt.Sprintf(
			"%s-%s",
			v.(map[string]interface{})["trading_name"].(string),
			v.(map[string]interface{})["pricing_name"].(string)))
	}
	return symbols
}
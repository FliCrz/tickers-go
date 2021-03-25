package huobi

import "strings"

func GenerateTradeUrl (coin, cur string) string {
	url := "https://www.huobi.com/en-us/exchange/" + strings.ToLower(coin) + "_" + strings.ToLower(cur)
	return url
}
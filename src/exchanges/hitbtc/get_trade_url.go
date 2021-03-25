package hitbtc

import "strings"

func GenerateTradeUrl (coin, cur string) string {
	url := "https://hitbtc.com/" + strings.ToLower(coin) + "-to-" + strings.ToLower(cur)
	return url
}
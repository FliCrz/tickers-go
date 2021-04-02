package bitstamp

import "strings"

func GenerateTradeUrl (coin, cur string) string {
	url := "https://www.bitstamp.net/markets/" + strings.ToLower(cur) + "/" +  strings.ToLower(coin)
	return url
}
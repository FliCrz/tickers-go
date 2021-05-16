package bitstamp

import "strings"

func GenerateTradeUrl (coin, cur string) string {
	url := "https://www.bitstamp.net/markets/" + strings.ToLower(coin) + "/" +  strings.ToLower(cur)
	return url
}
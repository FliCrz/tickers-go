package okex

import "strings"

func GenerateTradeUrl (coin, cur string) string {
	url := "https://www.okex.com/trade-spot/" + strings.ToLower(coin) + "-" + strings.ToLower(cur)
	return url
}
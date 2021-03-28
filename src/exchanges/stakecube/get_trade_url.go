package stakecube

import "strings"

func GenerateTradeUrl (coin, cur string) string {
	url := "https://stakecube.net/app/exchange/" + strings.ToLower(coin) + "_" + strings.ToLower(cur)
	return url
}
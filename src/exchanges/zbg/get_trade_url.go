package zbg

import "strings"

func GenerateTradeUrl (coin, cur string) string {
	url := "https://www.zbg.com/trade/" + strings.ToLower(coin) + "_" + strings.ToLower(cur)
	return url
}
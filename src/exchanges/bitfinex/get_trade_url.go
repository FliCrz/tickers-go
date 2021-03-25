package bitfinex

func GenerateTradeUrl (coin, cur string) string {
	url := "https://trading.bitfinex.com/t/" + coin + ":" + cur
	return url
}
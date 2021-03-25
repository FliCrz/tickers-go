package kucoin

func GenerateTradeUrl (coin, cur string) string {
	url := "https://trade.kucoin.com/spot/" + coin + "-" + cur
	return url
}
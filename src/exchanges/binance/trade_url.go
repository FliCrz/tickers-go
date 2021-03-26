package binance

func GenerateTradeUrl (coin, cur string) string {
	url := marketsURLBase + coin + "_" + cur
	return url
}
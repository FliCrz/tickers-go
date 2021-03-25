package kraken

func GenerateTradeUrl (coin, cur string) string {
	url := "https://trade.kraken.com/charts/KRAKEN:" + coin + "-" + cur
	return url
}
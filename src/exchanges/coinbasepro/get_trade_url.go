package coinbasepro

var makertsUrl = "https://pro.coinbase.com/trade/"

func GenerateTradeUrl (coin, cur string) string {
	url := makertsUrl + coin + "-" + cur
	return url
}
package bittrex

func GenerateTradeUrl (coin, cur string) string {
	url := "https://global.bittrex.com/Market/Index?MarketName=" + coin + "-" + cur
	return url
}
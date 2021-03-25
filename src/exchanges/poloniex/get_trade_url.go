package poloniex

func GenerateTradeUrl (coin, cur string) string {
	url := "https://poloniex.com/exchange/" + cur + "_" + coin
	return url
}
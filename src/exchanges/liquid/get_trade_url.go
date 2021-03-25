package liquid

func GenerateTradeUrl (coin, cur string) string {
	url := "https://app.liquid.com/exchange/" + coin + cur
	return url
}
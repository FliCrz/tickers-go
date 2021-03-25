package crex24

func GenerateTradeUrl (coin, cur string) string {
	url := "https://crex24.com/exchange/" + coin + "-" + cur
	return url
}
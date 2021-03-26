package bitstamp

func GenerateTradeUrl (coin, cur string) string {
	url := "https://www.bitstamp.net/markets/" + coin + "/" + cur
	return url
}
package btcpop

func GenerateTradeUrl (coin, cur string) string {
	url := "https://btcpop.co/Exchange/" + coin
	return url
}
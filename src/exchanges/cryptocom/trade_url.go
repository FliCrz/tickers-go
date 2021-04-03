package cryptocom

import "fmt"

func GenerateTradeUrl (coin, cur string) string {
	return fmt.Sprintf("https://crypto.com/exchange/trade/spot/%s_%s", coin, cur)
}
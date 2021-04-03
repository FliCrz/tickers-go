package coinex

import (
	"fmt"
	"strings"
)

// GenerateTradeUrl ...
func GenerateTradeUrl (coin, cur string) string {
	return fmt.Sprintf(
		"https://www.coinex.com/exchange/%s-%s", 
		strings.ToLower(coin),
		strings.ToLower(cur),
	)
}
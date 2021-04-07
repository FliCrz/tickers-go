package graviex

import (
	"fmt"
	"strings"
)

func GenerateTradeUrl (coin, cur string) string {
	return fmt.Sprintf(
		"https://graviex.net/markets/%s%s", 
		strings.ToLower(coin), 
		strings.ToLower(cur),
	)
}
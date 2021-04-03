package cexio

import (
	"fmt"
	"strings"
)

// GenerateTradeUrl ...
func GenerateTradeUrl (coin, cur string) string {
	return fmt.Sprintf("https://cex.io/%s-%s", strings.ToLower(coin), strings.ToLower(cur))
}
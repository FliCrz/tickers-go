package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// GetAvailableExchangesMethod ...
func GetAvailableExchangesMethod () []string {

	if Verbose {
		log.Println("Getting available exchanges")
	}

	exchanges := []string {
		"bequant",
		// "bibox",
		"binance",
		"bitfinex",
		"bitstamp",
		"bittrex",
		"btcpop",
		"cexio",
		"crex24",
		"crypto.com",
		// "coinbasepro",
		"coinex",
		"graviex",
		"hitbtc",
		"huobi",
		"kraken",
		"kucoin",
		"liquid", 
		"okex",
		"poloniex",
		"stakecube",
		"zbg",
	}

	return exchanges
}

var listExchangesCmd = &cobra.Command{
	Use:   "exchanges",
	Short: "list supported exchanges",
	Long: `list supported exchanges`,
	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		exchanges := GetAvailableExchangesMethod()
		if Verbose {
			for _, i := range exchanges {
				log.Println(i)
			}	
		}

		jsonData, _ := json.Marshal(exchanges)

		fmt.Printf("%s\n", string(jsonData))
	},
}
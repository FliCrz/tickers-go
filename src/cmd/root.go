package cmd

import (
	"fmt"
	"os"
	"tickers/src/libs/exchanges"
	"tickers/src/models"

	"github.com/spf13/cobra"
)

// Verbose ...
var Verbose bool

// TickersFuncs ...
var TickersFuncs = map[string]func() []models.Ticker {
	"binance": exchanges.GetBinanceTickers,
	"bitfinex": exchanges.GetBitfinexTickers,
	"bittrex": exchanges.GetBittrexTickers,
	"btcpop": exchanges.GetBtcPopTickers,
	"huobi": exchanges.GetHuobiTickers,
	"kraken": exchanges.GetKrakenTickers,
	"kucoin": exchanges.GetKucoinTickers,
	"okex": exchanges.GetOkexTickers,
	"poloniex": exchanges.GetPoloniexTickers,
	"liquid": exchanges.GetLiquidTickers,
	"hitbtc": exchanges.GetHitbtcTickers}

var rootCmd = &cobra.Command{
	Use:   "tickers",
	Short: "tickers - Arbitrage Tool",
	Long: `tickers is a cli tool to detect arbitrage opportunities among supported exchanges.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(listExchangesCmd)
	rootCmd.AddCommand(tickersCmd)
	rootCmd.AddCommand(listCommonSymbolsCmd)
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
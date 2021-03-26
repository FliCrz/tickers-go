package cmd

import (
	"fmt"
	"os"
	"tickers/src/exchanges/binance"
	"tickers/src/exchanges/bitfinex"
	"tickers/src/exchanges/bitstamp"
	"tickers/src/exchanges/bittrex"
	"tickers/src/exchanges/btcpop"
	"tickers/src/exchanges/coinbasepro"
	"tickers/src/exchanges/crex24"
	"tickers/src/exchanges/hitbtc"
	"tickers/src/exchanges/huobi"
	"tickers/src/exchanges/kraken"
	"tickers/src/exchanges/kucoin"
	"tickers/src/exchanges/liquid"
	"tickers/src/exchanges/okex"
	"tickers/src/exchanges/poloniex"
	"tickers/src/models"

	"github.com/spf13/cobra"
)

// Verbose ...
var Verbose bool

// TickersFuncs ...
var TickersFuncs = map[string]func() []models.Ticker {
	"binance": binance.GetTickers,
	"bitfinex": bitfinex.GetTickers,
	"bitstamp": bitstamp.GetTickers,
	"bittrex": bittrex.GetTickers,
	"btcpop": btcpop.GetTickers,
	"crex24": crex24.GetTickers,
	"coinbasepro": coinbasepro.GetTickers,
	"hitbtc": hitbtc.GetTickers,
	"huobi": huobi.GetTickers,
	"kraken": kraken.GetTickers,
	"kucoin": kucoin.GetTickers,
	"okex": okex.GetTickers,
	"poloniex": poloniex.GetTickers,
	"liquid": liquid.GetTickers}

// UrlFuncs ...
var UrlFuncs = map[string]func(string, string) string {
	"binance": binance.GenerateTradeUrl,
	"bitfinex": bitfinex.GenerateTradeUrl,
	"bitstamp": bitstamp.GenerateTradeUrl,
	"bittrex": bittrex.GenerateTradeUrl,
	"btcpop": btcpop.GenerateTradeUrl,
	"crex24": crex24.GenerateTradeUrl,
	"coinbasepro": coinbasepro.GenerateTradeUrl,
	"hitbtc": hitbtc.GenerateTradeUrl,
	"huobi": huobi.GenerateTradeUrl,
	"kraken": kraken.GenerateTradeUrl,
	"kucoin": kucoin.GenerateTradeUrl,
	"okex": okex.GenerateTradeUrl,
	"poloniex": poloniex.GenerateTradeUrl,
	"liquid": liquid.GenerateTradeUrl}

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
	rootCmd.AddCommand(getTradeURLCmd)
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
package utils

import (
	"log"
	"tickers/src/models"
)

// GetAvailableExchangesMethod ...
func GetAvailableExchangesMethod () []string {

	log.Println("Getting available exchanges")

	exchanges := []string {
		"binance",
		"bitfinex",
		"bittrex",
		"btcpop",
		"huobi",
		"kraken",
		"kucoin",
		"okex",
		"poloniex",
		"liquid"}

	return exchanges
}

// GetCommonSymbolsMethod ...
func GetCommonSymbolsMethod (args []string) []string {

	log.Printf("Getting common symbols for %s\n", args)
	var tickers [][]models.Ticker
		
	for _, i := range args {
		tickers = append(tickers, TickersFuncs[i]())
	}

	commonSymbols := GetCommonSymbols(tickers)
	return commonSymbols
}

package utils

import (
	"log"
	"tickers/src/models"
)

// GetTickersMethod ...
func GetTickersMethod(args []string) [][]models.Ticker{

	log.Printf("Getting Tickers for %s\n", args)

	var tickers [][]models.Ticker
	
	for _, i := range args {
		tickers = append(tickers, TickersFuncs[i]())
	}

	return tickers
}
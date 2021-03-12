package utils

import (
	"log"
	"tickers/src/models"
)

// GetCommonSymbols ...
func GetCommonSymbols(data [][]models.Ticker) []string {

	log.Println("Getting common symbols")

	var n int
	var symbols []string;

	for n = 0; n < len(data) -1; n++ {
		for _, d := range data[n] {
			for _, i := range data[n + 1] {
				if d.Symbol == i.Symbol {
					symbols = append(symbols, d.Symbol)
				}
			}
		}
	}

	if len(symbols) == 0 {
		log.Println("No common symbols found!")
	}
	
	return symbols
}
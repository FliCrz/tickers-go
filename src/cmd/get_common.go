package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"tickers/src/models"

	"github.com/spf13/cobra"
)

func generateMatrix () (map[string]map[string][]string) {
	
	common := make(map[string]map[string][]string)
	
	exchanges := GetAvailableExchangesMethod()

	var tickers [][]models.Ticker

	for n := 0; n < len(exchanges) - 1; n++ {
		tickers = append(tickers, TickersFuncs[exchanges[n]]())
		tickers = append(tickers, TickersFuncs[exchanges[n + 1]]())
		symbols := GetCommonSymbols(tickers)
		d := make(map[string][]string)
		if symbols != nil {
			d[exchanges[n + 1]] = symbols
			common[exchanges[n]] = d
		}
	}

	j, _ := json.Marshal(common)
	ioutil.WriteFile("common_matrix.json", j, 0643)
	return common
}

// GetCommonSymbols ...
func GetCommonSymbols(data [][]models.Ticker) []string {

	if Verbose {
		log.Println("Getting common symbols")
	}

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

// GetCommonSymbolsMethod ...
func GetCommonSymbolsMethod (args []string) []string {

	if Verbose {
		log.Printf("Getting common symbols for %s\n", args)
	}
	
	var tickers [][]models.Ticker
		
	for _, i := range args {
		tickers = append(tickers, TickersFuncs[i]())
	}

	commonSymbols := GetCommonSymbols(tickers)
	return commonSymbols
}

var listCommonSymbolsCmd = &cobra.Command{
	Use:   "common [exchange] [exchange - optional]",
	Short: "list common symbols for 2 or more exchanges",
	Long: `list common symbols for 2 or more exchanges`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "matrix" {
			commonSymbols := generateMatrix()
	
			common, _ := json.Marshal(commonSymbols)
			fmt.Printf("%s\n", string(common))	
		} else {
			commonSymbols := GetCommonSymbolsMethod(args)
	
			common, _ := json.Marshal(commonSymbols)
			fmt.Printf("%s\n", string(common))	
		}
	},
}
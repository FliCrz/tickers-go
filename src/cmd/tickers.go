package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"tickers/src/models"

	"github.com/spf13/cobra"
)

// GetTickersMethod ...
func GetTickersMethod(args []string) [][]models.Ticker{

	if Verbose {
		log.Printf("Getting Tickers for %s\n", args)
	}

	var tickers [][]models.Ticker
	
	for _, i := range args {
		tickers = append(tickers, TickersFuncs[i]())
	}

	return tickers
}

var tickersCmd = &cobra.Command{
	Use:   "tickers [exchange] [exchange - optional]",
	Short: "get tickers for 1 or more excanges",
	Long: `get tickers for 1 or more excanges`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		tickers := GetTickersMethod(args)

		if Verbose == true {
			for _, i := range tickers{
				for _, o := range i {
					
					jsonData, err := json.Marshal(o)
					
					if err != nil {
						log.Println(err)
					}

					log.Println(string(jsonData))
				}
			}
		}

		jsonData, err := json.Marshal(tickers)

		if err != nil {
			log.Println(err)
		}

		fmt.Printf("%s\n", string(jsonData))
	},
}


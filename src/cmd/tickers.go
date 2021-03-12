package cmd

import (
	"encoding/json"
	"log"
	"tickers/src/utils"

	"github.com/spf13/cobra"
)


var tickersCmd = &cobra.Command{
	Use:   "tickers [exchange] [exchange - optional]",
	Short: "get tickers for 1 or more excanges",
	Long: `get tickers for 1 or more excanges`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		tickers := utils.GetTickersMethod(args)

		for _, i := range tickers{
			for _, o := range i {
				if Verbose == true {
					jsonData, err := json.Marshal(o)
					if err != nil {
						log.Fatalln(err)
					}
					log.Println(string(jsonData))
				} else {
					log.Println(o)
				}
			}
		}
	},
}


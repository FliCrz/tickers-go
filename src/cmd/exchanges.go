package cmd

import (
	"log"
	"tickers/src/utils"

	"github.com/spf13/cobra"
)

var listExchangesCmd = &cobra.Command{
	Use:   "exchanges",
	Short: "list supported exchanges",
	Long: `list supported exchanges`,
	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		exchanges := utils.GetAvailableExchangesMethod()	
		for _, i := range exchanges {
			log.Println(i)
		}	
	},
}
package cmd

import (
	"log"
	"tickers/src/utils"

	"github.com/spf13/cobra"
)

var listCommonSymbolsCmd = &cobra.Command{
	Use:   "common [exchange] [exchange - optional]",
	Short: "list common symbols for 2 or more exchanges",
	Long: `list common symbols for 2 or more exchanges`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		commonSymbols := utils.GetCommonSymbolsMethod(args)
		log.Println("===== Common Pairs =====")
		log.Println(commonSymbols)	
	},
}
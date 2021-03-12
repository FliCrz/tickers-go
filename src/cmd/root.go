package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "tickers",
	Short: "tickers - Arbitrage Tool",
	Long: `tickers is a cli tool to detect arbitrage opportunities among supported exchanges.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {},
}

// Verbose ...
var Verbose bool

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
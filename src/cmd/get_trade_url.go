package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)


var getTradeURLCmd = &cobra.Command{
	Use:   "url <exchange> <coin> <cur>",
	Short: "get exchange url for coin cur pair",
	Long: `Get exchange url for coin cur pair`,
	Args: cobra.ExactArgs(3),

	Run: func(cmd *cobra.Command, args []string) {
		url := UrlFuncs[args[0]](strings.ToUpper(args[1]), strings.ToUpper(args[2]))

		jsonData, _ := json.Marshal(url)

		fmt.Printf("%s\n", string(jsonData))
	},
}
package cmd

import (
	"github.com/spf13/cobra"
)

var normalCmd = &cobra.Command{
	Use:   "normal",
	Short: "Normal distribution",
	Long: `
	Normally distributed random numbers

	Arguments: mean stdev  default: 0.0 1.0
	
	`,
	Run: Normal,
}

func init() {
	rootCmd.AddCommand(normalCmd)
}

func Normal(cmd *cobra.Command, args []string) {

}

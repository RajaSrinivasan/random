package cmd

import (
	"github.com/spf13/cobra"
	//"gitlab.com/RajaSrinivasan/random/impl/series"
)

var uniformCmd = &cobra.Command{
	Use:   "uniform",
	Short: "Uniformly distributed random variables",
	Long: `
	Uniformly distributed random variables

	Arguments:  min max default 0.0 1.0
	`,
	Run: Uniform,
}

func init() {
	rootCmd.AddCommand(uniformCmd)
}

func Uniform(cmd *cobra.Command, args []string) {

}

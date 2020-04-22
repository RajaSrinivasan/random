package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/random/impl/generate"
)

var Verbose bool
var crypto bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "random",
	Short: "Explore random number sources",
	Long: `
	This utility generates random numbers and plots the data as a series or as a histogram.
	`,
	Version: "v0.2",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&Verbose, "verbose", false, "be verbose")

	rootCmd.PersistentFlags().BoolVar(&generate.Series, "series", false, "generate time series plots. default is histogram")
	rootCmd.PersistentFlags().StringVarP(&generate.Output, "output", "o", "plot.png", "output file name")
	rootCmd.PersistentFlags().StringVarP(&generate.Table, "table", "t", "", "tabular output filename")
	rootCmd.PersistentFlags().IntVarP(&generate.Samples, "samples", "s", 1000, "number of samples")
	rootCmd.PersistentFlags().IntVarP(&generate.Slices, "slices", "c", 10, "number of slices - for histograms")
	rootCmd.PersistentFlags().IntVarP(&generate.Seed, "seed", "d", 1729, "seed for random numbers")
	rootCmd.PersistentFlags().BoolVarP(&crypto, "crypto", "C", false, "use cryptographic random source ")

}

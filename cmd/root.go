package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool
var series bool
var output string
var table string
var samples int
var slices int
var seed int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "random",
	Short: "Explore random number sources",
	Long: `
	This utility generates random numbers and plots the data as a series or as a histogram.
	`,
	Version: "v0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "be verbose")

	rootCmd.PersistentFlags().BoolVar(&series, "series", false, "generate time series plots. default is histogram")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "plot.png", "output file name")
	rootCmd.PersistentFlags().StringVarP(&table, "table", "t", "", "tabular output filename")
	rootCmd.PersistentFlags().IntVarP(&samples, "samples", "s", 1000, "number of samples")
	rootCmd.PersistentFlags().IntVarP(&slices, "slices", "c", 10, "number of slices - for histograms")
	rootCmd.PersistentFlags().IntVarP(&seed, "seed", "d", 1729, "seed for random numbers")

}

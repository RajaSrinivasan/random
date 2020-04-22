package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/random/impl/generate"
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
	var mean = 0.0
	var stdev = 1.0
	var err error

	if len(args) >= 1 {
		mean, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			panic(err)
		}
		if len(args) >= 2 {
			stdev, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				panic(err)
			}
		}
	}

	if Verbose {
		fmt.Printf("Normal random variables mean %f stdev %f\n", mean, stdev)
	}
	if crypto {
		series := generate.CryptoNormal(mean, stdev)
		series.Show()
		series.Plot()
	} else {
		series := generate.Normal(mean, stdev)
		series.Show()
		series.Plot()
	}
}

package cmd

import (
	"fmt"
	"os"
	"strconv"

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
	var minimum = 0.0
	var maximum = 1.0
	var err error

	if len(args) >= 1 {
		minimum, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			panic(err)
		}
		if len(args) >= 2 {
			maximum, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				panic(err)
			}
		}
	}
	if maximum <= minimum {
		fmt.Printf("Invalid Range\n")
		os.Exit(1)
	}
	if verbose {
		fmt.Printf("Uniform random variables in the range min %f max %f\n", minimum, maximum)
	}
}

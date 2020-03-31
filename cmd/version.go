package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/random/impl/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Report the version of the application",
	Long: `
	Report source repository versions
	`,
	Run: Version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Version(cmd *cobra.Command, args []string) {
	version.Report()
}

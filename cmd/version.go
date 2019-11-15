package cmd

import (
  "github.com/vulogov/MonitoringBund/ctx"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "bund version",
	Long:  `Displays current version of the MonitoringBund package`,
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(ctx.Logo.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nhn-cloud-api",
	Long:  `All software has versions. This is nhn-cloud-api's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nhn-cloud-api generator v0.9")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

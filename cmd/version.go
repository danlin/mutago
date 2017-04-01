package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of MuTaGo",
	Long:  "All software has versions. This is MuTaGo's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MuTaGo Musik Library v0.1")
	},
}

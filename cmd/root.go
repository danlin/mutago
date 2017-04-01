package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "mutago",
	Short: "MuTaGo Musik Library",
	Long: `A Fast Musik Indexing Libray built with love by danlin in Go.
Complete documentation is available at http://mutago.8plugs.de`,
}

var Verbose bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
